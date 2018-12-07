package api

import (
	"github.com/copernet/whcexplorer/model/view"
	"math/big"
	"strconv"

	"github.com/bcext/gcash/blockchain"
	"github.com/bcext/gcash/chaincfg"
	"github.com/copernet/whccommon/log"
	common "github.com/copernet/whccommon/model"
	"github.com/copernet/whcexplorer/config"
	"github.com/copernet/whcexplorer/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type BlockAPI struct {
	*common.Block

	NonceString   string               `json:"nonce_str"`
	BitsString    string               `json:"bits_str"`
	Difficulty    float64              `json:"difficulty"`
	Confirmations int32                `json:"confirmations"`
	FeeRates      *view.FeeRateSummary `json:"fee_rates"`
}

func ListBlocks(c *gin.Context) {
	if c.Query("from") != "" {
		BlocksInPeriod(c)
		return
	}

	total, err := model.GetBlockCounts()
	if err != nil {
		c.JSON(200, apiError(ErrFetchBlockListCount))
		return
	}

	paginator := common.GetPageInfo(c)
	if total == 0 {
		c.JSON(200, apiSuccess(common.EmptyPageInfo(paginator)))
		return
	}

	blocks, err := model.FetchBlockList(paginator.PageSize, paginator.Page)
	if err != nil {
		c.JSON(200, apiError(ErrFetchBlockList))
		return
	}

	paginator.Total = total
	currentHeight := model.GetLastBlock().BlockHeight
	paginator.List = formatBlocksForAPi(blocks, currentHeight)
	c.JSON(200, apiSuccess(paginator))
}

// BlocksInPeriod returns block list in the specified block time period.
func BlocksInPeriod(c *gin.Context) {
	paramFrom := c.Query("from")
	from, ok := checkBlockTime(paramFrom)
	if !ok {
		c.JSON(200, apiError(ErrInvalidBlockTimestamp))
		return
	}
	paramTo := c.Query("to")
	to, ok := checkBlockTime(paramTo)
	if !ok {
		c.JSON(200, apiError(ErrInvalidBlockTimestamp))
		return
	}

	paginator := common.GetPageInfo(c)
	total, err := model.GetBlockListInperiodCount(from, to)
	if total == 0 {
		c.JSON(200, apiSuccess(common.EmptyPageInfo(paginator)))
		return
	}

	blocks, err := model.FetchBlockListInperiod(paginator.PageSize, paginator.Page, from, to)
	if err != nil {
		c.JSON(200, apiError(ErrFetchBlockList))
		return
	}

	paginator.Total = total
	currentHeight := model.GetLastBlock().BlockHeight
	paginator.List = formatBlocksForAPi(blocks, currentHeight)
	c.JSON(200, apiSuccess(paginator))
}

// GetBlock return a block information detail according to the specified
// block height or block hash
func GetBlock(c *gin.Context) {
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

		ret := transferBlock(model.GetLastBlock().BlockHeight, block)
		c.JSON(200, apiSuccess(ret))
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
		ret := transferBlock(model.GetLastBlock().BlockHeight, block)
		c.JSON(200, apiSuccess(ret))
		return
	case TxHash:
		// this API does not accept txhash
		c.JSON(200, apiError(ErrInvalidBlockHashOrHeight))
		return
	default:
		c.JSON(200, apiError(ErrInvalidParam))
		return
	}
}
func transferBlock(height int64, block *common.Block) *BlockAPI {
	ret := formatBlockForAPi(block, height)

	summary, _ := model.GetFeeRateSummary(block.BlockHeight)
	ret.FeeRates = summary

	return ret
}

// getDifficultyRatio returns the proof-of-work difficulty as a multiple of the
// minimum difficulty using the passed bits field from the header of a block.
func getDifficultyRatio(bits uint32, params *chaincfg.Params) float64 {
	// The minimum difficulty is the max possible proof-of-work limit bits
	// converted back to a number.  Note this is not the same as the proof of
	// work limit directly because the block difficulty is encoded in a block
	// with the compact form which loses precision.
	max := blockchain.CompactToBig(params.PowLimitBits)
	target := blockchain.CompactToBig(bits)

	difficulty := new(big.Rat).SetFrac(max, target)
	outString := difficulty.FloatString(8)
	diff, err := strconv.ParseFloat(outString, 64)
	if err != nil {
		return 0
	}
	return diff
}

func formatBlockForAPi(block *common.Block, height int64) *BlockAPI {
	if block == nil {
		return nil
	}

	var apiRet BlockAPI
	apiRet.Block = block
	apiRet.NonceString = "0x" + strconv.FormatInt(int64(block.Nonce), 16)
	apiRet.BitsString = "0x" + strconv.FormatInt(int64(block.Bits), 16)
	apiRet.Difficulty = getDifficultyRatio(block.Bits, config.GetChainParam())
	apiRet.Confirmations = int32(height-block.BlockHeight) + 1

	return &apiRet
}

func formatBlocksForAPi(blocks []common.Block, height int64) []BlockAPI {
	if blocks == nil {
		return nil
	}

	ret := make([]BlockAPI, 0, len(blocks))
	for idx, _ := range blocks {
		item := formatBlockForAPi(&blocks[idx], height)
		ret = append(ret, *item)
	}

	return ret
}
