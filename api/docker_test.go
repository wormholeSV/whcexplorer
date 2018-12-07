package api_test

import (
	"testing"
	"github.com/copernet/whcexplorer/util"
	"github.com/copernet/whcexplorer/model"
	"github.com/copernet/whcexplorer/config"
	"github.com/copernet/whcexplorer/routers"
)

func TestMain(m *testing.M) {
	if config.GetConf().UseDockerTest {
		resource, pool, dbOption := util.SetupDockerDB(dbDocker, "../model")
		model.InitWhenTesting(dbOption)
		testRouter = routers.InitRouter()
		code := m.Run()
		util.TeardownDockerDB(resource, pool, code)
	} else {
		m.Run()
	}
}
