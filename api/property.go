package api

import (
	"github.com/copernet/whccommon/log"
	cmodel "github.com/copernet/whccommon/model"
	"github.com/copernet/whcexplorer/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	whcPropertyName = "WHC"
)

func Properties(c *gin.Context) {
	pageInfo := cmodel.GetPageInfo(c)
	propertyType := cmodel.DefaultInt("propertyType", -1, c)
	propertyTypes := make([]int, 0)
	if propertyType != -1 {
		propertyTypes = append(propertyTypes, propertyType)
	}

	total := model.CountProperties(propertyTypes)
	list, err := model.ListProperties(propertyTypes, pageInfo)
	pageInfo.Total = total
	pageInfo.List = list
	if err == nil {
		c.JSON(200, apiSuccess(pageInfo))
		return
	}
	c.JSON(200, apiError(ErrGetProperties)) 
}

func PropertyByID(c *gin.Context) {
	propertyId := cmodel.DefaultInt("propertyId", -1, c)
	detail, err := model.PropertyDetail(int64(propertyId))
	if err == nil {
		c.JSON(200, apiSuccess(map[string]interface{}{
			"data": detail,
		}))
		return
	}
	c.JSON(200, apiError(ErrGetPropertyByID)) 
}

func PropertyChangeHistory(c *gin.Context) {
	pageInfo := cmodel.GetPageInfo(c)
	propertyId := int64(cmodel.DefaultInt("propertyId", -1, c))
	pageInfo.Total = model.CountPropertyChangeLog(propertyId)
	list, err := model.ListPropertyChangeLog(propertyId, pageInfo)
	pageInfo.List = list
	if err == nil {
		c.JSON(200, apiSuccess(pageInfo))
		return
	}
	c.JSON(200, apiError(ErrGetPropertyChangeHistory)) 
}

func PropertyTransactions(c *gin.Context) {
	pageInfo := cmodel.GetPageInfo(c)
	propertyId := int64(cmodel.DefaultInt("propertyId", -1, c))
	eventType := cmodel.DefaultInt("eventType", -1, c)
	list, err := model.ListPropertyTxes(int64(propertyId), eventType, pageInfo)
	pageInfo.Total = model.CountPropertyTxes(int64(propertyId), eventType)
	pageInfo.List = list
	if err == nil {
		c.JSON(200, apiSuccess(pageInfo))
		return
	}
	c.JSON(200, apiError(ErrGetPropertyTransactions)) 
}


func ListOwners(c *gin.Context) {
	pid, err := strconv.Atoi(c.Param("propertyId"))
	total, err := model.ListOwnersCount(pid)
	if err != nil {
		log.WithCtx(c).Errorf("get address balance total number failed: %v", err)
		c.JSON(200, apiError(ErrListPropertiesCount))
		return
	}

	if total == 0 {
		c.JSON(200, apiSuccess(map[string]interface{}{
			"total": 0,
			"list":  []string{},
		}))
		return
	}

	pageSize, pageNo := paginator(c)
	list, err := model.ListOwners(pageSize, pageNo, pid)
	if err != nil {
		log.WithCtx(c).Errorf("get properties list failed: %v", err)

		c.JSON(200, apiError(ErrListProperties))
		return
	}

	c.JSON(200, apiSuccess(map[string]interface{}{
		"total": total,
		"list":  list,
	}))
}

func IsAvailableQueryForProperty(keyword string) bool {
	if len(keyword) > propertyNameLimit {
		return false
	}

	return true
}

// GetProperty query property data via property name of property id.
// will return most ten items for select.
func GetProperty(c *gin.Context) {
	pageInfo := cmodel.GetPageInfo(c)
	keyword := c.Query("keyword")

	if !IsAvailableQueryForProperty(keyword) {
		c.JSON(200, apiError(ErrIncorrectPropertyQueryByKeyword))
		return
	}

	property, err := model.GetPropertyByKeyword(keyword, pageInfo)
		if err != nil {
		log.WithCtx(c).Errorf("fetch property data by keyword(id/name) failed: %v", err)
		c.JSON(200, apiError(ErrGetPropertyByName))
		return
	}
	pageInfo.List = property
	pageInfo.Total = len(property)

	c.JSON(200, apiSuccess(pageInfo))
}