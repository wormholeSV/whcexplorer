package model

import (
	"encoding/json"
	"github.com/copernet/whc.go/btcjson"
	"github.com/copernet/whccommon/model"
	"github.com/copernet/whcexplorer/config"
	"github.com/copernet/whcexplorer/model/view"
)

func GetTxByHash(hash string) (*btcjson.GenerateTransactionResult, *string, error) {
	var txj model.TxJson
	err := db.Table("tx_jsons").Select("tx_jsons.tx_data, tx_jsons.raw_data").
		Joins("INNER JOIN txes ON txes.tx_id = tx_jsons.tx_id AND txes.tx_hash = ? and txes.protocol = ?", hash, model.Wormhole).
		Order("tx_jsons.id").First(&txj).Error

	if err != nil {
		return nil, nil, err
	}

	var tx btcjson.GenerateTransactionResult
	err = json.Unmarshal([]byte(txj.TxData), &tx)
	if err != nil {
		return nil, nil, err
	}

	return &tx, &txj.RawData, nil
}

func GetTxByBlockHeightCount(height int32) (int, error) {
	var tx model.Block
	err := db.Table("blocks").Select("whccount").Where("block_height = ?", height).First(&tx).Error
	if err != nil {
		return 0, err
	}

	return tx.Whccount, nil
}

