package model_test

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"github.com/copernet/whcexplorer/config"
	"github.com/copernet/whcexplorer/util"
	"github.com/copernet/whcexplorer/model"
)

var dbDocker *sql.DB

func TestMain(m *testing.M) {
	if config.GetConf().UseDockerTest {
		resource, pool, dbOpiton := util.SetupDockerDB(dbDocker,"./")
		model.InitWhenTesting(dbOpiton)
		code := m.Run()
		util.TeardownDockerDB(resource, pool, code)
	} else {
		m.Run()
	}
}


