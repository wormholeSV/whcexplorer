package model

import (
	"github.com/shopspring/decimal"
	"github.com/copernet/whccommon/model"
	"github.com/copernet/whcexplorer/api/common"
	"reflect"
	"github.com/jinzhu/gorm"
	"encoding/json"
	"github.com/copernet/whc.go/btcjson"
)


type wchBalance struct {
	Balance_available float64	`json:"balanceAvailable"`
	Balance_total float64	`json:"balanceTotal"`
}
 func WhcBalance(address string)(decimal.Decimal, decimal.Decimal,error){
 	sql := "SELECT balance_available,(balance_available+balance_frozen) as balance_total FROM address_balances where property_id=1 and address=?"
 	var item wchBalance
 	err := db.Raw(sql, address).Scan(&item).Error
 	if err == nil {
 		return decimal.NewFromFloat(item.Balance_available), decimal.NewFromFloat(item.Balance_total), nil
	}
 	if gorm.IsRecordNotFoundError(err) {
 		err = nil
	}
 	return decimal.Zero, decimal.Zero, err
 }
type wchStatistics struct {
	Amount float64	`json:"amount"`
	AddressRole string	`json:"addressRole"`
}
func WhcStatistics(address string) (decimal.Decimal, decimal.Decimal,decimal.Decimal,error) {
	sql := "SELECT sum(balance_available_credit_debit) as amount, address_role FROM " +
		"addresses_in_txes join txes " +
		"on addresses_in_txes.tx_id=txes.tx_id  " +
		"where txes.tx_state='valid' " +
		"and address=? and property_id=1 group by address_role"
	rows,err := db.Raw(sql, address).Rows()

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = nil
		}
		return decimal.Zero, decimal.Zero, decimal.Zero, err
	}
	items := make([]wchStatistics,0)
	defer rows.Close()
	for rows.Next() {
		var item wchStatistics
		db.ScanRows(rows, &item)
		items = append(items, item)
	}
	var cosumed, sended, received decimal.Decimal
	for _, item := range items {
		role := model.AddressRole(item.AddressRole)
		if  role == model.Sender {
			sended = sended.Add(decimal.NewFromFloat(item.Amount))
		}
		if  role == model.Recipient || role == model.Payee || role == model.Buyer {
			received = received.Add(decimal.NewFromFloat(item.Amount))
		}
		if  role == model.Feepayer {
			cosumed = cosumed.Add(decimal.NewFromFloat(item.Amount))
		}
	}
	return cosumed, sended, received, nil

}

func WhcTxCountStatistics(address string) (int, error) {
	sql := "SELECT count(1) FROM " +
		"addresses_in_txes join txes " +
		"on addresses_in_txes.tx_id=txes.tx_id  " +
		"where txes.tx_state='valid' " +
		"and address=? and property_id=1"
	var count int
	db.Raw(sql, address).Count(&count)
	return count, nil
}
type AddressTxInfo struct {
	TxHash string	`json:"txHash"`
	TxType int		`json:"txType"`
	TxBlockHeight int	`json:"txBlockHeight"`
	AddressRole string	`json:"addressRole"`
	BalanceAvailableCreditDebit  float64	`json:"balanceAvailableCreditDebit"`
	BalanceFrozenCreditDebit float64	`json:"balanceFrozenCreditDebit"`
	PropertyId int	`json:"propertyId"`
}
func ListAddressTxes(address string, txType int, times []int64, propertyId int64, pageInfo *model.PageInfo) ([]btcjson.GenerateTransactionResult, []int, error){
	values := make([]interface{}, 0)
	values = append(values, address)
	start, limit := model.PageLimit(pageInfo)
	var byEventTypeQuery, byTimeQuery, byPropertyIdQuery string
	if txType != -1 {
		byEventTypeQuery = "and txes.tx_type=? "
		values = append(values, txType)
	}
	if len(times) ==2 {
		byTimeQuery = " and txes.block_time between ? and ? "
		values = append(values, times[0],times[1])
	}
	if propertyId != -1 {
		byPropertyIdQuery = "and addresses_in_txes.property_id=? "
		values = append(values, propertyId)
	}
	values = append(values, start, limit)
	//sql := "select " +
	//	"txes.tx_hash, txes.tx_type, txes.tx_block_height, addresses_in_txes.address_role, " +
	//	"addresses_in_txes.balance_available_credit_debit, addresses_in_txes.balance_frozen_credit_debit, addresses_in_txes.property_id " +
	//	"from addresses_in_txes join txes on addresses_in_txes.tx_id=txes.tx_id " +
	//	"where txes.tx_state='valid' and addresses_in_txes.address=? " +
	//	byEventTypeQuery +
	//	byTimeQuery +
	//	byPropertyIdQuery +
	//	"order by addresses_in_txes.id desc limit ?,?"

	sql := "select tx_jsons.tx_data, sign(addresses_in_txes.balance_available_credit_debit)  "+
		"from addresses_in_txes join txes on addresses_in_txes.tx_id=txes.tx_id " +
		" join tx_jsons on txes.tx_id=tx_jsons.tx_id " +
		"where txes.tx_state='valid' and addresses_in_txes.address=? " +
		byEventTypeQuery +
		byTimeQuery +
		byPropertyIdQuery +
		"order by addresses_in_txes.id desc limit ?,?"


	args := make([]reflect.Value, 0)
	args = append(args, reflect.ValueOf(sql))
	for _, v := range values {
		args = append(args, reflect.ValueOf(v))
	}
	ret := reflect.ValueOf(db).MethodByName("Raw").Call(args)
	rdb, _ := ret[0].Interface().(*gorm.DB)
	rows, err := rdb.Rows()
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = nil
		}
		return nil, nil, err
	}
	//items := make([]AddressTxInfo,0)
	//defer rows.Close()
	//for rows.Next() {
	//	var item AddressTxInfo
	//	db.ScanRows(rows, &item)
	//	items = append(items, item)
	//}
	//return items, nil

	items := make([]btcjson.GenerateTransactionResult, 0)
	directions := make([]int, 0)
	defer rows.Close()
	for rows.Next() {
		var itemStr string
		var direction int
		rows.Scan(&itemStr, &direction)
		var tx btcjson.GenerateTransactionResult
		err = json.Unmarshal([]byte(itemStr), &tx)
		if err != nil {
			return nil, nil, err
		}
		items = append(items, tx)
		directions = append(directions, direction)
	}
	return items, directions, nil

}

