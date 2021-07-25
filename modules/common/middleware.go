package common

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// MiddlewareDB -
// Sets the active DB to the context
func MiddlewareDB(activeDB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(KContextDB, activeDB)
		c.Next()
	}
}
