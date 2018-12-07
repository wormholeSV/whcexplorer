package model

import (
	common "github.com/copernet/whccommon/model"
	"github.com/jinzhu/gorm"
)

func GetBlockCounts() (int, error) {
	var count int
	err := db.Table("blocks").Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func FetchBlockList(pageSize, page int) ([]common.Block, error) {
	rows, err := db.Table("blocks").
		Order("block_height DESC").
		Offset(pageSize * (page - 1)).
		Limit(pageSize).
		Rows()

	if err != nil {
		return nil, err
	}

	var blocks []common.Block
	for rows.Next() {
		var block common.Block
		err = db.ScanRows(rows, &block)
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, block)
	}

	return blocks, nil
}

func GetBlockListInperiodCount(from, to uint32) (int, error) {
	var count int
	err := db.Table("blocks").
		Where("block_time between ? and ?", from, to).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func FetchBlockListInperiod(pageSize, page int, from, to uint32) ([]common.Block, error) {
	rows, err := db.Table("blocks").
		Where("block_time between ? and ?", from, to).
		Order("block_height DESC").
		Offset(pageSize * (page - 1)).
		Limit(pageSize).
		Rows()

	if err != nil {
		return nil, err
	}

	var blocks []common.Block
	for rows.Next() {
		var block common.Block
		err = db.ScanRows(rows, &block)
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, block)
	}

	return blocks, nil
}

func GetBlockByHash(hash string) (*common.Block, error) {
	var block common.Block
	err := db.Table("blocks").
		Where("block_hash = ?", hash).
		First(&block).Error

	if err != nil {
		return nil, err
	}

	return &block, nil
}

func GetBlockByHeight(height int32) (*common.Block, error) {
	var block common.Block
	err := db.Table("blocks").
		Where("block_height = ?", height).
		First(&block).Error

	if err != nil {
		return nil, err
	}

	return &block, nil
}

func GetLastBlock() *common.Block {
	var block = common.Block{}
	err := db.Last(&block).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}

	return &block
}