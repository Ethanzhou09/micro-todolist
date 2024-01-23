package http

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"todolist/app/gatewayhz/rpc"
	"todolist/idl/kitex_gen/api"
	"todolist/pkg/jwt"
	"todolist/types"
)

func UserRegisterHandler(c context.Context, ctx *app.RequestContext) {
	req := api.UserRequest{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, "UserRegisterHandler-ShouldBindJSON")
		return
	}
	_, err := rpc.UserRegister(c, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, "UserRegisterHandler-UserRegister-RPC")
		return
	}
	ctx.JSON(http.StatusOK, "create success")
}

func UserLoginHandler(c context.Context, ctx *app.RequestContext) {
	var req api.UserRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, "UserLoginHandler-ShouldBindJSON")
		return
	}
	userResp, err := rpc.UserLogin(c, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, "UserLoginHandler-UserLogin-RPC")
		return
	}
	token, err := jwt.GenerateToken(uint(userResp.User.Id))
	if err != nil {
		ctx.JSON(http.StatusOK, "UserLoginHandler-GenerateToken")
		return
	}
	res := &types.TokenData{
		userResp,
		token,
	}
	ctx.JSON(http.StatusOK, res)
}
