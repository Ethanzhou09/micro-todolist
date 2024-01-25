package dao

import (
	"context"
	"gorm.io/gorm"
	"todolist/app/taskkx/repository/db/model"
	"todolist/idl/task/kitex_gen/api"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewDBClient(ctx)}
}

func (dao *TaskDao) CreateTask(data *model.Task) (err error) {
	return dao.Model(&model.Task{}).Create(data).Error
}

func (dao *TaskDao) ListTaskByUserId(userid uint64, start, limit int) (r []*model.Task, count int64, err error) {
	r = []*model.Task{}
	err = dao.Model(&model.Task{}).Offset(start).Limit(limit).Where("uid = ?", userid).Find(r).Error
	if err != nil {
		return
	}
	err = dao.Model(&model.Task{}).Where("uid = ?", userid).Count(&count).Error
	return
}

func (dao *TaskDao) GetTaskById(tid, uid uint64) (r *model.Task, err error) {
	r = &model.Task{}
	err = dao.Model(&model.Task{}).Where("id = ? and uid = ?", tid, uid).Find(r).Error
	return
}

func (dao *TaskDao) DeleteTask(tid, uid uint64) (err error) {
	err = dao.Model(&model.Task{}).Where("id = ? and uid = ?", tid, uid).Delete(&model.Task{}).Error
	return
}

func (dao *TaskDao) UpdateTask(req *api.TaskRequest) (err error) {
	r := &model.Task{}
	err = dao.Model(&model.Task{}).Where("id = ? and uid = ?", req.Id, req.Uid).Find(r).Error
	if err != nil {
		return
	}
	r.Title = req.Title
	r.Content = req.Content
	r.Status = int(req.Status)
	return dao.Save(r).Error
}
