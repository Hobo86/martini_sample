package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DB() gorm.DB {
	sqlConnection := "root:meet771027@tcp(127.0.0.1:3306)/im?parseTime=True"
	db, err := gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic(err)
	}

	return db
}
