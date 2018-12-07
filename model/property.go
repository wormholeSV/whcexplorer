package model

import (
	"errors"
	"github.com/copernet/whccommon/model"
	"github.com/copernet/whcexplorer/api/common"
	"github.com/copernet/whcexplorer/model/view"
	"github.com/copernet/whcexplorer/util"
	"strings"
	"github.com/jinzhu/gorm"
)

type PropertyObj struct {
	PropertyID          int64  `json:"propertyId"`
	PropertyName        string `json:"propertyName"`
	PropertyCategory    string `json:"propertyCategory"`
	PropertySubcategory string `json:"propertySubcategory"`
	PropertyServiceURL  string `json:"propertyServiceUrl"`
	Issuer              string `json:"issuer"`
	TxType              string `json:"txType"`
}

func ListProperties(propertyType []int, pageInfo *model.PageInfo) ([]PropertyObj, error) {
	if len(propertyType) == 0 {
		propertyType = []int{common.EVENT_TYPE_CREATE_FIXED_TOKEN, common.EVENT_TYPE_CREATE_CROWD_TOKEN, common.EVENT_TYPE_CREATE_MANAGE_TOKEN}
	}
	start, limit := model.PageLimit(pageInfo)
	sql := "select property_id,property_name,property_category,property_subcategory,property_data->'$.url' as property_service_url,issuer,txes.tx_type as tx_type " +
		"from smart_properties as sp join txes" +
		" on sp.create_tx_id=txes.tx_id " +
		"where txes.tx_type IN (?) order by sp.id limit ?,?"
	rows, err := db.Raw(sql, propertyType, start, limit).Rows()
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = nil
		}
		return nil, err
	}
	items := make([]PropertyObj, 0)
	defer rows.Close()
	for rows.Next() {
		var item PropertyObj
		db.ScanRows(rows, &item)
		item.PropertyServiceURL = strings.Trim(item.PropertyServiceURL, "\"")
		items = append(items, item)
	}
	return items, nil
}

func CountProperties(propertyType []int) int {
	if len(propertyType) == 0 {
		propertyType = []int{common.EVENT_TYPE_CREATE_FIXED_TOKEN, common.EVENT_TYPE_CREATE_CROWD_TOKEN, common.EVENT_TYPE_CREATE_MANAGE_TOKEN}
	}
	var count int
	sql := "select count(1) " +
		"from smart_properties as sp join txes" +
		" on sp.create_tx_id=txes.tx_id " +
		"where txes.tx_type IN (?)"
	db.Raw(sql, propertyType).Count(&count)
	return count
}

func PropertyDetail(propertyId int64) (*util.Result, error) {
	sql := "select property_data, txes.tx_type as tx_type from  smart_properties as sp join txes on sp.create_tx_id=txes.tx_id and sp.property_id=?"
	row := db.Raw(sql, propertyId).Row()
	var propertyData string
	var txType string
	row.Scan(&propertyData, &txType)
	if len(propertyData) > 0 {
		ret := util.JsonStringToMap(propertyData)
		(*ret)["txType"] = txType
		return ret, nil
	}
	return nil, errors.New("property not found")

}

type ChangeLog struct {
	BlockTime          int64 `json:"blockTime"`
	TxType             string    `json:"txType"`
	TxHash             string    `json:"txHash"`
}

func ListPropertyChangeLog(propertyId int64, pageInfo *model.PageInfo) ([]ChangeLog, error) {
	sql := "SELECT txes.block_time, txes.tx_type as tx_type, txes.tx_hash as tx_hash FROM property_histories as h join txes  on h.tx_id=txes.tx_id where h.property_id=? and txes.tx_type in (50,51,53,54,56,70) order by h.id desc limit ?,?"
	start, limit := model.PageLimit(pageInfo)
	rows, err := db.Raw(sql, propertyId, start, limit).Rows()
	items := make([]ChangeLog, 0)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = nil
		}
		return items, err
	}

	defer rows.Close()
	for rows.Next() {
		var item ChangeLog
		db.ScanRows(rows, &item)
		items = append(items, item)
	}
	return items, nil
}
func CountPropertyChangeLog(propertyId int64) int {
	sql := "SELECT count(1) FROM property_histories as h join txes  on h.tx_id=txes.tx_id where h.property_id=? and txes.tx_type in (50,51,53,54,56,70)"
	ret := 0
	db.Raw(sql, propertyId).Count(&ret)
	return ret
}

