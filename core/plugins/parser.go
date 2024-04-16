package plugins

import (
	"errors"
	"log"
)

func checkPluginType(pluginType PluginType) {
	if pluginType <= UndefinedType || pluginType >= UndefinedType {
		pParseError("Invalid Login Type")
	}
}

func pParseError(errorMessage string) {
	log.Panicln("Error: Failed to parse", errorMessage)
}

func eParseError(errorMessage string) error {
	return errors.New(errorMessage)
}
