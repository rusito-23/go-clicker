package db

import (
	"fmt"
)

// CFG - contains the DB environment specs
type CFG struct {
	Username string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Database string `env:"POSTGRES_DB"`
	Kind     string `env:"POSTGRES_KIND" envDefault:"postgres"`
	Port     int    `env:"POSTGRES_PORT" envDefault:"5432"`
	Host     string `env:"POSTGRES_HOST" envDefault:"database"`
	SSLMode  string `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
}

// Formatted - Parses the database configuration into `gorm` standards
func (cfg CFG) Formatted() string {
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