func GetTxByBlockHeightCountTxtypeLimited(height int32, txType int) (int, error) {
	var count int
	err := db.Table("blocks").
		Where("block_height = ?", height).
		Joins("INNER JOIN txes ON blocks.block_height = txes.tx_block_height AND txes.tx_type = ?", txType).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetTxByBlockHeight(height int32, pageNo, pageSize int) ([]btcjson.GenerateTransactionResult, error) {
	rows, err := db.Table("tx_jsons").Select("tx_jsons.tx_data").
		Joins("INNER JOIN txes ON txes.tx_id = tx_jsons.tx_id AND txes.tx_block_height = ?", height).
		Order("tx_jsons.id").
		Offset(pageNo*pageSize - pageSize).
		Limit(pageSize).
		Rows()

	if err != nil {
		return nil, err
	}

	list := make([]btcjson.GenerateTransactionResult, 0)
	for rows.Next() {
		var txj model.TxJson
		err = db.ScanRows(rows, &txj)
		if err != nil {
			return nil, err
		}

		var tx btcjson.GenerateTransactionResult
		err = json.Unmarshal([]byte(txj.TxData), &tx)
		if err != nil {
			return nil, err
		}

		list = append(list, tx)
	}

	return list, nil
}

func GetTxByBlockHeightTxtypeLimited(height int32, txType, pageNo, pageSize int) ([]btcjson.GenerateTransactionResult, error) {
	rows, err := db.Table("tx_jsons").
		Select("tx_jsons.tx_data").
		Joins("INNER JOIN txes ON txes.tx_id = tx_jsons.tx_id AND txes.tx_block_height = ? AND txes.tx_type = ?", height, txType).
		Order("tx_jsons.id").
		Offset(pageNo*pageSize - pageSize).
		Limit(pageSize).
		Rows()

	if err != nil {
		return nil, err
	}

	list := make([]btcjson.GenerateTransactionResult, 0)
	for rows.Next() {
		var txj model.TxJson
		err = db.ScanRows(rows, &txj)
		if err != nil {
			return nil, err
		}

		var tx btcjson.GenerateTransactionResult
		err = json.Unmarshal([]byte(txj.TxData), &tx)
		if err != nil {
			return nil, err
		}

		list = append(list, tx)
	}

	return list, nil
}

func GetTxByBlockHashCount(hash string) (int, error) {
	var tx model.Block
	err := db.Table("blocks").Select("whccount").Where("block_hash = ?", hash).First(&tx).Error
	if err != nil {
		return 0, err
	}

	return tx.Whccount, nil
}

func GetTxByBlockHashCountTxtypeLimited(hash string, txType int) (int, error) {
	var count int
	err := db.Table("blocks").
		Where("block_hash = ?", hash).
		Joins("INNER JOIN txes ON blocks.block_height = txes.tx_block_height AND txes.tx_type = ?", txType).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetTxByBlockHash(hash string, pageNo, pageSize int) ([]btcjson.GenerateTransactionResult, error) {
	rows, err := db.Table("tx_jsons").Select("tx_jsons.tx_data").
		Joins("INNER JOIN txes ON txes.tx_id = tx_jsons.tx_id").
		Joins("INNER JOIN blocks ON txes.tx_block_height = blocks.block_height AND blocks.block_hash = ?", hash).
		Order("tx_jsons.id").
		Offset(pageNo*pageSize - pageSize).
		Limit(pageSize).
		Rows()

	if err != nil {
		return nil, err
	}

	list := make([]btcjson.GenerateTransactionResult, 0)
	for rows.Next() {
		var txj model.TxJson
		err = db.ScanRows(rows, &txj)
		if err != nil {
			return nil, err
		}

		var tx btcjson.GenerateTransactionResult
		err = json.Unmarshal([]byte(txj.TxData), &tx)
		if err != nil {
			return nil, err
		}

		list = append(list, tx)
	}

	return list, nil
}

func GetTxByBlockHashTxtypeLimited(hash string, txType, pageNo, pageSize int) ([]btcjson.GenerateTransactionResult, error) {
	rows, err := db.Table("tx_jsons").Select("tx_jsons.tx_data").
		Joins("INNER JOIN txes ON txes.tx_id = tx_jsons.tx_id AND txes.tx_type = ?", txType).
		Joins("INNER JOIN blocks ON txes.tx_block_height = blocks.block_height AND blocks.block_hash = ?", hash).
		Order("tx_jsons.id").
		Offset(pageNo*pageSize - pageSize).
		Limit(pageSize).
		Rows()

	if err != nil {
		return nil, err
	}

	list := make([]btcjson.GenerateTransactionResult, 0)
	for rows.Next() {
		var txj model.TxJson
		err = db.ScanRows(rows, &txj)
		if err != nil {
			return nil, err
		}

		var tx btcjson.GenerateTransactionResult
		err = json.Unmarshal([]byte(txj.TxData), &tx)
		if err != nil {
			return nil, err
		}

		list = append(list, tx)
	}

	return list, nil
}

func GetBurnCount() (int, error) {
	var count int
	err := db.Table("txes").Where("tx_state=? and tx_type=?", model.Valid, 68).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetBurnList(pageSize int, page int) ([]*view.BurnInfo, error) {
	rows, err := db.Table("txes").Select("txes.tx_hash,txes.block_time as block_time,adr.address,if(adr.balance_frozen_credit_debit = 0, " +
		"adr.balance_available_credit_debit, adr.balance_frozen_credit_debit) as whc," +
		"if(adr.balance_frozen_credit_debit = 0, adr.balance_available_credit_debit / 100,adr.balance_frozen_credit_debit / 100) as bch" +
		",txes.tx_block_height ,b.block_time as mature_time").
		Joins(" JOIN addresses_in_txes adr ON txes.tx_id = adr.tx_id  left join blocks b on  txes.tx_block_height+ ? = b.block_height "+
			"where tx_state=? and tx_type=?", config.GetConfirms(), model.Valid, 68).
		Offset(pageSize * (page - 1)).
		Order(" txes.tx_id DESC").
		Limit(pageSize).
		Rows()

	if err != nil {
		return nil, err
	}

	var burns []*view.BurnInfo
	for rows.Next() {
		var block view.BurnInfo
		err = db.ScanRows(rows, &block)
		if err != nil {
			return nil, err
		}

		burns = append(burns, &block)
	}

	return burns, nil

}

func GetBurnSummary() (view.BurnSummary, error) {
	var summary view.BurnSummary
	err := db.Table("txes").Select("SUM(IF(adr.balance_frozen_credit_debit = 0, adr.balance_available_credit_debit," +
		"adr.balance_frozen_credit_debit)) AS total,SUM(adr.balance_available_credit_debit) AS avail").
		Joins(" JOIN addresses_in_txes adr ON adr.tx_id = txes.tx_id").Where("tx_state=? and tx_type=?", model.Valid, 68).
		First(&summary).Error

	return summary, err
}

func GetLatestWormholeTxInfo(startTime int) ([]model.Block, error) {
	var wormholeInfo []model.Block
	err := db.Table("blocks").Select("whccount, block_time").Order("id").Where("block_time >= ? AND whccount > ?", startTime, 0).Find(&wormholeInfo).Error
	if err != nil {
		return nil, err
	}

	return wormholeInfo, nil
}

func GetFeeRateSummary(height int64) (*view.FeeRateSummary, error) {

	rows, err := db.Raw("select min(c.fee_rate) as min_fee_rate, CONVERT(avg(c.fee_rate) , DECIMAL(64, 8)) as avg_fee_rate, "+
		"max(c.fee_rate) as max_fee_rate from (select CONVERT(txj.tx_data -> '$.fee_rate', DECIMAL(64, 8)) as fee_rate "+
		"from txes tx join tx_jsons txj on tx.tx_id = txj.tx_id where tx.tx_block_height = ? ) c", height).Rows()

	if err != nil {
		return nil, err
	}

	var item view.FeeRateSummary
	for rows.Next() {
		db.ScanRows(rows, &item)
		break
	}

	if item.AvgFeeRate == nil || item.MaxFeeRate == nil || item.MinFeeRate == nil {
		return nil, nil
	}

	return &item, nil
}
