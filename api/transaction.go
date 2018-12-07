package api

import (
	"strconv"
	"time"

	"github.com/copernet/whc.go/btcjson"
	"github.com/copernet/whccommon/log"
	common "github.com/copernet/whccommon/model"
	"github.com/copernet/whcexplorer/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type TransactionDetail struct {
	TxHash              string              `json:"tx_hash"`
	TypeInt             uint64              `json:"type_int"`
	TypeStr             string              `json:"type_str"`
	SendingAddress      string              `json:"sending_address"`
	ReferenceAddress    interface{}         `json:"reference_address"` // only accept array
	Amount              string              `json:"amount"`
	BlockHeight         int32               `json:"block_height"`
	BlockTime           int64               `json:"block_time"`
	State               string              `json:"state"`
	PropertyId          int64               `json:"property_id"`
	PropertyName        string              `json:"property_name"`
	SendingPropertyName string              `json:"sending_property_name"`
	FeeRate             string              `json:"fee_rate,omitempty"`
	MinerFee            string              `json:"miner_fee"`
	WhcFee              string              `json:"whc_fee"`
	Recipients          []btcjson.Recipient `json:"recipients,omitempty"`
	Confirmations       int32               `json:"confirmations"`
	RawData             string              `json:"raw_data,omitempty"`
	Direction           int                 `json:"direction",omitempty`
}

// GetTransaction returns transaction detail according to the
// specified transaction hash.
func GetTransaction(c *gin.Context) {
	hash := c.Param("txhash")
	if !checkSha256Hash(hash) {
		c.JSON(200, apiError(ErrInvalidParam))
		return
	}

	txd, rawdata, err := model.GetTxByHash(hash)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(200, apiError(ErrNotFoundTransactionHash))
			return
		}

		c.JSON(200, apiError(ErrFetchTransactionByHash))
		return
	}

	currentHeight := model.GetLastBlock().BlockHeight
	ret, err := formatTransactionForAPI(txd, currentHeight, c)
	if err != nil {
		log.WithCtx(c).Errorf("format transaction failed: %v", err)
		return
	}

	ret.RawData = *rawdata

	c.JSON(200, apiSuccess(ret))
}

// GetTransactionList returns transaction list with brief information
// according to the specified block height or hash.
// The condition:
//   If block hash is not empty, the API will search transaction list in
//   the specified block hash.
//   If block hash is empty, next to search block height.
//   If the parameter is a invalid block hash or block height, error returns
//   immediately。
func GetTransactionList(c *gin.Context) {
	txTypeStr := c.DefaultQuery("txType", "-1")
	txType, err := strconv.Atoi(txTypeStr)
	if err != nil {
		c.JSON(200, apiError(ErrInvalidParam))
		return
	}

	var total int
	var txds []btcjson.GenerateTransactionResult
	paginator := common.GetPageInfo(c)

	paramHash := c.Query("block_hash")
	if paramHash != "" {
		ok := checkSha256Hash(paramHash)
		if !ok {
			c.JSON(200, apiError(ErrInvalidBlockHash))
			return
		}

		if txType == -1 {
			total, err = model.GetTxByBlockHashCount(paramHash)
			if err != nil {
				c.JSON(200, apiError(ErrGetTransactionListCount))
				return
			}

			if total == 0 {
				c.JSON(200, apiSuccess(common.EmptyPageInfo(paginator)))
				return
			}

			txds, err = model.GetTxByBlockHash(paramHash, paginator.Page, paginator.PageSize)
			if err != nil {
				c.JSON(200, apiError(ErrFetchTransactionListInBlockHash))
				return
			}
		} else {
			total, err = model.GetTxByBlockHashCountTxtypeLimited(paramHash, txType)
			if err != nil {
				c.JSON(200, apiError(ErrGetTransactionListCount))
				return
			}

			if total == 0 {
				c.JSON(200, apiSuccess(common.EmptyPageInfo(paginator)))
				return
			}

			txds, err = model.GetTxByBlockHashTxtypeLimited(paramHash, txType, paginator.Page, paginator.PageSize)
			if err != nil {
				c.JSON(200, apiError(ErrFetchTransactionListInBlockHash))
				return
			}
		}

		currentHeight := model.GetLastBlock().BlockHeight
		ret, err := formatTransactionsForAPI(txds, currentHeight, c)
		if err != nil {
			return
		}

		paginator.Total = total
		paginator.List = ret
		c.JSON(200, apiSuccess(paginator))
		return
	}

	// if the parameter is block_height
	param := c.Query("block_height")
	height, ok := checkBlockHeight(param)
	if !ok {
		c.JSON(200, apiError(ErrInvalidBlockHeight))
		return
	}

	if txType == -1 {
		total, err = model.GetTxByBlockHeightCount(height)
		if err != nil {
			c.JSON(200, apiError(ErrGetTransactionListCount))
			return
		}

		paginator := common.GetPageInfo(c)
		if total == 0 {
			c.JSON(200, apiSuccess(common.EmptyPageInfo(paginator)))
			return
		}

		txds, err = model.GetTxByBlockHeight(height, paginator.Page, paginator.PageSize)
		if err != nil {
			c.JSON(200, apiError(ErrFetchTransactionListInBlockHeight))
			return
		}
	} else {
		total, err = model.GetTxByBlockHeightCountTxtypeLimited(height, txType)
		if err != nil {
			c.JSON(200, apiError(ErrGetTransactionListCount))
			return
		}

		if total == 0 {
			c.JSON(200, apiSuccess(common.EmptyPageInfo(paginator)))
			return
		}

		txds, err = model.GetTxByBlockHeightTxtypeLimited(height, txType, paginator.Page, paginator.PageSize)
		if err != nil {
			c.JSON(200, apiError(ErrFetchTransactionListInBlockHeight))
			return
		}
	}

	currentHeight := model.GetLastBlock().BlockHeight
	ret, err := formatTransactionsForAPI(txds, currentHeight, c)
	if err != nil {
		return
	}

	paginator.Total = total
	paginator.List = ret
	c.JSON(200, apiSuccess(paginator))
}

