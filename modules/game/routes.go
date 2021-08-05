package game

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-clicker/modules/common"
	"net/http"
)

// RegisterRoutes -
// Links all of the handlers in the `/game` path
func RegisterRoutes(group *gin.RouterGroup) {
	userRoutes := group.Group("/game")
	{
		userRoutes.GET("/ping", pingGame)
		userRoutes.POST("/", postGame)
	}
}

// Ping Route Test
func pingGame(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

// Creates a new game
func postGame(c *gin.Context) {
	db := c.MustGet(common.KContextDB).(*gorm.DB)
	game := Model{
		ClickScore: 0,
		Status:     Created,
	}

	InsertGame(db, &game)
	serializer := Serializer{c, game}
	c.JSON(http.StatusCreated, gin.H{"game": serializer.Response()})
}
