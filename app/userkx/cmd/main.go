package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	"time"
	"todolist/app/userkx/repository/db/dao"
	"todolist/app/userkx/service"
	"todolist/config"
	"todolist/idl/kitex_gen/api/myservice"
)

func main() {
	config.Init()
	dao.InitDB()
	// etcd注册
	etcdReg, err := etcd.NewEtcdRegistry([]string{fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)})
	if err != nil {
		log.Fatal(err)
	}
	// 注册微服务
	addr, err := net.ResolveTCPAddr("tcp", config.UserServiceAddress)
	if err != nil {
		log.Fatal(err)
	}
	microserve := myservice.NewServer(service.GetUserSrv(), server.WithRegistry(etcdReg), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "rpcUserService",
	}), server.WithServiceAddr(addr), server.WithReadWriteTimeout(5*time.Second), server.WithExitWaitTime(5*time.Second))
	microserve.Run()
}
