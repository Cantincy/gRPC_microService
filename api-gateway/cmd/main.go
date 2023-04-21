package main

import (
	"pro01/grpc/api-gateway/router"
	"pro01/grpc/api-gateway/util"
)

func main() {
	util.InitCli()
	engine := router.NewRouter()
	engine.Run(":9090")
}
