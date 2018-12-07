package api

import (
	cmodel "github.com/copernet/whccommon/model"
	"github.com/copernet/whcexplorer/model"
	"github.com/gin-gonic/gin"
)

func AddressDetail(c *gin.Context) {
	address := c.Param("address")
	balanceAvail, balanceTotal, err1 := model.WhcBalance(address)
	consumedBal, sendedBal, receivedBal, err2 := model.WhcStatistics(address)
	whcTxCount, err3 := model.WhcTxCountStatistics(address)
	if err1 == nil && err2 == nil && err3 == nil {
		c.JSON(200, apiSuccess(map[string]interface{}{
			"balanceAvail": balanceAvail,
			"balanceTotal": balanceTotal,
			"consumedBal":  consumedBal,
			"sendedBal":    sendedBal,
			"receivedBal":  receivedBal,
			"whcTxCount":   whcTxCount,
		}))
		return
	}
	c.JSON(200, apiError(ErrGetAddressDetail))

}

func setDirection(list []TransactionDetail, directions []int) {
	for i:=0; i<len(list);i++ {
		list[i].Direction = directions[i]
	}
}
func AddressTransactions(c *gin.Context) {
	address := c.Param("address")
	pageInfo := cmodel.GetPageInfo(c)
	txType := cmodel.DefaultInt("txType", -1, c)
	beginTime := cmodel.DefaultInt64("beginTime", -1, c)
	endTime := cmodel.DefaultInt64("endTime", -1, c)
	times := make([]int64, 0)
	if beginTime != -1 && endTime != -1 {
		times = append(times, beginTime, endTime)
	}
	list, directions, err := model.ListAddressTxes(address, txType, times, -1, pageInfo)
	currentHeight := model.GetLastBlock().BlockHeight
	ret, err2 := formatTransactionsForAPI(list, currentHeight, c)
	setDirection(ret, directions)
	pageInfo.Total = model.CountAddressTxes(address, txType, times, -1)
	pageInfo.List = ret

	if err == nil && err2 == nil{
		c.JSON(200, apiSuccess(pageInfo))
		return
	}
	c.JSON(200, apiError(ErrGetAddressTransactions))
}

func AddressProperties(c *gin.Context) {
	address := c.Param("address")
	propertyType := cmodel.DefaultInt("propertyType", -1, c)
	pageInfo := cmodel.GetPageInfo(c)
	list, err := model.ListAddressProperties(address, propertyType, pageInfo)
	pageInfo.Total = model.CountAddressProperties(address, propertyType)
	pageInfo.List = list
	if err == nil {
		c.JSON(200, apiSuccess(pageInfo))
		return
	}
	c.JSON(200, apiError(ErrGetAddressProperties))
}

func AddressPropertyTransactions(c *gin.Context) {
	address := c.Param("address")
	propertyId := int64(cmodel.DefaultInt("propertyId", -1, c))
	pageInfo := cmodel.GetPageInfo(c)

	list, directions, err := model.ListAddressTxes(address, -1, nil, propertyId, pageInfo)
	currentHeight := model.GetLastBlock().BlockHeight
	ret, err2 := formatTransactionsForAPI(list, currentHeight, c)
	setDirection(ret, directions)
	pageInfo.Total = model.CountAddressTxes(address, -1, nil, propertyId)
	pageInfo.List = ret
	if err == nil && err2 == nil {
		c.JSON(200, apiSuccess(pageInfo))
		return
	}
	c.JSON(200, apiError(ErrGetAddressPropertyTransactions))
}
