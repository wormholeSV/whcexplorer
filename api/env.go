package api

import (
	"github.com/copernet/whcexplorer/config"
	"github.com/copernet/whcexplorer/api/common"
	"github.com/gin-gonic/gin"
)

func GetEnv(c *gin.Context) {
	if config.GetConf().TestNet {
		c.JSON(200, gin.H{
			"code":    common.Success,
			"message": "",
			"result":  map[string]bool{"testnet": true},
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    common.Success,
		"message": "",
		"result":  map[string]bool{"testnet": false},
	})

	c.JSON(200, apiSuccess(map[string]bool{"testnet": false}))

	return
}
