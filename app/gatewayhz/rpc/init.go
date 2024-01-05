package rpc

import (
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"time"
	"todolist/config"
	"todolist/idl/kitex_gen/api/myservice"
)

var (
	Client myservice.Client
)

func InitRPC() {
	// 初始化rpc服务
	r, err := etcd.NewEtcdResolver([]string{fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)})
	if err != nil {
		log.Fatal(err)
	}
	c, err := myservice.NewClient("rpcUserService", client.WithResolver(r), client.WithConnectTimeout(5*time.Second),
		client.WithRPCTimeout(0))
	if err != nil {
		log.Fatal(err)
	}
	Client = c
}