func CountAddressTxes(address string, txType int, times []int64, propertyId int64) (int){
	values := make([]interface{}, 0)
	values = append(values, address)
	var byEventTypeQuery, byTimeQuery, byPropertyIdQuery string
	if txType != -1 {
		byEventTypeQuery = "and txes.tx_type=? "
		values = append(values, txType)
	}
	if len(times) ==2 {
		byTimeQuery = " and txes.block_time between ? and ? "
		values = append(values, times[0],times[1])
	}
	if propertyId != -1 {
		byPropertyIdQuery = "and addresses_in_txes.property_id=? "
		values = append(values, propertyId)
	}
	sql := "select count(1) " +
		"from addresses_in_txes join txes on addresses_in_txes.tx_id=txes.tx_id " +
		"where txes.tx_state='valid' and addresses_in_txes.address=? " +
		byEventTypeQuery +
		byTimeQuery +
		byPropertyIdQuery

	args := make([]reflect.Value, 0)
	args = append(args, reflect.ValueOf(sql))
	for _, v := range values {
		args = append(args, reflect.ValueOf(v))
	}
	ret := reflect.ValueOf(db).MethodByName("Raw").Call(args)
	rdb, _ := ret[0].Interface().(*gorm.DB)
	var count int
	rdb.Count(&count)
	return count
}

type AddressProperty struct {
	PropertyId int64	`json:"propertyId"`
	PropertyName string	`json:"propertyName"`
	TxType int	`json:"txType"`
	BalanceAvailable float64	`json:"balanceAvailable"`
}
func ListAddressProperties(address string, propertyType int, pageInfo *model.PageInfo) ([]AddressProperty, error) {
	sql := "SELECT smart_properties.property_id as property_id,smart_properties.property_name as property_name,txes.tx_type as tx_type, address_balances.balance_available as balance_available FROM " +
		"address_balances join smart_properties on address_balances.property_id=smart_properties.property_id " +
		"join txes on smart_properties.create_tx_id=txes.tx_id " +
		"where smart_properties.property_id !=1 and address_balances.address=? and txes.tx_type in (?) order by address_balances.id limit ?,?"
	txTypeArray := []int{common.EVENT_TYPE_CREATE_FIXED_TOKEN, common.EVENT_TYPE_CREATE_CROWD_TOKEN, common.EVENT_TYPE_CREATE_MANAGE_TOKEN}
	if propertyType != -1 {
		txTypeArray = []int{propertyType}
	}
	start, limit := model.PageLimit(pageInfo)
	rows, err := db.Raw(sql, address, txTypeArray, start, limit).Rows()

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = nil
		}
		return nil, err
	}
	items := make([]AddressProperty,0)
	defer rows.Close()
	for rows.Next() {
		var item AddressProperty
		db.ScanRows(rows, &item)
		items = append(items, item)
	}
	return items, nil
}

func CountAddressProperties(address string, propertyType int) (int) {
	sql := "SELECT count(1) FROM " +
		"address_balances join smart_properties on address_balances.property_id=smart_properties.property_id " +
		"join txes on smart_properties.create_tx_id=txes.tx_id " +
		"where smart_properties.property_id !=1 and address_balances.address=? and txes.tx_type in (?)"
	txTypeArray := []int{common.EVENT_TYPE_CREATE_FIXED_TOKEN, common.EVENT_TYPE_CREATE_CROWD_TOKEN, common.EVENT_TYPE_CREATE_MANAGE_TOKEN}
	if propertyType != -1 {
		txTypeArray = []int{propertyType}
	}
	var count int
	db.Raw(sql, address, txTypeArray).Count(&count)
	return count
}
