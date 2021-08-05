package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-clicker/db"
	"go-clicker/modules/common"
	"go-clicker/modules/game"
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
	game.AutoMigrate(db)

	// Create router
	engine := gin.Default()

	// Set up middlewares
	engine.Use(common.MiddlewareDB(db))

	// Create `v1` group &&
	// Register sub-routes
	v1 := engine.Group("/v1")
	{
		game.RegisterRoutes(v1)
	}

	// Start listening
	log.Fatal(engine.Run())
}
