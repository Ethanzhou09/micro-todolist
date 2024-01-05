package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/app/gateway/rpc"
	"todolist/idl/pb"
	"todolist/pkg/ctl"
	"todolist/pkg/jwt"
	"todolist/types"
)

func UserRegisterHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-ShouldBindJSON"))
		return
	}
	userResp, err := rpc.UserRegister(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-UserRegister-RPC"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, userResp))
}

func UserLoginHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserLoginHandler-ShouldBindJSON"))
		return
	}
	userResp, err := rpc.UserLogin(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserLoginHandler-UserLogin-RPC"))
		return
	}
	token, err := jwt.GenerateToken(uint(userResp.UserDetail.Id))
	//ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, userResp))
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserLoginHandler-GenerateToken"))
		return
	}
	res := &types.TokenData{
		userResp,
		token,
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, res))
}
