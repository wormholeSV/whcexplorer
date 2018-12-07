package model_test

import (
	"testing"
	"github.com/copernet/whccommon/model"
	model2 "github.com/copernet/whcexplorer/model"
	"github.com/stretchr/testify/assert"
	"github.com/copernet/whcexplorer/api/common"
)

var pageInfos = []model.PageInfo {
	{Page:1, PageSize:model.DEFAULT_PAGE_SIZE},
	{Page:1, PageSize:20},
	{Page:1, PageSize:2},
}
func TestListProperties(t *testing.T) {

	tests := []struct {
		name string
		PageInfo model.PageInfo
		ProertyType []int
		Total int
		Error error
	}{
		{"lp0",pageInfos[0], nil, 50, nil},
		{"lp1",pageInfos[0], []int{common.EVENT_TYPE_CREATE_FIXED_TOKEN}, 50, nil},
		{"lp2",pageInfos[1], []int{common.EVENT_TYPE_CREATE_FIXED_TOKEN, common.EVENT_TYPE_CREATE_CROWD_TOKEN, common.EVENT_TYPE_CREATE_MANAGE_TOKEN}, 20, nil},
		{"lp3",pageInfos[1], []int{1000}, 0, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret, err := model2.ListProperties(test.ProertyType, &test.PageInfo)
			assert.Equal(t, err ,test.Error, "err not expected")
			assert.Equal(t, len(ret), test.Total, "length not expected")
		})
	}
}

func TestCountProperties(t *testing.T) {
	tests := []struct {
		name string
		ProertyType []int
		Total int
	}{
		{"cp0", nil, 50},
		{"cp1", []int{common.EVENT_TYPE_CREATE_FIXED_TOKEN}, 50},
		{"cp2", []int{common.EVENT_TYPE_CREATE_FIXED_TOKEN, common.EVENT_TYPE_CREATE_CROWD_TOKEN, common.EVENT_TYPE_CREATE_MANAGE_TOKEN}, 20},
		{"cp3", []int{1000}, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret := model2.CountProperties(test.ProertyType)
			assert.Equal(t, ret, test.Total, "length not expected")
		})
	}
}

func TestPropertyDetail(t *testing.T) {
	tests := []struct {
		name string
		queryPropertyId int64
		expectPropertyId int64
		err error
	}{
		{"pd0", 1, 1,nil},
		{"pd1", 3, 3,nil},
		{"pd2", -1, -1,nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret, err := model2.PropertyDetail(test.queryPropertyId)
			assert.Equal(t, ret, test.expectPropertyId, "propertyDetail not expected")
			assert.Equal(t, err, test.err, "error not expected")
		})
	}
}

func TestListPropertyChangeLog(t *testing.T) {
	tests := []struct {
		name string
		PageInfo model.PageInfo
		ProertyId int64
		Total int
		Error error
	}{
		{"lpc0",pageInfos[0], -1, 0, nil},
		{"lpc1",pageInfos[0], 27, 50, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret, err := model2.ListPropertyChangeLog(test.ProertyId, &test.PageInfo)
			assert.Equal(t, err ,test.Error, "err not expected")
			assert.Equal(t, len(ret), test.Total, "length not expected")
		})
	}
}
func TestCountPropertyChangeLog(t *testing.T) {
	tests := []struct {
		name string
		ProertyId int64
		Total int
	}{
		{"lpc0", -1, 0},
		{"lpc1", 27, 50},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret := model2.CountPropertyChangeLog(test.ProertyId)
			assert.Equal(t, ret, test.Total, "total not expected")
		})
	}
}

func TestListPropertyTxes(t *testing.T) {
	tests := []struct {
		name string
		PageInfo model.PageInfo
		ProertyId int64
		EventTypes int
		Total int
		Error error
	}{
		{"lpc0",pageInfos[0], -1, -1,0, nil},
		{"lpc1",pageInfos[0], 27, 50,50, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret, err := model2.ListPropertyTxes(test.ProertyId, test.EventTypes, &test.PageInfo)
			assert.Equal(t, err ,test.Error, "err not expected")
			assert.Equal(t, len(ret), test.Total, "length not expected")
		})
	}
}

func TestCountPropertyTxes(t *testing.T) {
	tests := []struct {
		name string
		ProertyId int64
		EventTypes int
		Total int
	}{
		{"lpc0", -1, -1,0},
		{"lpc1", 27, 50,50},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret := model2.CountPropertyTxes(test.ProertyId, test.EventTypes)
			assert.Equal(t, ret, test.Total, "total not expected")
		})
	}
}
