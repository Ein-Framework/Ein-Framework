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
	AssessmentService IAssessmentService
	JobService        IJobService
	TaskService       ITaskService
}

type IAssessmentService interface {
	AddNewAssessment(name string, assessmentType entity.AssessmentType, scope entity.Scope) (*entity.Assessment, error)
	DeleteAssessment(id uint) error
	UpdateAssessment(id uint, updatedAssessment *entity.Assessment) error
	GetAssessmentById(id uint) (*entity.Assessment, error)
}
type ITaskService interface {
	AddNewTask(state entity.TaskState, output string, outputFormat entity.OutputFormat, args map[string]string, assessmentStageId uint) (*entity.Task, error)
	DeleteTask(id uint) error
	UpdateTask(id uint, updatedTask *entity.Task) error
	GetTaskById(id uint) (*entity.Task, error)
	GetAllTasks() ([]*entity.Task, error)
}

type IJobService interface {
	AddNewJob(name string, templates []entity.Template) (*entity.Job, error)
	DeleteJob(id uint) error
	UpdateJob(id uint, updatedJob *entity.Job) error
	GetJobById(id uint) (*entity.Job, error)
	GetAllJobs() ([]*entity.Job, error)
}
