package game

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"go-clicker/modules/common"
	"net/http"
)

// RegisterRoutes -
// Links all of the handlers in the `/game` path
func RegisterRoutes(group *gin.RouterGroup) {
	userRoutes := group.Group("/game")
	{
		userRoutes.GET("/ping", ping)
		userRoutes.POST("/", create)
		userRoutes.GET("/:external_id", retrieve)
		userRoutes.PUT("/:external_id/click", click)
	}
}

// Handle ping
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// Handle post request
func create(c *gin.Context) {
	db := c.MustGet(common.KContextDB).(*gorm.DB)

	// Create game with default values
	game := Model{
		ExternalID: uuid.New().String(),
		ClickScore: 0,
		Status:     Created,
	}

	// Insert game into the DB
	err := InsertGame(db, &game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save game"})
		return
	}

	// Serialize created game
	serializer := Serializer{c, game}
	c.JSON(http.StatusCreated, gin.H{"game": serializer.Response()})
}

// Handle retrieve request
func retrieve(c *gin.Context) {
	db := c.MustGet(common.KContextDB).(*gorm.DB)

	// Retrieve game or send error
	game, err := FindGameByExternalID(db, c.Param("external_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Game Not Found"})
		return
	}

	// Serialize game
	serializer := Serializer{c, game}
	c.JSON(http.StatusOK, gin.H{"game": serializer.Response()})
}

// Handle click request
func click(c *gin.Context) {
	db := c.MustGet(common.KContextDB).(*gorm.DB)

	// Retrieve game or send error
	game, err := FindGameByExternalID(db, c.Param("external_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Game Not Found"})
		return
	}

	// Increment game score & start progress
	game.ClickScore++
	game.Status = Started

	// Update the game in the DB
	err = SaveGame(db, &game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save game"})
		return
	}

	// Serialize game
	serializer := Serializer{c, game}
	c.JSON(http.StatusOK, gin.H{"game": serializer.Response()})
}
