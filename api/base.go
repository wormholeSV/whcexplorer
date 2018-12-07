package api

import (
	"github.com/copernet/whccommon/log"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
	"strings"

	"github.com/bcext/cashutil"
	"github.com/bcext/gcash/chaincfg/chainhash"
	"github.com/pkg/errors"
)

type Response struct {
	Code    errCode     `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type ParamType int

const (
	UnknownType ParamType = iota

	BlockHash
	BlockHeight
	TxHash

	// pagination feature
	defaultPageSizeNumber = 50
	defaultPageNo         = 1

	propertyNameLimit = 520
)

// How to distinguish block hash /transaction hash / block height:
//
// block powlimit:
// mainnet: 0x00000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff
// testnet: 0x00000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff
// regtest: 0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff
//
// so distinguish block hash from transaction hash:
// 1. The parameter string must prefix with '00000000' (8 '0' at least), here it may be block most likely.
// When it have been found in blockchain, we assume it represents block hash. Next, if it
// not a existed block hash, we start search transaction with the string.
func paramType(keyword string) (ParamType, error) {
	// First justify whether is block height or not. Max int32 is 2147483647, and the length
	// is 10.
	if len(keyword) <= 10 {
		_, ok := checkBlockHeight(keyword)
		if !ok {
			return UnknownType, errors.New("invalid block height")
		}

		return BlockHeight, nil
	}

	if checkSha256Hash(keyword) {
		// The keyword maybe represent block hash most likely. But it maybe represent
		// transaction hash with less possibility.
		if strings.HasPrefix(keyword, "00000000") {
			return BlockHash, nil
		} else {
			return TxHash, nil
		}
	}

	return UnknownType, errors.New("invalid parameter")
}

func checkSha256Hash(hash string) bool {
	if len(hash) != 2*chainhash.HashSize &&
		len(hash) != 2*chainhash.HashSize-1 {
		return false
	}

	_, err := chainhash.NewHashFromStr(hash)
	if err != nil {
		return false
	}

	return true
}

// checkBlockHeight is a convenient function for checking http request parameter.
// So the function accept the block height string format.
func checkBlockHeight(height string) (int32, bool) {
	h, err := strconv.Atoi(height)
	if err != nil {
		return cashutil.BlockHeightUnknown, false
	}

	if h < 0 || h > math.MaxInt32 {
		return cashutil.BlockHeightUnknown, false
	}

	return int32(h), true
}

// Time the block was created.  This is, unfortunately, encoded as a
// uint32 on the wire and therefore is limited to 2106.
func checkBlockTime(t string) (uint32, bool) {
	timestamp, err := strconv.Atoi(t)
	if err != nil {
		return 0, false
	}

	if timestamp > math.MaxUint32 || timestamp < 0 {
		return 0, false
	}

	return uint32(timestamp), true
}

func paginator(c *gin.Context) (int, int) {
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", strconv.Itoa(defaultPageSizeNumber)))
	pageNo, err := strconv.Atoi(c.DefaultQuery("pageNo", strconv.Itoa(defaultPageNo)))
	if err != nil {
		log.WithCtx(c).Errorf("query property history list parameter error, pagesize: %s pageno: %s",
			c.Query("pageSize"), c.Query("pageNo"))
	}

	if pageSize <= 0 {
		pageSize = defaultPageSizeNumber
	}

	if pageNo <= 0 {
		pageNo = defaultPageNo
	}

	return pageSize, pageNo
}