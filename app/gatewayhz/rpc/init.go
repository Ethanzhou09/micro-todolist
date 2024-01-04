package rpc

import (
	"go-micro.dev/v4"
	"todolist/idl/pb"
)

var (
	UserService pb.UserService
)

func InitRPC() {
	// 初始化rpc服务
	userMicroService := micro.NewService(micro.Name("userService.client"))
	// 初始化rpc客户端
	userService := pb.NewUserService("rpcUserService", userMicroService.Client())
	// 赋值给全局变量
	UserService = userService
}
