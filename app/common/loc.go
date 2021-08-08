package common

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// CreateLocalizationBundle -
func CreateLocalizationBundle() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	registerMessageFiles(bundle)
	return bundle
}

// Register messages files
func registerMessageFiles(bundle *i18n.Bundle) {
	bundle.MustLoadMessageFile(KEnglishLocPath)
	// Add more resources files here if needed
}

// Localize -
// Create the localization configuration and run in a single step
func Localize(message string, localizer *i18n.Localizer) string {
	// Create and run the localization configuration
	config := i18n.LocalizeConfig{MessageID: message}
	localized, err := localizer.Localize(&config)

	// If the localization fails we just return the localization key
	if err != nil {
		return message
	}

	return localized
}
