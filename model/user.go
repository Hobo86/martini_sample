package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int64     `form:"id"`
	Nickname  string    `form:"nickname"`
	Gender    int64     `form:"gender"`
	CreatedAt time.Time `form:"create_time"`
	UpdatedAt time.Time `xorm:"update_time"`
}

func (u User) TableName() string {
	return "user"
}

func (u User) GetUserById(id int64, db gorm.DB) *User {
	user := User{}
	db.Where("id = ?", id).Find(&user)
	return &user
}
