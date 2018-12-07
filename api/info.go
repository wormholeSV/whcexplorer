package api

import (
	"time"

	"github.com/copernet/whcexplorer/model"
	"github.com/gin-gonic/gin"
)

const (
	oneDay = 24 * 60 * 60
)

type TipInfo struct {
	BlockHeight     int32 `json:"block_height"`
	BlockTime       int64 `json:"block_time"`
	CurrentWhcCount int   `json:"current_whc_count"`
	OneDayWhcCount  int   `json:"one_day_whc_count"`
	TotalWhcCount   int   `json:"total_whc_count"`
}

func Info(c *gin.Context) {
	tip, err := model.GetInfoForNetwork()
	if err != nil {
		c.JSON(200, apiError(ErrGetNetworkInfo))
		return
	}

	total, err := model.GetTotalWormholeTxCount()
	if err != nil {
		c.JSON(200, apiError(ErrGetTotalWormholeTxCount))
		return
	}

	dayCount, err := model.GetWormholeTxCountInPeriod(time.Now().Unix() - oneDay)
	if err != nil {
		c.JSON(200, apiError(ErrGetWormholeTxCountInPeriod))
		return
	}

	info := TipInfo{
		BlockHeight:     int32(tip.BlockHeight),
		BlockTime:       tip.BlockTime,
		CurrentWhcCount: tip.Whccount,
		OneDayWhcCount:  dayCount,
		TotalWhcCount:   total,
	}

	c.JSON(200, apiSuccess(&info))
}
