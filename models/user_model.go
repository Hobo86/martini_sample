package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/sessionauth"
)

type User struct {
	Id        int64
	Nickname  string `form:"nickname"`
	Password  string `form:"password"`
	Gender    int64
	Birthday  time.Time
	CreatedAt time.Time `gorm:"column:created_time"`
	UpdatedAt time.Time `gorm:"column:updated_time"`

	authenticated bool `form:"-" db:"-"`
}

// GetAnonymousUser should generate an anonymous user model
// for all sessions. This should be an unauthenticated 0 value struct.
func GenerateAnonymousUser() sessionauth.User {
	return &User{}
}

func (u User) TableName() string {
	return "user"
}

func (u User) GetUserById(id int64, db gorm.DB) *User {
	user := User{}
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil
	}

	return &user
}

func (u User) GetUserByNicknamePwd(nickname string, pwd string, db gorm.DB) *User {
	user := User{}
	if err := db.Where("nickname = ? AND password = ?", nickname, pwd).First(&user).Error; err != nil {
		return nil
	}
	return &user
}

// Login will preform any actions that are required to make a user model
// officially authenticated.
func (u *User) Login() {
	// Update last login time
	// Add to logged-in user's list
	// etc ...
	u.authenticated = true
}

// Logout will preform any actions that are required to completely
// logout a user.
func (u *User) Logout() {
	// Remove from logged-in user's list
	// etc ...
	u.authenticated = false
}

func (u *User) IsAuthenticated() bool {
	return u.authenticated
}

func (u *User) UniqueId() interface{} {
	return u.Id
}

// GetById will populate a user object from a database model with
// a matching id.
func (u *User) GetById(id interface{}) error {
	// err := dbmap.SelectOne(u, "SELECT * FROM users WHERE id = $1", id)
	// if err != nil {
	// 	return err
	// }

	return nil
}
