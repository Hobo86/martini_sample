package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DB() gorm.DB {
	sqlConnection := "root:123456@tcp(127.0.0.1:3306)/im?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic(err)
	}

	return db
}
