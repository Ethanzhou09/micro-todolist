package rpc

import "C"
import (
	"context"
	"todolist/idl/kitex_gen/api"
	"todolist/pkg/e"
)

func UserLogin(ctx context.Context, req *api.UserRequest) (resp *api.UserResponse, err error) {
	if req != nil {
		resp, err = Client.UserLogin(ctx, req)
		if err != nil || resp.Code != e.Success {
			resp.Code = e.Error
			return
		}
	}
	return
}

func UserRegister(ctx context.Context, req *api.UserRequest) (resp *api.UserResponse, err error) {
	if req != nil {
		resp, err = Client.UserRegister(ctx, req)
		if err != nil || resp.Code != e.Success {
			resp.Code = e.Error
			return
		}
	}
	return
}
