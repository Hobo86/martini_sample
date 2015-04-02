package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int64
	Nickname  string
	Gender    int64
	Birthday  time.Time
	CreatedAt time.Time `gorm:"column:create_time"`
	UpdatedAt time.Time `gorm:"column:modify_time"`
}

func (u User) TableName() string {
	return "user"
}

func (u User) GetUserById(id int64, db gorm.DB) *User {
	user := User{}
	db.Where("id = ?", id).Find(&user)
	return &user
}
