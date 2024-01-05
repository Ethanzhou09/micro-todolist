package main

import (
	"fmt"
	"github.com/hertz-contrib/registry/etcd"
	"log"
	"todolist/app/gatewayhz/router"
	"todolist/app/gatewayhz/rpc"
	"todolist/config"
)

func main() {
	config.Init()
	rpc.InitRPC()
	etcdreg, err := etcd.NewEtcdRegistry([]string{fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)})
	if err != nil {
		log.Fatal(err)
	}
	svc := router.NewhzRouter(etcdreg, "127.0.0.1:4000")
	svc.Spin()
}
