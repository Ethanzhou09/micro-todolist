package ctl

import (
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"golang.org/x/net/context"
)

type key int

var userKey key

type UserInfo struct {
	Id uint `json:"id"`
}

func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userKey).(*UserInfo)
	return u, ok
}

// 从context中获取用户信息
func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	u, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("获取用户信息失败")
	}
	return u, nil
}

func FromHzContext(ctx *app.RequestContext) (*UserInfo, bool) {
	u, ok := ctx.Value(userKey).(*UserInfo)
	return u, ok
}

func GetUserInfoHz(ctx *app.RequestContext) (*UserInfo, error) {
	u, ok := FromHzContext(ctx)
	if !ok {
		return nil, errors.New("获取用户信息失败")
	}
	return u, nil
}
