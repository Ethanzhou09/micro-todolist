package dao

import (
	"context"
	"gorm.io/gorm"
	"todolist/app/taskkitex/repository/db/model"
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
	return dao.Model(&model.Task{}).Create(&data).Error
}
