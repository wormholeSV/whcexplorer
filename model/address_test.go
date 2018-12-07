package model_test

import (
	"testing"
	"github.com/copernet/whcexplorer/model"
	"github.com/shopspring/decimal"
	model2 "github.com/copernet/whccommon/model"
	"github.com/stretchr/testify/assert"
)

func TestWhcBalance(t *testing.T) {
	tests := []struct {
		name string
		address string
		avail decimal.Decimal
		total decimal.Decimal
		err error
	}{
		{},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			avail, total, err := model.WhcBalance(test.address)
			assert.Equal(t, avail, test.avail)
			assert.Equal(t, total, test.total)
			assert.Equal(t, err, test.err)
		})
	}
}
func TestWhcStatistics(t *testing.T) {
	tests := []struct {
		name string
		address string
		consumed decimal.Decimal
		sended decimal.Decimal
		received decimal.Decimal
		err error
	}{
		{},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			consumed, sended, received, err := model.WhcStatistics(test.address)
			assert.Equal(t, consumed, test.consumed)
			assert.Equal(t, sended, test.sended)
			assert.Equal(t, received, test.received)
			assert.Equal(t, err, test.err)
		})
	}
}

func TestWhcTxCountStatistics(t *testing.T) {
	tests := []struct {
		name string
		address string
		count int
		err error
	}{
		{},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			count, err := model.WhcTxCountStatistics(test.address)
			assert.Equal(t, count, test.count)
			assert.Equal(t, err, test.err)
		})
	}
}

func TestListAddressTxes(t *testing.T) {
	tests := []struct {
		name string
		address string
		txType int
		times []int64
		propertyId int64
		pageInfo model2.PageInfo
		count int
		directions []int
		err error
	}{
		{},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list, directions, error := model.ListAddressTxes(test.address, test.txType, test.times, test.propertyId, &test.pageInfo)
			assert.Equal(t, len(list), test.count)
			assert.Equal(t, error, test.err)
			assert.Equal(t, directions, test.directions)
		})
	}
}
func TestCountAddressTxes(t *testing.T) {
	tests := []struct {
		name string
		address string
		txType int
		times []int64
		propertyId int64
		count int
	}{
		{},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			count := model.CountAddressTxes(test.address, test.txType, test.times, test.propertyId)
			assert.Equal(t, count, test.count)
		})
	}
}
func TestListAddressProperties(t *testing.T) {
	tests := []struct {
		name string
		address string
		propertyType int
		pageInfo model2.PageInfo
		count int
		err error
	}{
		{},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list, err := model.ListAddressProperties(test.address, test.propertyType, &test.pageInfo)
			assert.Equal(t, len(list), test.count)
			assert.Equal(t, err, test.err)
		})
	}
}
func TestCountAddressProperties(t *testing.T) {
	tests := []struct {
		name string
		address string
		propertyType int
		count int
	}{
		{},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			count := model.CountAddressProperties(test.address, test.propertyType)
			assert.Equal(t, count, test.count)
		})
	}
}




