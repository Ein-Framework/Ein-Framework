package services

import (
	"fmt"

	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/pkg/repository"
	"gorm.io/gorm"
)

type JobService struct {
	Service
}

func NewJobService(ctx Context) *JobService {
	repo := repository.NewGormRepository(
		ctx.OrmConnection.Db,
		ctx.Logger.Sugar(),
	)
	return &JobService{
		Service{
			repo:          repo,
			ormConnection: ctx.OrmConnection,
			logger:        ctx.Logger,
		},
	}
}

func (s *JobService) AddNewJob(name string, templates []entity.Template) (*entity.Job, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid input: name is required")
	}

	job := &entity.Job{
		Name:      name,
		Templates: templates,
	}

	if err := s.repo.Create(job); err != nil {
		return nil, fmt.Errorf("failed to create job: %w", err)
	}

	return job, nil
}

func (s *JobService) DeleteJob(id uint) error {
	job := &entity.Job{Model: gorm.Model{ID: id}}
	if err := s.repo.Delete(job); err != nil {
		return fmt.Errorf("failed to delete job with ID %d: %w", id, err)
	}
	return nil
}

func (s *JobService) UpdateJob(id uint, updatedJob *entity.Job) error {
	job, err := s.GetJobById(id)
	if err != nil {
		return fmt.Errorf("job with ID %d not found: %w", id, err)
	}

	job.Name = updatedJob.Name
	job.Templates = updatedJob.Templates

	if err := s.repo.Save(job); err != nil {
		return fmt.Errorf("failed to update job with ID %d: %w", id, err)
	}

	return nil
}

func (s *JobService) GetJobById(id uint) (*entity.Job, error) {
	var job entity.Job
	if err := s.repo.GetOneByID(&job, id); err != nil {
		return nil, fmt.Errorf("job with ID %d not found: %w", id, err)
	}
	return &job, nil
}

func (s *JobService) GetAllJobs() ([]*entity.Job, error) {
	var jobs []*entity.Job
	if err := s.repo.GetAll(&jobs); err != nil {
		return nil, fmt.Errorf("failed to retrieve jobs: %w", err)
	}
	return jobs, nil
}
