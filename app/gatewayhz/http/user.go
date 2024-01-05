package http

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"todolist/app/gatewayhz/rpc"
	"todolist/idl/kitex_gen/api"
)

func UserRegisterHandler(c context.Context, ctx *app.RequestContext) {
	req := api.UserRequest{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, "UserRegisterHandler-ShouldBindJSON")
		return
	}
	resp, err := rpc.UserRegister(c, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, "UserRegisterHandler-UserRegister-RPC")
		return
	}
	ctx.JSON(http.StatusOK, resp.User.Username+" create success")
}

//func UserLoginHandler(ctx *gin.Context) {
//	var req api.UserRequest
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserLoginHandler-ShouldBindJSON"))
//		return
//	}
//	userResp, err := rpc.UserLogin(ctx, &req)
//	if err != nil {
//		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserLoginHandler-UserLogin-RPC"))
//		return
//	}
//	token, err := jwt.GenerateToken(uint(userResp.UserDetail.Id))
//	//ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, userResp))
//	if err != nil {
//		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserLoginHandler-GenerateToken"))
//		return
//	}
//	res := &types.TokenData{
//		userResp,
//		token,
//	}
//	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, res))
//}
