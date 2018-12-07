package api

import (
	"fmt"
	common "github.com/copernet/whccommon/model"
	"github.com/copernet/whcexplorer/config"
	"github.com/copernet/whcexplorer/model"
	"github.com/copernet/whcexplorer/model/view"
	"github.com/gin-gonic/gin"
)

func GetBurnList(c *gin.Context) {
	total, err := model.GetBurnCount()
	if err != nil {
		c.JSON(200, apiError(ErrFetchBurnListCount))
		return
	}

	paginator := common.GetPageInfo(c)
	if total == 0 {
		c.JSON(200, apiSuccess(common.EmptyPageInfo(paginator)))
		return
	}

	list, err := model.GetBurnList(paginator.PageSize, paginator.Page)
	if err != nil {
		c.JSON(200, apiError(ErrFetchBurnListCount))
		return
	}

	block := model.GetLastBlock()
	for _, info := range list {
		generateProcess(info, block)
	}

	paginator.Total = total
	paginator.List = list

	c.JSON(200, apiSuccess(paginator))
}

func Summary(c *gin.Context) {
	summary, err := model.GetBurnSummary()
	if err != nil {
		c.JSON(200, apiError(ErrFetchBurnListCount))
		return
	}

	c.JSON(200, apiSuccess(summary))
}

func generateProcess(info *view.BurnInfo, block *common.Block) {
	diff := int(block.BlockHeight) - int(info.TxBlockHeight) + 1
	confirms := config.GetConfirms()

	if diff >= confirms {
		info.Process = fmt.Sprintf("%d/%d", confirms, confirms)
	} else {
		info.MatureTime = info.BlockTime + (confirms-diff)*10*60
		info.Process = fmt.Sprintf("%d/%d", diff, confirms)
	}
}
