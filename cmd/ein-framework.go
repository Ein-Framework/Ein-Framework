package main

import (
	"log"

	"github.com/Ein-Framework/Ein-Framework/core/config"
	"github.com/Ein-Framework/Ein-Framework/core/plugins"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Panicln("Error: loading config", err)
	}

	pluginManager := plugins.CreatePluginManager(cfg)
	pluginManager.LoadAllPlugins()
}
