package taskmanager

import (
	"github.com/Ein-Framework/Ein-Framework/core/services"
	"github.com/Ein-Framework/Ein-Framework/core/templating"
)

func New(templateManager templating.ITemplateManager, services *services.Services) ITaskManager {
	return &TaskManager{
		templateManager: templateManager,
		coreServices:    services,
		executions:      make(map[uint]*ConcurrentExecution),
	}
}
