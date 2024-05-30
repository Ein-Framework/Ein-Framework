package taskmanager

import (
	"github.com/Ein-Framework/Ein-Framework/core/services"
	"github.com/Ein-Framework/Ein-Framework/core/templating"
)

type TaskManager struct {
	templateManager templating.ITemplateManager
	coreServices    *services.Services
}

type ITaskManager interface {
	ExecuteTemplate(TemplateID uint)
	ExecuteJob(jobID uint)
	CancelJob(jobID uint)
	ViewJobStatus(jobID uint)
}
