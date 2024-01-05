package ctl

import (
	"github.com/gin-gonic/gin"
	"todolist/pkg/e"
)

type Response struct {
	Status int         `json:"status,omitempty"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}

func RespError(ctx *gin.Context, err error, data string, code ...int) *Response {
	status := e.Error
	if code != nil {
		status = code[0]
	}
	return &Response{
		status,
		data,
		e.GetMsg(status),
		err.Error(),
	}
}

func RespSuccess(ctx *gin.Context, data interface{}, code ...int) *Response {
	status := e.Success
	if code != nil {
		status = code[0]
	}
	if data == nil {
		data = "操作成功"
	}
	return &Response{
		status,
		data.(string),
		e.GetMsg(status),
		"nil",
	}
}
