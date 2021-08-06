package game

import (
	"github.com/jinzhu/gorm"
)

// FindGameByExternalID -
// Search for a particular game
func FindGameByExternalID(db *gorm.DB, ExternalID string) (Model, error) {
	game := Model{ExternalID: ExternalID}
	err := db.First(&game).Error
	return game, err
}

// InsertGame -
// Inserts a new game into the database
func InsertGame(db *gorm.DB, game *Model) error {
	return db.Create(game).Error
}

// SaveGame -
// Updates the game into the DB
func SaveGame(db *gorm.DB, game *Model) error {
	return db.Save(game).Error
}
