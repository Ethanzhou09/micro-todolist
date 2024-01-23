package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"todolist/pkg/ctl"
	"todolist/pkg/jwt"
)

// JWT token验证中间件
//func JWT() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var code uint32
//
//		code = 200
//		token := c.GetHeader("Authorization")
//		if token == "" {
//			code = 404
//			c.JSON(500, gin.H{
//				"code": code,
//				"msg":  "鉴权失败",
//			})
//		}
//		claims, err := jwt.ParseToken(token)
//		if err != nil {
//			code = 401
//			c.JSON(500, gin.H{
//				"code": code,
//				"msg":  "鉴权失败",
//			})
//			c.Abort()
//		}
//		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{Id: claims.Id}))
//		c.Next()
//	}
//}

func JWT() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var code uint32
		code = 200
		token := string(ctx.GetHeader("Authorization"))
		if token == "" {
			code = 404
			ctx.JSON(500, map[string]interface{}{
				"code": code,
				"msg":  "鉴权失败",
			})
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			code = 401
			ctx.JSON(500, map[string]interface{}{
				"code": code,
				"msg":  "鉴权失败",
			})
			ctx.Abort()
		}
		ctx.Next(ctl.NewContext(c, &ctl.UserInfo{Id: claims.Id}))
	}
}
