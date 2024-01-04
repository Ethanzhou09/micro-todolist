package main

import (
	"fmt"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
	"time"
	"todolist/app/gateway/router"
	"todolist/app/gateway/rpc"
	"todolist/config"
)

func main() {
	config.Init()
	rpc.InitRPC()
	// etcd注册
	etcdReg := registry.NewRegistry(registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)))
	// 注册路由
	webService := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:4000"),
		web.Registry(etcdReg),
		web.Handler(router.NewRouter()),
		web.RegisterTTL(time.Second*30),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	_ = webService.Init()
	_ = webService.Run()
	//reghzgateway()
}

//func reghzgateway() {
//	config.Init()
//	etcdreg, err := etcd.NewEtcdRegistry([]string{fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)})
//	if err != nil {
//		log.Fatal(err)
//	}
//	svc := router.NewhzRouter(etcdreg, "127.0.0.1:4000")
//	svc.Spin()
//}