func GetLatestInfo(c *gin.Context) {
	offsetStr := c.Query("timeOffset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(200, apiError(ErrInvalidInteger))
		return
	}

	utc, err := time.LoadLocation("UTC")
	if err != nil {
		c.JSON(200, apiError(ErrSetUTCLocation))
		return
	}

	now := time.Now().Unix()
	timeOffset := now - int64(offset*60)
	timeLocal := time.Unix(timeOffset, 0)
	hour, minute, second := timeLocal.In(utc).Clock()
	offsetSeconds := hour*3600 + minute*60 + second
	startTime := int(now) - 6*24*3600 - offsetSeconds

	countSlice, err := model.GetLatestWormholeTxInfo(startTime)
	if err != nil {
		c.JSON(200, apiError(ErrFetchLatestWhccountInfo))
		return
	}

	var count int
	dayCursor := 1
	ret := make(map[string]int)
	for _, item := range countSlice {
		if item.BlockTime < int64(startTime+dayCursor*24*3600-offset*60) {
			count += item.Whccount
		} else {
			ret[time.Unix(int64(startTime+dayCursor*24*3600-offset*60-1), 0).In(utc).Format("01-02")] = count
			dayCursor++
			count = item.Whccount
		}
	}

	ret[time.Unix(int64(startTime+7*24*3600-offset*60-1), 0).In(utc).Format("01-02")] = count

	c.JSON(200, apiSuccess(ret))
}

func formatTransactionsForAPI(txds []btcjson.GenerateTransactionResult, height int64, ctx *gin.Context) ([]TransactionDetail, error) {
	ret := make([]TransactionDetail, 0, len(txds))
	for _, tx := range txds {
		item, err := formatTransactionForAPI(&tx, height, ctx)
		if err != nil {
			return nil, err
		}

		ret = append(ret, *item)
	}

	return ret, nil
}

type ApiSubSends struct {
	btcjson.SubSend
	Address      string `json:"address"`
	PropertyName string `json:"property_name"`
}

