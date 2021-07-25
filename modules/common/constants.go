package common

// KContextKey -
// Keys for middle wares context values
type KContextKey string

const (
	// KContextDB - Pass the gorm DB reference through the MiddlewareDB
	KContextDB = "GIN_CONTEXT_DB"
)
