package taskmanager

import (
	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/core/services"
	"github.com/Ein-Framework/Ein-Framework/core/templating"
	"github.com/Ein-Framework/Ein-Framework/pkg/queue"
)

type TaskManager struct {
	templateManager templating.ITemplateManager
	coreServices    *services.Services
	q               queue.FifoQueue[entity.Task]
}

type ITaskManager interface {
	ExecuteTemplate(TemplateID uint) (*entity.JobExecution, error)
	ExecuteJob(jobID uint) (*entity.JobExecution, error)
	CancelJob(jobID uint) error
	ViewJobStatus(jobID uint) (*entity.TaskState, error)
}
