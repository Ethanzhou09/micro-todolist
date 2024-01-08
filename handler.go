package main

import (
	"context"
	"todolist/idl/task/kitex_gen/api"
)

// TaskServiceImpl implements the last service interface defined in the IDL.
type TaskServiceImpl struct{}

// GetTaskList implements the TaskServiceImpl interface.
func (s *TaskServiceImpl) GetTaskList(ctx context.Context, req *api.TaskRequest) (resp *api.TaskListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetTask implements the TaskServiceImpl interface.
func (s *TaskServiceImpl) GetTask(ctx context.Context, req *api.TaskRequest) (resp *api.TaskDetailResponse, err error) {
	// TODO: Your code here...
	return
}

// CreateTask implements the TaskServiceImpl interface.
func (s *TaskServiceImpl) CreateTask(ctx context.Context, req *api.TaskRequest) (resp *api.TaskDetailResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateTask implements the TaskServiceImpl interface.
func (s *TaskServiceImpl) UpdateTask(ctx context.Context, req *api.TaskRequest) (resp *api.TaskDetailResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteTask implements the TaskServiceImpl interface.
func (s *TaskServiceImpl) DeleteTask(ctx context.Context, req *api.TaskRequest) (resp *api.TaskDetailResponse, err error) {
	// TODO: Your code here...
	return
}
