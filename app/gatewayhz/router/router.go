package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/gin-gonic/gin"
	"todolist/app/gateway/http"
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

		//v1.POST("user/login", http.UserLoginHandler)
		//v1.POST("user/register", http.UserRegisterHandler)
	}
	return hertzRouter
}

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "success")
		})
		v1.POST("user/login", http.UserLoginHandler)
		v1.POST("user/register", http.UserRegisterHandler)
	}
	return ginRouter
}
