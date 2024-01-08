package service

import (
	"context"
	"encoding/json"
	"sync"
	"todolist/app/taskkitex/repository/db/dao"
	"todolist/app/taskkitex/repository/db/model"
	"todolist/app/taskkitex/repository/mq"
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
