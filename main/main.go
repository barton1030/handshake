package main

import (
	"context"
	"github.com/fvbock/endless"
	"handshake/conf"
	"handshake/domain"
	"handshake/persistent"
	"handshake/router"
	"strconv"
)

func main() {
	conf.Init()
	context.Background()
	persistent.Init()
	defer persistent.Close()
	go domain.Init()
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
