package plugins

import (
	"errors"
	"log"
	"os"
)

func validatePlugin(pluginInstance IPlugin) {
	meta := pluginInstance.MetaInfo()

	checkPluginType(meta.Type)
}

func validateLibraryPath(libraryPath string) error {
	info, err := os.Lstat(libraryPath)
	if err != nil {
		if os.IsNotExist(err) {
			return eParseError("plugin not found")
		}
		return err
	}

	if info.IsDir() {
		return eParseError("bad file path, path is a directory")
	}

	return nil
}

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
