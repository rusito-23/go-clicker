package game

import (
	"github.com/jinzhu/gorm"
)

// AutoMigrate -
// Migration logic for the game model
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Model{})
}
