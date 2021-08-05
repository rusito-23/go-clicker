package game

import (
	"github.com/jinzhu/gorm"
)

// FindGameByID -
// Search for a particular game
func FindGameByID(db *gorm.DB, ID uint) Model {
	var game Model
	db.First(&game, ID)
	return game
}

// InsertGame -
// Inserts a new game into the database
func InsertGame(db *gorm.DB, game *Model) {
	db.Create(game)
}
