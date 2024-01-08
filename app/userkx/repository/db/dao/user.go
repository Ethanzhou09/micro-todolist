package dao

import (
	"context"
	"gorm.io/gorm"
	"todolist/app/userkx/repository/db/model"
)

// 定义user的curd操作
type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) FindUserByUsername(username string) (r *model.User, err error) {
	r = &model.User{}
	err = dao.Where("username = ?", username).Find(r).Error
	return r, err
}
func (dao *UserDao) CreateUser(in *model.User) (err error) {
	return dao.Model(&model.User{}).Create(&in).Error

}
