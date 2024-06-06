package services

import (
	"fmt"

	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/pkg/repository"
	"gorm.io/gorm"
)

type AlertService struct {
	Service
}

func NewAlertService(ctx Context) *AlertService {
	repo := repository.NewGormRepository(
		ctx.OrmConnection.Db,
		ctx.Logger.Sugar(),
	)
	return &AlertService{
		Service{
			repo:          repo,
			ormConnection: ctx.OrmConnection,
			logger:        ctx.Logger,
		},
	}
}

func (s *AlertService) AddNewAlerts(alerts ...entity.Alert) []error {
	errs := make([]error, 0)
	for _, alert := range alerts {
		if err := s.repo.Create(&alert); err != nil {
			errs = append(errs, fmt.Errorf("failed to create alert: %w", err))
		}
	}
	return errs
}

func (s *AlertService) AddNewAlert(title string, description string, scope string) (*entity.Alert, error) {
	if title == "" || description == "" || scope == "" {
		return nil, fmt.Errorf("invalid input: title, description & scope are required")
	}

	alert := &entity.Alert{
		Title:       title,
		Description: description,
		Scope:       scope,
	}

	if err := s.repo.Create(alert); err != nil {
		return nil, fmt.Errorf("failed to create alert: %w", err)
	}
	return alert, nil
}

func (s *AlertService) DeleteAlert(id uint) error {
	alert := &entity.Alert{Model: gorm.Model{ID: id}}
	if err := s.repo.Delete(alert); err != nil {
		return fmt.Errorf("failed to delete alert with ID %d: %w", id, err)
	}
	return nil
}

func (s *AlertService) UpdateAlert(id uint, updatedJob *entity.Alert) error {
	job, err := s.GetAlertById(id)
	if err != nil {
		return fmt.Errorf("alert with ID %d not found: %w", id, err)
	}

	job.Title = updatedJob.Title
	job.Description = updatedJob.Description
	job.Scope = updatedJob.Scope

	if err := s.repo.Save(job); err != nil {
		return fmt.Errorf("failed to update job with ID %d: %w", id, err)
	}

	return nil
}

func (s *AlertService) GetAlertById(id uint) (*entity.Alert, error) {
	var alert entity.Alert
	if err := s.repo.GetOneByID(&alert, id); err != nil {
		return nil, fmt.Errorf("alert with ID %d not found: %w", id, err)
	}
	return &alert, nil
}

func (s *AlertService) GetAllAlerts() (*[]entity.Alert, error) {
	var alerts []entity.Alert
	if err := s.repo.GetAll(&alerts); err != nil {
		return nil, fmt.Errorf("failed to retrieve alerts: %w", err)
	}
	return &alerts, nil
}
