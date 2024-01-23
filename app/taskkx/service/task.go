package service

import (
	"context"
	"encoding/json"
	"sync"
	"todolist/app/taskkx/repository/db/dao"
	"todolist/app/taskkx/repository/db/model"
	"todolist/app/taskkx/repository/mq"
	"todolist/idl/task/kitex_gen/api"
	"todolist/pkg/e"
)

type Tasksrv struct{}

var TasksrvIns *Tasksrv
var TasksrvOnce sync.Once

// GetTaskSrv 获取单例
func GetTaskSrv() *Tasksrv {
	// 通过sync.Once实现单例
	TasksrvOnce.Do(func() {
		TasksrvIns = &Tasksrv{}
	})
	return TasksrvIns
}

// create task ----->送到mq ----->落库
func (t *Tasksrv) CreateTask(ctx context.Context, req *api.TaskRequest) (resp *api.TaskDetailResponse, err error) {
	resp = &api.TaskDetailResponse{}
	resp.Code = e.Success
	body, _ := json.Marshal(req)
	err = mq.SendMessage2MQ(body)
	if err != nil {
		resp.Code = e.Error
		return
	}
	return
}

func TaskMq2DB(ctx context.Context, req *api.TaskRequest) (err error) {
	m := &model.Task{
		Uid:       uint(req.Uid),
		Title:     req.Title,
		Content:   req.Content,
		Status:    int(req.Status),
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	return dao.NewTaskDao(ctx).CreateTask(m)

}

func (*Tasksrv) GetTaskList(ctx context.Context, req *api.TaskRequest) (resp *api.TaskListResponse, err error) {
	resp = &api.TaskListResponse{}
	resp.Code = e.Success
	if req.Limit == 0 {
		req.Limit = 10
	}
	r, count, err := dao.NewTaskDao(ctx).ListTaskByUserId(uint64(req.Uid), int(req.Start), int(req.Limit))
	if err != nil {
		resp.Code = e.Error
		return
	}
	var taskRes []*api.TaskModel
	for _, item := range r {
		taskRes = append(taskRes, BuildTask(item))
	}
	resp.TaskList = taskRes
	resp.Count = count
	return
}

func (*Tasksrv) GetTask(ctx context.Context, req *api.TaskRequest) (resp *api.TaskDetailResponse, err error) {
	resp = &api.TaskDetailResponse{}
	resp.Code = e.Success
	r, err := dao.NewTaskDao(ctx).GetTaskById(uint64(req.Id), uint64(req.Uid))
	if r.ID == 0 {
		resp.Code = e.Error
		return
	}
	resp.TaskDetail = BuildTask(r)
	return
}

func (*Tasksrv) UpdateTask(ctx context.Context, req *api.TaskRequest) (resp *api.TaskDetailResponse, err error) {
	resp = &api.TaskDetailResponse{}
	resp.Code = e.Success
	err = dao.NewTaskDao(ctx).UpdateTask(req)
	if err != nil {
		resp.Code = e.Error
	}
	return
}

func (*Tasksrv) DeleteTask(ctx context.Context, req *api.TaskRequest) (resp *api.TaskDetailResponse, err error) {
	resp = &api.TaskDetailResponse{}
	resp.Code = e.Success
	err = dao.NewTaskDao(ctx).DeleteTask(uint64(req.Id), uint64(req.Uid))
	if err != nil {
		resp.Code = e.Error
	}
	return
}

func BuildTask(item *model.Task) *api.TaskModel {
	return &api.TaskModel{
		Id:         int64(item.ID),
		Uid:        int64(item.Uid),
		Title:      item.Title,
		Content:    item.Content,
		StartTime:  item.StartTime,
		EndTime:    item.EndTime,
		Status:     int64(item.Status),
		CreateTime: item.CreatedAt.Unix(),
		UpdateTime: item.UpdatedAt.Unix(),
	}
}
