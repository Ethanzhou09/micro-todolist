package dao

import "todolist/app/userkx/repository/db/model"

func migration() {
	// 把模型与数据库中的表对应起来
	_db.Set(`gorm:table_options`, "charset=utf8mb4").AutoMigrate(&model.User{})
}
