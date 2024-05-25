package templating

import (
	"github.com/Ein-Framework/Ein-Framework/core/services"
	"github.com/Ein-Framework/Ein-Framework/pkg/config"
	"github.com/Ein-Framework/Ein-Framework/pkg/plugins"
)

type TemplatingManager struct {
	coreServices   *services.Services
	PluginsManager plugins.IPluginManager
}

func New(cfg *config.Config, services *services.Services) *TemplatingManager {
	return &TemplatingManager{
		coreServices:   services,
		PluginsManager: plugins.New(cfg),
	}
}
