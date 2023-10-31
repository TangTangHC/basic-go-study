package wire

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("dsc"))
	if err != nil {
		return nil
	}
	return db
}
