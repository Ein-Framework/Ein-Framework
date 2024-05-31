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
	ExecuteTemplate(Template entity.Template, AssesementID uint) (*entity.Task, error)
	ExecuteJob(jobID uint, AssesementID uint) (*entity.JobExecution, error)
	CancelJob(jobID uint) error
	ViewJobStatus(jobID uint) (*entity.TaskState, error)
}
