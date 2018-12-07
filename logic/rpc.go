package logic

import (
	"github.com/copernet/whc.go/rpcclient"
	"github.com/copernet/whccommon/model"
	"github.com/copernet/whcexplorer/config"
)

var client *rpcclient.Client

func GetRPCIns() *rpcclient.Client {
	if client != nil {
		return client
	}

	return model.ConnRpc(config.GetConf().RPC)
}
