package api

import (
	"strconv"

	"github.com/copernet/whccommon/log"
	"github.com/copernet/whcexplorer/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Search supports block or transaction search, the parameter
// may be block height, block hash or transaction hash.
func Search(c *gin.Context) {
	keyword := c.Param("keyword")
	t, err := paramType(keyword)
	if err != nil {
		log.WithCtx(c).Errorf("user input parameter invalid： %v", err)
		c.JSON(200, apiError(ErrInvalidParam))
		return
	}

	switch t {
	case BlockHeight:
		// the function paramType() has check the parameter is a valid block height,
		// so here ignore the returned error
		height, _ := strconv.Atoi(keyword)
		block, err := model.GetBlockByHeight(int32(height))
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				log.WithCtx(c).Errorf("fetch block information by block height failed: %v", err)
				c.JSON(200, apiError(ErrNotFoundBlockHeight))
				return
			}

			log.WithCtx(c).Errorf("fetch block information by block height failed: %v", err)
			c.JSON(200, apiError(ErrFetchBlockByHeight))
			return
		}

		currentHeight := model.GetLastBlock().BlockHeight
		c.JSON(200, apiSuccess(formatBlockForAPi(block, currentHeight)))
		return

	case BlockHash:
		block, err := model.GetBlockByHash(keyword)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				log.WithCtx(c).Errorf("fetch block information by block hash failed: %v", err)
				c.JSON(200, apiError(ErrNotFoundBlockHash))
				return
			}

			log.WithCtx(c).Errorf("fetch block information by block hash failed: %v", err)
			c.JSON(200, apiError(ErrFetchBlockByHash))
			return
		}

		currentHeight := model.GetLastBlock().BlockHeight
		c.JSON(200, apiSuccess(formatBlockForAPi(block, currentHeight)))
		return

	case TxHash:
		tx, rawdata, err := model.GetTxByHash(keyword)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				c.JSON(200, apiError(ErrNotFoundTransactionHash))
				return
			}

			c.JSON(200, apiError(ErrFetchTransactionByHash))
			return
		}

		currentHeight := model.GetLastBlock().BlockHeight
		ret, err := formatTransactionForAPI(tx, currentHeight, c)
		if err != nil {
			return
		}
		ret.RawData = *rawdata

		c.JSON(200, apiSuccess(ret))
		return

	default:
		c.JSON(200, apiError(ErrInvalidParam))
		return
	}

	c.JSON(200, errNotFound)
}
