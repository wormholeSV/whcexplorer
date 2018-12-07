package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bcext/cashutil"
	"github.com/copernet/whccommon/model"
	"github.com/copernet/whcexplorer/config"
	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

const (
	BalancePrefix = "balance:"

	// WARNING: the maintained address set only supports bash32 encoded address
	AddressSetKey = "address:set"

	// channels
	// update block tip
	UpdateBlockTip = "block:tip"
	// notify balance for wormhole
	BalanceUpdated = "balance:wormhole:updated"

	BalanceExpire = 120
)

type RedisOperationSet interface {
	GetBalanceForAddress([]string) (map[string][]model.BalanceForAddress, error)
	StoreBalanceForAddress(bal map[string][]model.BalanceForAddress) error
	Publish(channel string, msg string) error
	Subscribe(ctx context.Context, channel string, ret chan<- []byte) error
	AddAddress(addrs []string) error
	DelAddress(addr string) error
}

type RedisOption struct {
	Host             string
	Port             int
	Password         string
	DB               int
	IdleTimeout      int
	MaxIdleConns     int
	MaxOpenConns     int
	InitialOpenConns int
}

type RedisStorage struct {
	pool   *redis.Pool
	config RedisOption
}

func NewRedis(option RedisOption) (RedisOperationSet, error) {
	r := &RedisStorage{
		config: option,
		pool: &redis.Pool{
			MaxIdle:     option.MaxIdleConns,
			MaxActive:   option.MaxOpenConns,
			IdleTimeout: time.Duration(option.IdleTimeout) * time.Second,
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		},
	}

	r.pool.Dial = r.createRedisConnect

	if option.InitialOpenConns > option.MaxIdleConns {
		option.InitialOpenConns = option.MaxIdleConns
	} else if option.InitialOpenConns == 0 {
		option.InitialOpenConns = 1
	}

	return r, r.initRedisConnect()
}

func (r *RedisStorage) createRedisConnect() (redis.Conn, error) {
	return redis.Dial(
		"tcp",
		fmt.Sprintf("%s:%d", r.config.Host, r.config.Port),
		redis.DialDatabase(r.config.DB),
		redis.DialPassword(r.config.Password),
	)
}

func (r *RedisStorage) initRedisConnect() error {
	cons := make([]redis.Conn, r.config.InitialOpenConns)
	defer func() {
		for _, c := range cons {
			if c != nil {
				c.Close()
			}
		}
	}()

	for i := 0; i < r.config.InitialOpenConns; i++ {
		cons[i] = r.pool.Get()
		if _, err := cons[i].Do("PING"); err != nil {
			return err
		}
	}

	return nil
}

func (r *RedisStorage) GetBalanceForAddress(addresses []string) (map[string][]model.BalanceForAddress, error) {
	ri := r.pool.Get()
	defer ri.Close()

	// the balMap is first return result, and initialize its value via make.
	// so the caller should not use `ret == nil` to justify whether the result is empty
	// or not.
	balMap := make(map[string][]model.BalanceForAddress)
	for _, addr := range addresses {
		_, err := cashutil.DecodeAddress(addr, config.GetChainParam())
		if err != nil {
			// skip to the next address.
			continue
		}

		res, err := redis.String(ri.Do("GET", BalancePrefix+addr))

		if err != nil {
			// continue for the next address if null data in redis database
			if strings.HasSuffix(err.Error(), "nil returned") {
				continue
			}

			return nil, err
		}

		var bal []model.BalanceForAddress
		err = json.Unmarshal(bytes.NewBufferString(res).Bytes(), &bal)
		if err != nil {
			return nil, err
		}

		balMap[addr] = bal
	}

	return balMap, nil
}

func (r *RedisStorage) StoreBalanceForAddress(bals map[string][]model.BalanceForAddress) error {
	if bals == nil {
		return errors.New("empty balance result")
	}

	ri := r.pool.Get()
	defer ri.Close()

	for addr, bal := range bals {
		b, err := json.Marshal(bal)
		if err != nil {
			return err
		}

		// use pipline style to optimise redis operation
		err = ri.Send("SET", BalancePrefix+addr, string(b), "EX", BalanceExpire)
		if err != nil {
			logrus.WithField("address", addr).
				Error("send store balance command for address failed:", err)
			continue
		}
	}

	err := ri.Flush()
	if err != nil {
		return err
	}

	_, err = ri.Receive()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisStorage) Publish(channel string, msg string) error {
	ri := r.pool.Get()
	defer ri.Close()

	_, err := ri.Do("PUBLISH", channel, msg)

	return err
}

func (r *RedisStorage) Subscribe(ctx context.Context, channel string, ret chan<- []byte) error {
	ri := r.pool.Get()
	defer ri.Close()

	// A ping is set to the server with this period to test for the health of
	// the connection and server.
	const healthCheckPeriod = time.Minute

	psc := redis.PubSubConn{Conn: ri}

	if err := psc.Subscribe(channel); err != nil {
		return err
	}

	done := make(chan error, 1)

	// Start a goroutine to receive notifications from the server.
	go func() {
		for {
			switch n := psc.Receive().(type) {
			case error:
				done <- n
				return
			case redis.Message:
				ret <- n.Data
			case redis.Subscription:
				if n.Count == 0 {
					// Return from the goroutine when all channels are unsubscribed.
					done <- nil
					return
				}
			}
		}
	}()

	ticker := time.NewTicker(healthCheckPeriod)
	defer ticker.Stop()

	var err error
loop:
	for err == nil {
		select {
		case <-ticker.C:
			// Send ping to test health of connection and server. If
			// corresponding pong is not received, then receive on the
			// connection will timeout and the receive goroutine will exit.
			if err = psc.Ping(""); err != nil {
				break loop
			}
		case <-ctx.Done():
			break loop
		case err := <-done:
			// Return error from the receive goroutine.
			return err
		}
	}

	// Signal the receiving goroutine to exit by unsubscribing from all channels.
	psc.Unsubscribe()

	// Wait for goroutine to complete.
	return <-done
}

func (r *RedisStorage) AddAddress(addrs []string) error {
	ri := r.pool.Get()
	defer ri.Close()

	for _, addr := range addrs {
		err := ri.Send("SADD", AddressSetKey, addr)
		if err != nil {
			return err
		}
	}

	return ri.Flush()
}

func (r *RedisStorage) DelAddress(addrs string) error {
	ri := r.pool.Get()
	defer ri.Close()

	_, err := ri.Do("SREM", AddressSetKey, addrs)
	return err
}
