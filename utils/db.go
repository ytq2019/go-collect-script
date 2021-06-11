package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	dsn := "linqi_yingtejia_:7hweTfH4e3BHFaxP@tcp(122.51.33.239:3306)/linqi_yingtejia_?charset=utf8mb4&parseTime=True&loc=Local"
	db, err1 := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err1 != nil {
		panic(err1)
	}
	return db
}
