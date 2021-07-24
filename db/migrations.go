package db

import (
	"go-friends/models"
)

// AutoMigrate - Run automatic migrations for every model
func (db DB) AutoMigrate() {
	db.conn.AutoMigrate(&models.User{})
	// Add other models here if needed
}
