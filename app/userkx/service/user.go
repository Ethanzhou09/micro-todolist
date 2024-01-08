package service

import (
	"errors"
	"sync"
	"todolist/app/userkx/repository/db/dao"
	"todolist/app/userkx/repository/db/model"
	"todolist/idl/kitex_gen/api"
	"todolist/pkg/e"
)
import "context"

type Usersrv struct{}

var UsersrvIns *Usersrv
var UsersrvOnce sync.Once

// GetUserSrv 获取单例
func GetUserSrv() *Usersrv {
	// 通过sync.Once实现单例
	UsersrvOnce.Do(func() {
		UsersrvIns = &Usersrv{}
	})
	return UsersrvIns
}

// 饿汉式单例
func GetUserSrvHungury() *Usersrv {
	if UsersrvIns == nil {
		UsersrvIns = &Usersrv{}
	}
	return UsersrvIns
}

func (u *Usersrv) UserLogin(ctx context.Context, req *api.UserRequest) (resp *api.UserResponse, err error) {
	// 查看用户是否存在
	resp = &api.UserResponse{}
	resp.Code = e.Success
	user, err := dao.NewUserDao(ctx).FindUserByUsername(req.Username)
	if err != nil {
		return
	}
	if user.ID == 0 {
		err = errors.New("用户不存在")
		resp.Code = e.Error
		return
	}
	// 验证密码是否正确
	if user.CheckPassword(req.Password) == false {
		err = errors.New("密码错误")
		resp.Code = e.Error
		return
	}
	resp.User = BuildUser(user)
	return
}

func (u *Usersrv) UserRegister(ctx context.Context, req *api.UserRequest) (resp *api.UserResponse, err error) {
	resp = &api.UserResponse{}
	resp.Code = e.Success
	// 验证两次输入密码一致
	if req.Password != req.PasswordConfirm {
		err = errors.New("两次输入密码不一致")
		resp.Code = e.Error
		return
	}
	user, err := dao.NewUserDao(ctx).FindUserByUsername(req.Username)
	if err != nil {
		return
	}
	if user.ID > 0 {
		err = errors.New("用户已存在")
		resp.Code = e.Error
		return
	}
	user = &model.User{
		Username: req.Username,
	}
	// 加密密码
	err = user.SetPassword(req.Password)
	if err != nil {
		resp.Code = e.Error
		return
	}
	// 创建用户
	if err = dao.NewUserDao(ctx).CreateUser(user); err != nil {
		resp.Code = e.Error
		return
	}
	return
}

func BuildUser(item *model.User) *api.UserModel {
	return &api.UserModel{
		Id:        int64(item.ID),
		Username:  item.Username,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
}
