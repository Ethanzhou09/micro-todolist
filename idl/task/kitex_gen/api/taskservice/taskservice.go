// Code generated by Kitex v0.8.0. DO NOT EDIT.

package taskservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"todolist/idl/task/kitex_gen/api"
)

func serviceInfo() *kitex.ServiceInfo {
	return taskServiceServiceInfo
}

var taskServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "TaskService"
	handlerType := (*api.TaskService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetTaskList": kitex.NewMethodInfo(getTaskListHandler, newTaskServiceGetTaskListArgs, newTaskServiceGetTaskListResult, false),
		"GetTask":     kitex.NewMethodInfo(getTaskHandler, newTaskServiceGetTaskArgs, newTaskServiceGetTaskResult, false),
		"CreateTask":  kitex.NewMethodInfo(createTaskHandler, newTaskServiceCreateTaskArgs, newTaskServiceCreateTaskResult, false),
		"UpdateTask":  kitex.NewMethodInfo(updateTaskHandler, newTaskServiceUpdateTaskArgs, newTaskServiceUpdateTaskResult, false),
		"DeleteTask":  kitex.NewMethodInfo(deleteTaskHandler, newTaskServiceDeleteTaskArgs, newTaskServiceDeleteTaskResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "api",
		"ServiceFilePath": `idl\taskSrv.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func getTaskListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.TaskServiceGetTaskListArgs)
	realResult := result.(*api.TaskServiceGetTaskListResult)
	success, err := handler.(api.TaskService).GetTaskList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTaskServiceGetTaskListArgs() interface{} {
	return api.NewTaskServiceGetTaskListArgs()
}

func newTaskServiceGetTaskListResult() interface{} {
	return api.NewTaskServiceGetTaskListResult()
}

func getTaskHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.TaskServiceGetTaskArgs)
	realResult := result.(*api.TaskServiceGetTaskResult)
	success, err := handler.(api.TaskService).GetTask(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTaskServiceGetTaskArgs() interface{} {
	return api.NewTaskServiceGetTaskArgs()
}

func newTaskServiceGetTaskResult() interface{} {
	return api.NewTaskServiceGetTaskResult()
}

func createTaskHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.TaskServiceCreateTaskArgs)
	realResult := result.(*api.TaskServiceCreateTaskResult)
	success, err := handler.(api.TaskService).CreateTask(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTaskServiceCreateTaskArgs() interface{} {
	return api.NewTaskServiceCreateTaskArgs()
}

func newTaskServiceCreateTaskResult() interface{} {
	return api.NewTaskServiceCreateTaskResult()
}

func updateTaskHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.TaskServiceUpdateTaskArgs)
	realResult := result.(*api.TaskServiceUpdateTaskResult)
	success, err := handler.(api.TaskService).UpdateTask(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTaskServiceUpdateTaskArgs() interface{} {
	return api.NewTaskServiceUpdateTaskArgs()
}

func newTaskServiceUpdateTaskResult() interface{} {
	return api.NewTaskServiceUpdateTaskResult()
}

func deleteTaskHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.TaskServiceDeleteTaskArgs)
	realResult := result.(*api.TaskServiceDeleteTaskResult)
	success, err := handler.(api.TaskService).DeleteTask(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTaskServiceDeleteTaskArgs() interface{} {
	return api.NewTaskServiceDeleteTaskArgs()
}

func newTaskServiceDeleteTaskResult() interface{} {
	return api.NewTaskServiceDeleteTaskResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetTaskList(ctx context.Context, req *api.TaskRequest) (r *api.TaskListResponse, err error) {
	var _args api.TaskServiceGetTaskListArgs
	_args.Req = req
	var _result api.TaskServiceGetTaskListResult
	if err = p.c.Call(ctx, "GetTaskList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetTask(ctx context.Context, req *api.TaskRequest) (r *api.TaskDetailResponse, err error) {
	var _args api.TaskServiceGetTaskArgs
	_args.Req = req
	var _result api.TaskServiceGetTaskResult
	if err = p.c.Call(ctx, "GetTask", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateTask(ctx context.Context, req *api.TaskRequest) (r *api.TaskDetailResponse, err error) {
	var _args api.TaskServiceCreateTaskArgs
	_args.Req = req
	var _result api.TaskServiceCreateTaskResult
	if err = p.c.Call(ctx, "CreateTask", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateTask(ctx context.Context, req *api.TaskRequest) (r *api.TaskDetailResponse, err error) {
	var _args api.TaskServiceUpdateTaskArgs
	_args.Req = req
	var _result api.TaskServiceUpdateTaskResult
	if err = p.c.Call(ctx, "UpdateTask", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteTask(ctx context.Context, req *api.TaskRequest) (r *api.TaskDetailResponse, err error) {
	var _args api.TaskServiceDeleteTaskArgs
	_args.Req = req
	var _result api.TaskServiceDeleteTaskResult
	if err = p.c.Call(ctx, "DeleteTask", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
