package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"todolist/app/gatewayhz/http"
	"todolist/app/gatewayhz/middleware"
)

func NewhzRouter(r registry.Registry, addr string) *server.Hertz {
	hertzRouter := server.Default(server.WithHostPorts(addr),
		server.WithRegistry(r, &registry.Info{
			ServiceName: "gateway.test.demo",
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      10,
			Tags:        nil,
		}),
	)
	v1 := hertzRouter.Group("/api/v1")
	{
		v1.GET("ping", func(c context.Context, ctx *app.RequestContext) {
			ctx.JSON(200, "success")
		})
		v1.POST("user/login", http.UserLoginHandler)
		v1.POST("user/register", http.UserRegisterHandler)
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("task", http.CreateTaskHandler)
			authed.POST("update_task", http.UpdateTaskHandler)
			authed.DELETE("delete_task", http.DeleteTaskHandler)
			authed.GET("get_tasklist", http.TaskListGetHandler)
			authed.GET("get_task", http.GetTaskHandler)
		}
	}
	return hertzRouter
}
