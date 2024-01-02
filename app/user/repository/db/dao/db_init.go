package dao

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"todolist/config"
)

var _db *gorm.DB

func InitDB() {
	DbHost := config.DbHost
	DbPort := config.DbPort
	DbUser := config.DbUser
	DbPassWord := config.DbPassWord
	DbName := config.DbName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassWord, DbHost, DbPort, DbName)
	fmt.Println(dsn)
	err := Database(dsn)
	if err != nil {
		fmt.Println(err)
	}
}

func Database(connstring string) error {
	var ormLogger logger.Interface = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connstring, // DSN data source name
		DefaultStringSize:         256,        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,      // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	_db = db
	migration()
	return err
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
