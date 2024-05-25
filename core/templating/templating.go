package templating

import (
	"os"

	"github.com/Ein-Framework/Ein-Framework/core/services"
	"github.com/Ein-Framework/Ein-Framework/pkg/config"
	"github.com/Ein-Framework/Ein-Framework/pkg/plugins"
	"go.uber.org/zap"
)

func New(cfg *config.Config, services *services.Services, logger *zap.Logger) ITemplateManager {
	pluginsManager := plugins.New(cfg)
	pluginsManager.LoadAllPlugins()

	_, err := os.ReadDir(cfg.TemplatesDir)
	if os.IsNotExist(err) {
		os.Mkdir(cfg.TemplatesDir, os.ModePerm)
	}

	logger.Info("Templating service started")
	return &TemplatingManager{
		coreServices:   services,
		PluginsManager: pluginsManager,
	}
}
