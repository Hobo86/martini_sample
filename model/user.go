package model

import (
	"time"
)

type User struct {
	Id       int64
	Name     string
	Age      int64
	Birthday time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
