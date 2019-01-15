# Wormhole Explorer API Service
-
## What is Wormhole

Wormhole is a fast, portable Omni Layer implementation that is based off the Bitcoin ABC codebase (currently 0.18.2). This implementation requires no external dependencies extraneous to Bitcoin ABC, and is native to the Bitcoin Cash network just like other Bitcoin Cash nodes. It currently supports a wallet mode and is seamlessly available on three platforms: Windows, Linux and Mac OS. Wormhole Cash Layer extensions are exposed via the JSON-RPC interface. Development has been consolidated on the Wormhole product, and it is the reference client for the Wormhole Cash Layer.

## What is Wormhole Explorer
[https://explorer.wormhole.cash](https://explorer.wormhole.cash)

## Quick Start

#### Prerequesites
| Package | Version | 
| :------| ------|
| Mysql | 5.7+ |
| Golang | 1.10+ |
| Redis | 4.0+ |
|[Wormhole](https://github.com/copernet/wormhole)|0.2.2|
|[Engine](https://github.com/wormholeSV/whcengine)|0.0.1|

#### Config Init

	go get -d github.com/wormholeSV/whcexplorer
	cd ${gopath}/src/github.com/wormholeSV/whcexplorer
	cp conf.yml.sample conf.yml

	#you need modify db、redis、rpc、log to your local config
#### How To Run

	cd ${gopath}/src/github.com/wormholeSV/whcexplorer
	mkdir logs
	touch logs/system.log
	go build

	#start
	tools/run start whcexplorer
	#stop 
	tools/run stop whcexplorer

## Document
- [API Document](./doc/api.md)
