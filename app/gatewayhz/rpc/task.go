package rpc

import (
	"context"
	"todolist/idl/task/kitex_gen/api"
	"todolist/pkg/e"
)

func TaskCreate(ctx context.Context, req *api.TaskRequest) (resp *api.TaskDetailResponse, err error) {
	if req != nil {
		//ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		resp, err = TaskClient.CreateTask(ctx, req)
		//cancel()
		if err != nil || resp.Code != e.Success {
			resp.Code = e.Error
			return
		}
	}
	return
}

func TaskUpdate(ctx context.Context, req *api.TaskRequest) (resp *api.TaskDetailResponse, err error) {
	if req != nil {
		resp, err = TaskClient.UpdateTask(ctx, req)
		if err != nil || resp.Code != e.Success {
			resp.Code = e.Error
			return
		}
	}
	return
}

func TaskDelete(ctx context.Context, req *api.TaskRequest) (resp *api.TaskDetailResponse, err error) {
	if req != nil {
		resp, err = TaskClient.DeleteTask(ctx, req)
		if err != nil || resp.Code != e.Success {
			resp.Code = e.Error
			return
		}
	}
	return
}

func TaskListGet(ctx context.Context, req *api.TaskRequest) (resp *api.TaskListResponse, err error) {
	if req != nil {
		resp, err = TaskClient.GetTaskList(ctx, req)
		if err != nil || resp.Code != e.Success {
			resp.Code = e.Error
			return
		}
	}
	return
}

func TaskGet(ctx context.Context, req *api.TaskRequest) (resp *api.TaskDetailResponse, err error) {
	if req != nil {
		resp, err = TaskClient.GetTask(ctx, req)
		if err != nil || resp.Code != e.Success {
			resp.Code = e.Error
			return
		}
	}
	return
}
