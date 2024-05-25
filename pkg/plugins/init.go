package plugins

import (
	"github.com/Ein-Framework/Ein-Framework/pkg/config"
)

func New(cfg *config.Config) IPluginManager {
	return PluginManager{
		config:        cfg,
		loadedPlugins: make(map[string]*LoadedPluginInfo),
	}
}