func CountPropertyTxes(propertyId int64, eventTypes int) int {
	eventTypesArray := []int{eventTypes}
	if eventTypes == -1 {
		eventTypesArray = []int{common.EVENT_TYPE_SIMPLE_SEND, common.EVENT_TYPE_SEND_ALL, common.EVENT_TYPE_GRANT_TOKEN, common.EVENT_TYPE_BUY_CROWD_TOKEN, common.EVENT_TYPE_AIR_DROP}
	}
	sql := "SELECT count(1) FROM property_histories  join txes  on property_histories.tx_id=txes.tx_id where property_histories.property_id=? and txes.tx_type in (?)"
	var count int
	db.Raw(sql, propertyId, eventTypesArray).Count(&count)
	return count
}

func ListPropertyTxes(propertyId int64, eventTypes int, pageInfo *model.PageInfo) ([]*util.Result, error) {
	eventTypesArray := []int{eventTypes}
	if eventTypes == -1 {
		eventTypesArray = []int{common.EVENT_TYPE_SIMPLE_SEND, common.EVENT_TYPE_SEND_ALL, common.EVENT_TYPE_GRANT_TOKEN, common.EVENT_TYPE_BUY_CROWD_TOKEN, common.EVENT_TYPE_AIR_DROP}
	}
	start, limit := model.PageLimit(pageInfo)
	sql := "SELECT tx_jsons.tx_data FROM property_histories  join txes  on property_histories.tx_id=txes.tx_id join tx_jsons on txes.tx_id=tx_jsons.tx_id where property_histories.property_id=? and txes.tx_type in (?) order by property_histories.id desc limit ?,?"
	rows, err := db.Raw(sql, propertyId, eventTypesArray, start, limit).Rows()
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = nil
		}
		return nil, err
	}
	items := make([]string, 0)
	defer rows.Close()
	for rows.Next() {
		var item string
		rows.Scan(&item)
		items = append(items, item)
	}
	re := make([]*util.Result, 0, len(items))
	for _, item := range items {
		jsonItem := util.JsonStringToMap(item)
		re = append(re, jsonItem)
	}
	return re, nil
}

func GetPropertyName(pid uint64) (string, error) {
	// The pending created issuance transaction, the property name is not set;
	// The invalid created issuance transaction, the property id = 0;
	// Bitcoin cash transaction in addresses_in_txes table, the property id = 0;
	if pid == 0 {
		return "", nil
	}

	var sp model.SmartProperty
	err := db.Select("property_name").Where("property_id = ?", pid).First(&sp).Error
	if err != nil {
		return "", err
	}

	return sp.PropertyName, nil
}

func ListOwnersCount(propertyId int) (int64, error) {
	var total int64
	var err error
	err = db.Table("address_balances").
		Where("property_id = ?", propertyId).
		Count(&total).Error

	if err != nil {
		return -1, err
	}
	return total, nil

}

func ListOwners(pageSize int, pageNo int, propertyId int) ([]view.AddressBalance, error) {
	rows, err := db.Table("address_balances adr").Select("adr.address,adr.balance_available,if(tx.tx_type=185,1,0) as status").
		Where("adr.property_id = ?", propertyId).
		Joins("JOIN txes tx on adr.last_tx_id = tx.tx_id ").
		Limit(pageSize).
		Offset(pageSize*pageNo - pageSize).
		Rows()

	models := make([]view.AddressBalance, 0)
	for rows.Next() {
		var vo view.AddressBalance
		err = db.ScanRows(rows, &vo)
		if err != nil {
			return nil, err
		}

		models = append(models, vo)
	}

	sql := "SELECT count(1) FROM property_histories as h join txes  on h.tx_id=txes.tx_id where h.property_id=? and txes.tx_type in (50,51,53,54,56,70)"
	ret := 0
	db.Raw(sql, propertyId).Count(&ret)

	return models, nil
}

// support query via property name or property id input
func GetPropertyByKeyword(keyword string, pageInfo *model.PageInfo) ([]PropertyObj, error) {
	propertyType := []int{common.EVENT_TYPE_CREATE_FIXED_TOKEN, common.EVENT_TYPE_CREATE_CROWD_TOKEN, common.EVENT_TYPE_CREATE_MANAGE_TOKEN}

	start, limit := model.PageLimit(pageInfo)
	sql := "select property_id,property_name,property_category,property_subcategory,property_data->'$.url' as property_service_url,issuer,txes.tx_type as tx_type " +
		"from smart_properties as sp join txes" +
		" on sp.create_tx_id=txes.tx_id " +
		"where txes.tx_type IN (?) and ((property_name LIKE ?) OR (property_id = ?)) order by sp.id limit ?,?;"
	rows, err := db.Raw(sql, propertyType, keyword+"%", keyword, start, limit).Rows()

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = nil
		}
		return nil, err
	}

	items := make([]PropertyObj, 0)
	defer rows.Close()
	for rows.Next() {
		var item PropertyObj
		db.ScanRows(rows, &item)
		item.PropertyServiceURL = strings.Trim(item.PropertyServiceURL, "\"")
		items = append(items, item)
	}
	return items, nil
}