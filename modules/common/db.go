package common

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/jinzhu/gorm"
)

// InitializeDB -
// Creates the database connection
func InitializeDB() (*gorm.DB, error) {

	// Initialize
	var db *gorm.DB
	cfg := DBConfig{}

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

// DBConfig -
// Contains the DB environment specs
type DBConfig struct {
	Username string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Database string `env:"POSTGRES_DB"`
	Kind     string `env:"POSTGRES_KIND" envDefault:"postgres"`
	Port     int    `env:"POSTGRES_PORT" envDefault:"5432"`
	Host     string `env:"POSTGRES_HOST" envDefault:"database"`
	SSLMode  string `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
}

// Formatted DBConfig -
// Parses the database configuration into `gorm` standards
func (cfg DBConfig) Formatted() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Database,
		cfg.SSLMode,
	)
}
