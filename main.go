package main

import (
	// "github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-friends/db"
	"log"
)

func main() {

	// Get database credentials and specs
	var db, err = db.Initialize()

	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
		return
	}

	// Close database on exit
	defer db.Close()

	// Run automatic migrations
	db.AutoMigrate()

	// TODO: Create gin routers
}
