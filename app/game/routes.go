package game

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"go-clicker/app/common"
	"net/http"
)

// RegisterRoutes -
// Registers all of the handlers for the `/game` path
func RegisterRoutes(group *gin.RouterGroup) {
	userRoutes := group.Group(
		"/game",
		// Define the possible middlewares for the `game` path
		ErrorBuilderMiddleware(),
	)
	{
		// Define the possible routes for the `game` path
		userRoutes.GET("/ping", ping)
		userRoutes.POST("/", Create)
		userRoutes.GET("/:external_id", Retrieve)
		userRoutes.PUT("/:external_id/click", Click)
	}
}

// Handle ping
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// Create - Handle post request
func Create(c *gin.Context) {
	db := c.MustGet(common.KContextDB).(*gorm.DB)
	errBuilder := c.MustGet(common.KContextErrorBuilder).(ErrorBuilder)

	// Create game with default values
	game := Model{
		ExternalID: uuid.New().String(),
		ClickScore: 0,
		Status:     Created,
	}

	// Insert game into the DB
	err := InsertGame(db, &game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errBuilder.FailedToCreateGame())
		return
	}

	// Serialize created game
	serializer := Serializer{c, game}
	c.JSON(http.StatusCreated, gin.H{"game": serializer.Response()})
}

// Retrieve - Handle retrieve request
func Retrieve(c *gin.Context) {
	db := c.MustGet(common.KContextDB).(*gorm.DB)
	errBuilder := c.MustGet(common.KContextErrorBuilder).(ErrorBuilder)

	// Retrieve game or send error
	game, err := FindGameByExternalID(db, c.Param("external_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, errBuilder.GameNotFound())
		return
	}

	// Serialize game
	serializer := Serializer{c, game}
	c.JSON(http.StatusOK, gin.H{"game": serializer.Response()})
}

// Click - Handle click request
func Click(c *gin.Context) {
	db := c.MustGet(common.KContextDB).(*gorm.DB)
	errBuilder := c.MustGet(common.KContextErrorBuilder).(ErrorBuilder)

	// Retrieve game or send error
	game, err := FindGameByExternalID(db, c.Param("external_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, errBuilder.GameNotFound())
		return
	}

	// Validate and update the game state if needed
	if game.Status == Finished {
		c.JSON(http.StatusPreconditionFailed, errBuilder.InvalidGameState())
		return
	} else if game.Status == Created {
		game.Status = Started
	}

	// Increment game score
	game.ClickScore++

	// Update the game in the DB
	err = SaveGame(db, &game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errBuilder.FailedToUpdateGame())
		return
	}

	// Serialize game
	serializer := Serializer{c, game}
	c.JSON(http.StatusOK, gin.H{"game": serializer.Response()})
}
