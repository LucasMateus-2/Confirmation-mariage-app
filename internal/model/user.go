// internal/model/user.go
package model

type User struct {
	ID       int64  `gorm:"primaryKey"                   json:"id"`
	Email    string `gorm:"column:email;not null;unique" json:"email"`
	Password string `gorm:"column:password;not null"     json:"-"`
}
