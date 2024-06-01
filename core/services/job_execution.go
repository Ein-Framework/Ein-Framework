package services

import (
	"fmt"

	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/pkg/repository"
	"gorm.io/gorm"
)

type JobExecutionService struct {
	Service
}

func NewJobExecutionService(ctx Context) *JobExecutionService {
	repo := repository.NewGormRepository(
		ctx.OrmConnection.Db,
		ctx.Logger.Sugar(),
		"Job",
		"Assessment",
		"Tasks",
	)
	return &JobExecutionService{
		Service{
			repo:          repo,
			ormConnection: ctx.OrmConnection,
			logger:        ctx.Logger,
		},
	}
}

func (s *JobExecutionService) AddNewJobExecution(jobID, assessmentID uint, tasks []entity.Task, status entity.TaskState) (*entity.JobExecution, error) {
	if jobID == 0 || assessmentID == 0 || status == "" {
		return nil, fmt.Errorf("invalid input: jobID, assessmentID, and status are required")
	}

	jobExecution := &entity.JobExecution{
		JobID:        jobID,
		AssessmentId: assessmentID,
		Tasks:        tasks,
		Status:       status,
	}

	if err := s.repo.Create(jobExecution); err != nil {
		return nil, fmt.Errorf("failed to create job execution: %w", err)
	}

	return jobExecution, nil
}

func (s *JobExecutionService) DeleteJobExecution(id uint) error {
	jobExecution := &entity.JobExecution{Model: gorm.Model{ID: id}}
	if err := s.repo.Delete(jobExecution); err != nil {
		return fmt.Errorf("failed to delete job execution with ID %d: %w", id, err)
	}
	return nil
}

func (s *JobExecutionService) UpdateJobExecution(id uint, updatedJobExecution *entity.JobExecution) error {
	jobExecution, err := s.GetJobExecutionById(id)
	if err != nil {
		return fmt.Errorf("job execution with ID %d not found: %w", id, err)
	}
	//jobExecution.Tasks = updatedJobExecution.Tasks
	jobExecution.Status = updatedJobExecution.Status

	if err := s.repo.Save(jobExecution); err != nil {
		return fmt.Errorf("failed to update job execution with ID %d: %w", id, err)
	}

	return nil
}

func (s *JobExecutionService) UpdateJobExecutionStatus(id uint, state entity.TaskState) error {
	jobExecution, err := s.GetJobExecutionById(id)
	if err != nil {
		return fmt.Errorf("job execution with ID %d not found: %w", id, err)
	}
	jobExecution.Status = state

	if err := s.repo.Save(jobExecution); err != nil {
		return fmt.Errorf("failed to update job execution with ID %d: %w", id, err)
	}

	return nil
}

func (s *JobExecutionService) GetJobExecutionById(id uint) (*entity.JobExecution, error) {
	var jobExecution entity.JobExecution
	if err := s.repo.GetOneByID(&jobExecution, id); err != nil {
		return nil, fmt.Errorf("job execution with ID %d not found: %w", id, err)
	}
	return &jobExecution, nil
}

func (s *JobExecutionService) GetJobExecutionsByJobId(id uint) ([]*entity.JobExecution, error) {
	var jobExecutions []*entity.JobExecution
	if res := s.repo.DB().Where(&entity.JobExecution{JobID: id}).Find(&jobExecutions); res.Error != nil {
		return nil, fmt.Errorf("job execution with ID %d not found: %w", id, res.Error)
	}
	return jobExecutions, nil
}

func (s *JobExecutionService) GetJobExecutionsNotCanceledByJobId(id uint) ([]*entity.JobExecution, error) {
	var jobExecutions []*entity.JobExecution
	db := s.repo.DB()
	if res := db.Where(&entity.JobExecution{JobID: id}).Not(&entity.JobExecution{Status: entity.Canceled}).Find(&jobExecutions); res.Error != nil {
		return nil, fmt.Errorf("job execution with ID %d not found: %w", id, res.Error)
	}
	return jobExecutions, nil
}

func (s *JobExecutionService) GetAllJobExecutions() ([]*entity.JobExecution, error) {
	var jobExecutions []*entity.JobExecution
	if err := s.repo.GetAll(&jobExecutions); err != nil {
		return nil, fmt.Errorf("failed to retrieve job executions: %w", err)
	}
	return jobExecutions, nil
}
