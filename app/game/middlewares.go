package game

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-clicker/app/common"
)

// ErrorBuilderMiddleware -
func ErrorBuilderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		loc := c.MustGet(common.KContextLoc).(*i18n.Localizer)
		c.Set(common.KContextErrorBuilder, ErrorBuilder{loc: loc})
	}
}
