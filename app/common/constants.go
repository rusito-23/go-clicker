package common

// Context Keys
// Constants that will be used to pass information through the gin context

const (
	// KContextDB - The gorm DB reference
	KContextDB = "GIN_CONTEXT_DB"

	// KContextLoc - The Localizer reference based on the headers
	KContextLoc = "GIN_CONTEXT_LOC"

	// KContextErrorBuilder - The error builder reference
	KContextErrorBuilder = "GIN_CONTEXT_ERROR_BUILDER"
)

// Localization Resources Paths
// Paths to the different localization resources

const (
	// KEnglishLocPath - The default English localization path
	KEnglishLocPath = "resources/en.toml"
	// Add more localization files path here if needed
)
