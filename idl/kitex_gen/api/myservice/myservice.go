// Code generated by Kitex v0.7.2. DO NOT EDIT.

package myservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"todolist/idl/kitex_gen/api"
)

func serviceInfo() *kitex.ServiceInfo {
	return myServiceServiceInfo
}

var myServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "MyService"
	handlerType := (*api.MyService)(nil)
	methods := map[string]kitex.MethodInfo{
		"UserRegister": kitex.NewMethodInfo(userRegisterHandler, newMyServiceUserRegisterArgs, newMyServiceUserRegisterResult, false),
		"UserLogin":    kitex.NewMethodInfo(userLoginHandler, newMyServiceUserLoginArgs, newMyServiceUserLoginResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "api",
		"ServiceFilePath": `idl\userSrv.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.2",
		Extra:           extra,
	}
	return svcInfo
}

func userRegisterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.MyServiceUserRegisterArgs)
	realResult := result.(*api.MyServiceUserRegisterResult)
	success, err := handler.(api.MyService).UserRegister(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMyServiceUserRegisterArgs() interface{} {
	return api.NewMyServiceUserRegisterArgs()
}

func newMyServiceUserRegisterResult() interface{} {
	return api.NewMyServiceUserRegisterResult()
}

func userLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.MyServiceUserLoginArgs)
	realResult := result.(*api.MyServiceUserLoginResult)
	success, err := handler.(api.MyService).UserLogin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMyServiceUserLoginArgs() interface{} {
	return api.NewMyServiceUserLoginArgs()
}

func newMyServiceUserLoginResult() interface{} {
	return api.NewMyServiceUserLoginResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UserRegister(ctx context.Context, req *api.UserRequest) (r *api.UserResponse, err error) {
	var _args api.MyServiceUserRegisterArgs
	_args.Req = req
	var _result api.MyServiceUserRegisterResult
	if err = p.c.Call(ctx, "UserRegister", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserLogin(ctx context.Context, req *api.UserRequest) (r *api.UserResponse, err error) {
	var _args api.MyServiceUserLoginArgs
	_args.Req = req
	var _result api.MyServiceUserLoginResult
	if err = p.c.Call(ctx, "UserLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}