package users

import (
	"github.com/jinzhu/gorm"
)

// FindUserByUsername -
// Search for a particular user
func FindUserByUsername(db *gorm.DB, username string) UserModel {
	var user UserModel
	db.First(&user, username)
	return user
}

// CreateNewUser -
// Inserts a new user into the database
func CreateNewUser(db *gorm.DB, user UserModel) {
	db.Create(user)
}
