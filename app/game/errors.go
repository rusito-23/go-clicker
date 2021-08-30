package game

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-clicker/app/common"
)

// ErrorBuilder -
// Util to create the error responses
type ErrorBuilder struct {
	L *i18n.Localizer
}

// FailedToCreateGame -
func (builder ErrorBuilder) FailedToCreateGame() gin.H {
	return gin.H{"message": common.Localize("FailedToCreateGame", builder.L)}
}

// FailedToUpdateGame -
func (builder ErrorBuilder) FailedToUpdateGame() gin.H {
	return gin.H{"message": common.Localize("FailedToUpdateGame", builder.L)}
}

// GameNotFound -
func (builder ErrorBuilder) GameNotFound() gin.H {
	return gin.H{"message": common.Localize("GameNotFound", builder.L)}
}

// InvalidGameState -
func (builder ErrorBuilder) InvalidGameState() gin.H {
	return gin.H{"message": common.Localize("InvalidGameState", builder.L)}
}

// InvalidCountParameter -
func (builder ErrorBuilder) InvalidCountParameter() gin.H {
	return gin.H{"message": common.Localize("InvalidCountParameter", builder.L)}
}

// FailedToRetrieveGames -
func (builder ErrorBuilder) FailedToRetrieveGames() gin.H {
	return gin.H{"message": common.Localize("FailedToRetrieveGames", builder.L)}
}
