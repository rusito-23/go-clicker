package game

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-clicker/modules/common"
)

// ErrorBuilder -
// Util to create the error responses
type ErrorBuilder struct {
	loc *i18n.Localizer
}

// FailedToCreateGame -
func (builder ErrorBuilder) FailedToCreateGame() gin.H {
	return gin.H{"message": common.Localize("FailedToCreateGame", builder.loc)}
}

// FailedToUpdateGame -
func (builder ErrorBuilder) FailedToUpdateGame() gin.H {
	return gin.H{"message": common.Localize("FailedToUpdateGame", builder.loc)}
}

// GameNotFound -
func (builder ErrorBuilder) GameNotFound() gin.H {
	return gin.H{"message": common.Localize("GameNotFound", builder.loc)}
}

// InvalidGameState -
func (builder ErrorBuilder) InvalidGameState() gin.H {
	return gin.H{"message": common.Localize("InvalidGameState", builder.loc)}
}
