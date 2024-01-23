package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	"time"
	"todolist/app/taskkx/repository/db/dao"
	"todolist/app/taskkx/repository/mq"
	"todolist/app/taskkx/script"
	"todolist/app/taskkx/service"
	"todolist/config"
	"todolist/idl/task/kitex_gen/api/taskservice"
)

func main() {
	config.Init()
	dao.InitDB()
	mq.InitRabbitMq()
	loadingScript()
	// etcd注册
	etcdReg, err := etcd.NewEtcdRegistry([]string{fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)})
	if err != nil {
		log.Fatal(err)
	}
	// 注册微服务
	addr, err := net.ResolveTCPAddr("tcp", config.TaskServiceAddress)
	if err != nil {
		log.Fatal(err)
	}
	microserve := taskservice.NewServer(service.GetTaskSrv(), server.WithRegistry(etcdReg), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "rpcTaskService",
	}), server.WithServiceAddr(addr), server.WithReadWriteTimeout(5*time.Second), server.WithExitWaitTime(5*time.Second))
	microserve.Run()
}

func loadingScript() {
	ctx := context.Background()
	go script.TaskCreateSync(ctx)
}
