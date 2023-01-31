package main

import (
	"context"
	"github.com/fvbock/endless"
	"handshake/conf"
	"handshake/engine"
	"handshake/persistent"
	"handshake/router"
	"strconv"
)

func main() {
	conf.Init()
	context.Background()
	persistent.Init()
	defer persistent.Close()
	// 启动引擎
	engine.Init()
	r := router.Router()
	addr := structure()
	endless.ListenAndServe(addr, r)
}

func structure() (addr string) {
	server := conf.ServerConf()
	addr = ":"
	portString := strconv.Itoa(server.ServerPort)
	addr = addr + portString
	return addr
}
