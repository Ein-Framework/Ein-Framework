package plugins

import (
	"errors"
	"log"
	"os"
)

func validatePlugin(pluginInstance IPlugin) {
	// meta := pluginInstance.MetaInfo()
	if pluginInstance.Info().Protocol == "" {
		pParseError("[-] error: plugin protocol is missing")
	}

	// checkPluginType(meta.Type)
}

func validateLibraryPath(libraryPath string) error {
	info, err := os.Lstat(libraryPath)
	if err != nil {
		if os.IsNotExist(err) {
			return eParseError("[-] error: plugin not found")
		}
		return err
	}

	if info.IsDir() {
		return eParseError("[-] error: bad file path, path is a directory")
	}

	return nil
}

// func checkPluginType(pluginType PluginType) {
// 	if pluginType <= UndefinedType || pluginType >= UndefinedType {
// 		pParseError("[-] error: Invalid Login Type")
// 	}
// }

func pParseError(errorMessage string) {
	log.Panicln("[-] error: Failed to parse", errorMessage)
}

func eParseError(errorMessage string) error {
	return errors.New(errorMessage)
}
