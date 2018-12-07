package routers

import (
	"github.com/copernet/whccommon/log"
	"github.com/copernet/whcexplorer/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(ginLogger())
	r.Use(log.LogContext([]string{"/static"}))

	root := r.Group("/explorer")
	{
		// acquire the current environment (mainnet / testnet)
		root.GET("/env", api.GetEnv)

		root.GET("/search/:keyword", api.Search)

		root.GET("/info", api.Info)
	}

	block := root.Group("/block")
	{
		block.GET("/list", api.ListBlocks)
		block.GET("/block/:keyword", api.GetBlock)
	}

	tx := root.Group("/tx")
	{
		tx.GET("/list", api.GetTransactionList)
		tx.GET("/hash/:txhash", api.GetTransaction)
		tx.GET("/latest", api.GetLatestInfo)
	}

	root.GET("/properties", api.Properties)
	// support fuzzy search via property name or property id
	root.GET("/properties/query", api.GetProperty)

	property := root.Group("/property")
	{
		property.GET("/:propertyId", api.PropertyByID)
		property.GET("/:propertyId/history", api.PropertyChangeHistory)
		property.GET("/:propertyId/txs", api.PropertyTransactions)
		property.GET("/:propertyId/listowners", api.ListOwners)

	}

	address := root.Group("/address")
	{
		address.GET("/:address", api.AddressDetail)
		address.GET("/:address/txs", api.AddressTransactions)
		address.GET("/:address/properties", api.AddressProperties)
		address.GET("/:address/property/:propertyId/txs", api.AddressPropertyTransactions)
	}

	burn := root.Group("/burn")
	{
		burn.GET("/list", api.GetBurnList)
		burn.GET("/summary", api.Summary)
	}

	r.StaticFile("/static", "./bone")

	return r
}
