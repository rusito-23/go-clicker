package users

import (
	"github.com/jinzhu/gorm"
)

// UserModel -
// The user structure that will be saved in the DB
type UserModel struct {
	gorm.Model

	Username  string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Bio       string
}
