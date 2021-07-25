package db

import (
	"github.com/caarlos0/env/v6"
	"github.com/jinzhu/gorm"
)

// Initialize -
// Creates the database connection
func Initialize() (*gorm.DB, error) {

	// Initialize
	var db *gorm.DB
	cfg := CFG{}

	// Parse the database configuration
	err := env.Parse(&cfg)
	if err != nil {
		return db, err
	}

	// Open the `gorm` connection
	db, err = gorm.Open(cfg.Kind, cfg.Formatted())
	if err != nil {
		return db, err
	}

	return db, nil
}
