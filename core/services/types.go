package services

import (
	"github.com/Ein-Framework/Ein-Framework/core/domain"
	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/pkg/config"
	"github.com/Ein-Framework/Ein-Framework/pkg/repository"
	"go.uber.org/zap"
)

type Context struct {
	Logger        *zap.Logger
	OrmConnection *domain.ORMConnection
	Config        *config.Config
}

func BuildContext(orm *domain.ORMConnection, logger *zap.Logger, config *config.Config) Context {
	return Context{
		Logger:        logger,
		OrmConnection: orm,
		Config:        config,
	}
}

type Service struct {
	ormConnection *domain.ORMConnection
	repo          repository.TransactionRepository
	logger        *zap.Logger
}

type Services struct {
	AssessmentService      IAssessmentService
	JobService             IJobService
	TaskService            ITaskService
	JobExecutionService    IJobExecutionService
	AssessmentStageService IAssessmentStageService
	AlertService           IAlertService
}

type IAlertService interface {
	GetAllAlerts() (*[]entity.Alert, error)
	AddNewAlert(title string, description string, scope string) (*entity.Alert, error)
	AddNewAlerts(...entity.Alert) []error
	DeleteAlert(id uint) error
	UpdateAlert(id uint, updatedAssessment *entity.Alert) error
	GetAlertById(id uint) (*entity.Alert, error)
}

type IAssessmentService interface {
	GetAllAssessments() (*[]entity.Assessment, error)
	AddNewAssessmentFromHackerone(url string) (*entity.Assessment, error)
	AddNewAssessment(name string, assessmentType entity.AssessmentType, scope entity.Scope, engagementRules entity.EngagementRules) (*entity.Assessment, error)
	DeleteAssessment(id uint) error
	UpdateAssessment(id uint, updatedAssessment *entity.Assessment) error
	GetAssessmentById(id uint) (*entity.Assessment, error)
}

type IAssessmentStageService interface {
	GetAllAssessmentStages() (*[]entity.AssessmentStage, error)
	AddNewAssessmentStage(name string, description string, completed bool) (*entity.AssessmentStage, error)
	DeleteAssessmentStage(id uint) error
	UpdateAssessmentStage(id uint, updatedAssessment *entity.AssessmentStage) error
	GetAssessmentStageById(id uint) (*entity.AssessmentStage, error)

	InitStages() error
}

type ITaskService interface {
	AddNewTask(state entity.TaskState, template entity.Template, assessmentId uint, assessmentStageId uint) (*entity.Task, error)
	DeleteTask(id uint) error
	DeleteTasks(tasks ...entity.Task) []error
	UpdateTask(id uint, updatedTask *entity.Task) error
	GetTaskById(id uint) (*entity.Task, error)
	GetAllTasks() ([]*entity.Task, error)

	UpdateTaskState(id uint, state entity.TaskState) error
}

type IJobService interface {
	AddNewJob(name string, description string, templates []entity.Template) (*entity.Job, error)
	DeleteJob(id uint) error
	UpdateJob(id uint, updatedJob *entity.Job) error
	GetJobById(id uint) (*entity.Job, error)
	GetAllJobs() ([]*entity.Job, error)

	UpdateJobTemplates(jobId uint, templates ...entity.Template) (*entity.Job, error)
	// AddTemplateToJob(jobId uint, template entity.Template) (*entity.JobTemplate, error)
	// AddTemplatesToJob(jobId uint, template ...entity.Template) []error
}

type IJobExecutionService interface {
	AddNewJobExecution(jobID, assessmentID uint, tasks []entity.Task, status entity.TaskState) (*entity.JobExecution, error)
	DeleteJobExecution(id uint) error
	UpdateJobExecution(id uint, updatedJobExecution *entity.JobExecution) error
	GetJobExecutionById(id uint) (*entity.JobExecution, error)
	GetJobExecutionsByJobId(id uint) ([]*entity.JobExecution, error)
	GetJobExecutionsNotCanceledByJobId(id uint) ([]*entity.JobExecution, error)
	GetAllJobExecutions() ([]*entity.JobExecution, error)

	UpdateJobExecutionStatus(id uint, state entity.TaskState) error
}
