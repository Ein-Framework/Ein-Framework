package plugins

import (
	"errors"
	"log"
)

func checkPluginType(pluginType int) {
	if pluginType < Scanner || pluginType > ChangeDetector {
		pParseError("Invalid Login Type")
	}
}

func pParseError(errorMessage string) {
	log.Panicln("Error: Failed to parse", errorMessage)
}

func eParseError(errorMessage string) error {
	return errors.New(errorMessage)
}
