package model

import (
	"fmt"
	"os"

	common "github.com/copernet/whccommon/model"
	"github.com/copernet/whcexplorer/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	conf := config.GetConf()

	var err error
	db, err = common.ConnectDatabase(conf.DB)
	if err != nil {
		fmt.Printf("initial database error: %v", err)
		os.Exit(1)
	}
}
func InitWhenTesting(option *common.DBOption) {
	var err error
	db, err = common.ConnectDatabase(option)
	if err != nil {
		fmt.Printf("initial database error: %v", err)
		os.Exit(1)
	}
}
