package game

import (
	"github.com/jinzhu/gorm"
)

// ListGames -
// Lists all of the games sorted by score
func ListGames(db *gorm.DB, count int) ([]Model, error) {
	var games []Model
	err := db.Limit(count).Order("click_score").Find(&games).Error
	return games, err
}

// FindGameByExternalID -
// Search for a particular game
func FindGameByExternalID(db *gorm.DB, ExternalID string) (Model, error) {
	var game Model
	err := db.Where(Model{ExternalID: ExternalID}).First(&game).Error
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
