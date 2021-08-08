package common

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// DBMiddleware -
// Injects the active DB into the context
func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(KContextDB, db)
		c.Next()
	}
}

// LocMiddleware -
// Injects the localizer into the context
func LocMiddleware(bundle *i18n.Bundle) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.Request.FormValue("lang")
		accept := c.Request.Header.Get("Accept-Language")
		localizer := i18n.NewLocalizer(bundle, lang, accept)
		c.Set(KContextLoc, localizer)
		c.Next()
	}
}
