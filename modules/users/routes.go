package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-friends/modules/common"
	"net/http"
)

// RegisterRoutes -
// Links all of the handlers in the `/users` path
func RegisterRoutes(group *gin.RouterGroup) {
	userRoutes := group.Group("/users")
	{
		userRoutes.GET("/ping", pingUser)
		userRoutes.POST("/", postUser)
	}
}

// Ping Route Test
func pingUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

// Creates a new user with the data given in the payload
func postUser(c *gin.Context) {
	userModelValidator := UserModelValidator{}
	db := c.MustGet(common.KContextDB).(*gorm.DB)

	user, err := userModelValidator.Bind(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	CreateNewUser(db, user)
	c.JSON(200, gin.H{"message": "created"})
}
