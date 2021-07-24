package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-friends/db"
	"go-friends/routes"
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

	// Create router
	engine := gin.Default()

	// Register sub-routes

	v1 := engine.Group("/v1")
	{
		routes.RegisterUserRoutes(v1)
	}

	// Start listening
	log.Fatal(engine.Run())
}
