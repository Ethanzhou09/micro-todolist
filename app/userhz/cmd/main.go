package main

import (
	"fmt"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"todolist/app/user/repository/db/dao"
	"todolist/app/user/service"
	"todolist/config"
	"todolist/idl/pb"
)

func main() {
	config.Init()
	dao.InitDB()
	// etcd注册
	etcdReg := registry.NewRegistry(registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)))
	// 注册微服务
	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address(config.UserServiceAddress),
		micro.Registry(etcdReg))
	microService.Init()
	_ = pb.RegisterUserServiceHandler(microService.Server(), service.GetUserSrv())
	_ = microService.Run()
}
