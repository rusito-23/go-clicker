package db

import (
	"github.com/caarlos0/env/v6"
	"github.com/jinzhu/gorm"
)

// DB - Wrapper around the database connection
type DB struct {
	conn *gorm.DB
}

// Initialize - Creates the database connection
func Initialize() (DB, error) {

	// Initialize
	db := DB{}
	cfg := CFG{}

	// Parse the database configuration
	err := env.Parse(&cfg)
	if err != nil {
		return db, err
	}

	// Open the `gorm` connection
	conn, err := gorm.Open(cfg.Kind, cfg.Formatted())
	if err != nil {
		return db, err
	}

	db.conn = conn
	return db, nil
}

// Close - Finalize the database connection
func (db DB) Close() {
	db.conn.Close()
}
