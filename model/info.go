package model

import (
	"time"

	"github.com/copernet/whccommon/model"
)

func GetInfoForNetwork() (*model.Block, error) {
	var block model.Block
	err := db.Table("blocks").
		Select("block_height, block_time, whccount").
		Last(&block).Error

	if err != nil {
		return nil, err
	}

	return &block, nil
}

func GetTotalWormholeTxCount() (int, error) {
	var count int
	row := db.Table("blocks").
		Select("sum(whccount) as count").Row()

	err := row.Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetWormholeTxCountInPeriod(from int64) (int, error) {
	now := time.Now().Unix()
	var count int
	row := db.Table("blocks").
		Select("sum(whccount) as count").Where("block_time between ? and ?", from, now).
		Row()

	err := row.Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
