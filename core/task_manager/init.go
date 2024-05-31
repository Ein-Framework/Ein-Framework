package taskmanager

import (
	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/core/services"
	"github.com/Ein-Framework/Ein-Framework/core/templating"
	"github.com/Ein-Framework/Ein-Framework/pkg/queue"
)

func New(templateManager templating.ITemplateManager, services *services.Services) ITaskManager {
	return &TaskManager{
		templateManager: templateManager,
		coreServices:    services,
		q:               queue.CreateQueue[entity.Task](),
	}
}
