package models

import (
	"github.com/jinzhu/gorm"
)

// The User Model
type User struct {
	gorm.Model

	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio"`
	CreatedAt string `json:"created_at"`
}
