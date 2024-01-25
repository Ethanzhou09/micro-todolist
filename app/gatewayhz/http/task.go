package http

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"todolist/app/gatewayhz/rpc"
	"todolist/idl/task/kitex_gen/api"
)

func CreateTaskHandler(c context.Context, ctx *app.RequestContext) {
	req := api.TaskRequest{}
	if err := ctx.Bind(&req); err != nil {
		print(err.Error())
		ctx.JSON(200, err.Error())
		return
	}
	taskRes, err := rpc.TaskCreate(c, &req)
	if err != nil {
		ctx.JSON(200, "CreateTaskHandler-TaskCreate")
		return
	}
	ctx.JSON(200, taskRes)
}

func UpdateTaskHandler(c context.Context, ctx *app.RequestContext) {
	req := api.TaskRequest{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(200, "CreateTaskHandler-ShouldBindJSON")
		return
	}
	taskRes, err := rpc.TaskUpdate(c, &req)
	if err != nil {
		ctx.JSON(200, "CreateTaskHandler-TaskUpdate")
		return
	}
	ctx.JSON(200, taskRes)
}

func DeleteTaskHandler(c context.Context, ctx *app.RequestContext) {
	req := api.TaskRequest{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(200, "CreateTaskHandler-ShouldBindJSON")
		return
	}
	taskRes, err := rpc.TaskDelete(c, &req)
	if err != nil {
		ctx.JSON(200, "CreateTaskHandler-TaskDelete")
		return
	}
	ctx.JSON(200, taskRes)
}

func GetTaskHandler(c context.Context, ctx *app.RequestContext) {
	req := api.TaskRequest{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(200, "CreateTaskHandler-ShouldBindJSON")
		return
	}
	taskRes, err := rpc.TaskGet(c, &req)
	if err != nil {
		ctx.JSON(200, "CreateTaskHandler-TaskGet")
		return
	}
	ctx.JSON(200, taskRes)
}

func TaskListGetHandler(c context.Context, ctx *app.RequestContext) {
	req := api.TaskRequest{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(200, "CreateTaskHandler-ShouldBindJSON")
		return
	}
	taskRes, err := rpc.TaskListGet(c, &req)
	if err != nil {
		ctx.JSON(200, "CreateTaskHandler-TaskListGet")
		return
	}
	ctx.JSON(200, taskRes)
}