func formatTransactionForAPI(txd *btcjson.GenerateTransactionResult, height int64, ctx *gin.Context) (*TransactionDetail, error) {
	ret := TransactionDetail{
		TxHash:              txd.TxID,
		TypeInt:             txd.TypeInt,
		TypeStr:             txd.Type,
		SendingAddress:      txd.SendingAddress,
		ReferenceAddress:    txd.ReferenceAddress,
		BlockHeight:         int32(txd.BlockHeight),
		BlockTime:           txd.BlockTime,
		State:               strconv.FormatBool(txd.Valid),
		PropertyId:          txd.PropertyID,
		PropertyName:        txd.PropertyName,
		MinerFee:            txd.Fee,
		WhcFee:              txd.TotalStoFee,
		FeeRate:             txd.FeeRate.String(),
		SendingPropertyName: txd.PropertyName,
		Confirmations:       int32(model.GetLastBlock().BlockHeight) - int32(txd.BlockHeight) + 1,
	}

	//pending
	if ret.BlockTime == 0 {
		ret.State = "pending"
		ret.Confirmations = 0
	}

	// format amount filed(remove appended '0')
	if txd.Amount != "" {
		amount, err := decimal.NewFromString(txd.Amount)
		if err != nil {
			return nil, err
		}

		ret.Amount = amount.String()
	}

	switch txd.TypeInt {
	case 1:
		ret.PropertyId = txd.PurchasedPropertyID
		ret.PropertyName = txd.PurchasedPropertyName

		ret.SendingAddress = txd.ReferenceAddress
		ret.ReferenceAddress = formatReferenceAddress(txd.SendingAddress)
		ret.WhcFee = ret.Amount

		tokens, _ := decimal.NewFromString(txd.PurchasedTokens)
		ret.Amount = tokens.String()
		ret.SendingPropertyName = "WHC"
	case 3:
		ret.ReferenceAddress = txd.Recipients
	case 4:
		subsends, err := formatSubsends(txd.SubSends, txd.ReferenceAddress)
		if err != nil {
			return nil, err
		}
		ret.ReferenceAddress = subsends
	case 50, 51, 54:
		ret.ReferenceAddress = formatReferenceAddressAmount(txd.SendingAddress, ret.Amount)
		ret.SendingPropertyName = "WHC"
		ret.Amount = "1"
	case 53, 55, 56:
		ret.ReferenceAddress = formatReferenceAddress(txd.SendingAddress)
		ret.SendingPropertyName = "WHC"
	case 70:
		ret.ReferenceAddress = formatReferenceAddress(txd.ReferenceAddress)
	case 68:
		ret.ReferenceAddress = formatReferenceAddressAmount(txd.SendingAddress, ret.Amount)
		ret.SendingPropertyName = "BCH"
		ret.PropertyName = "WHC"
		amount, err := decimal.NewFromString(txd.Amount)
		if err != nil {
			return nil, err
		}

		if !txd.Mature {
			ret.State = "unmature"
		}

		ret.Amount = amount.Div(decimal.NewFromFloat(100)).String()
	default:
		if txd.ReferenceAddress == "" {
			ret.ReferenceAddress = []string{}
		} else {
			ret.ReferenceAddress = formatReferenceAddress(txd.ReferenceAddress)
		}
	}

	if ret.PropertyName == "" {
		spName, err := model.GetPropertyName(uint64(txd.PropertyID))
		if err != nil {
			ctx.JSON(200, apiError(ErrGetPropertyNameByID))
			return nil, err
		}
		ret.PropertyName = spName
		ret.SendingPropertyName = spName
	}

	return &ret, nil
}

func formatSubsends(subsends []btcjson.SubSend, address string) ([]ApiSubSends, error) {
	ret := make([]ApiSubSends, 0, len(subsends))
	for _, item := range subsends {
		var api ApiSubSends
		api.SubSend = item
		api.Address = address

		// remove remaining '0'
		amount, err := decimal.NewFromString(item.Amount)
		if err != nil {
			return nil, err
		}
		api.Amount = amount.String()

		spName, err := model.GetPropertyName(uint64(item.PropertyID))
		if err != nil {
			return nil, err
		}
		api.PropertyName = spName

		ret = append(ret, api)
	}

	return ret, nil
}

func formatReferenceAddress(address string) interface{} {
	return []interface{}{
		map[string]string{"address": address},
	}
}

func formatReferenceAddressAmount(address string, amount string) interface{} {
	return []interface{}{
		map[string]string{"address": address, "amount": amount},
	}
}
