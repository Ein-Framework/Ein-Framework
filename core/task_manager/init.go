package taskmanager

import (
	"github.com/Ein-Framework/Ein-Framework/core/services"
	"github.com/Ein-Framework/Ein-Framework/core/templating"
	"go.uber.org/zap"
)

func New(templateManager templating.ITemplateManager, services *services.Services, logger *zap.Logger) ITaskManager {
	return &TaskManager{
		templateManager: templateManager,
		coreServices:    services,
		executions:      make(map[uint]*ConcurrentExecution),
		logger:          logger,
	}
}
