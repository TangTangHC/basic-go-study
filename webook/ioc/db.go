package ioc

import (
	"github.com/TangTangHC/basic-go-study/webook/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	if err != nil {
		panic("数据库初始化失败")
	}
	return db
}
