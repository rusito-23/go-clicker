package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes - Links all of the handlers in the `/users` path
func RegisterUserRoutes(group *gin.RouterGroup) {
	userRoutes := group.Group("/users")
	{
		userRoutes.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
