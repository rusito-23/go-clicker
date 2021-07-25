package users

import (
	"github.com/jinzhu/gorm"
)

// AutoMigrate -
// Migration logic for the user model
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&UserModel{})
}
