package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 定义数据库模型
type User struct {
	gorm.Model
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

const PasswordCost = 12

func (user *User) SetPassword(password string) error {
	// 用bcrypt对密码进行哈希加密
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	// 比较密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
