package services

import (
	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/pkg/repository"
	"gorm.io/gorm"
)

type AssessmentStageService struct {
	Service
}

func NewAssessmentStageService(ctx Context) *AssessmentStageService {
	repo := repository.NewGormRepository(
		ctx.OrmConnection.Db,
		ctx.Logger.Sugar(),
	)
	return &AssessmentStageService{
		Service: Service{
			repo:          repo,
			ormConnection: ctx.OrmConnection,
			logger:        ctx.Logger,
		},
	}
}

func (s *AssessmentStageService) InitStages() error {
	stages, err := s.GetAllAssessmentStages()
	if err != nil {
		return err
	}

	if len(*stages) > 0 {
		return nil
	}

	_, err = s.AddNewAssessmentStage(entity.ReconnaissanceStage, "", false)
	if err != nil {
		return nil
	}
	_, err = s.AddNewAssessmentStage(entity.ScanningStage, "", false)
	if err != nil {
		return nil
	}
	_, err = s.AddNewAssessmentStage(entity.AutomatedChecking, "", false)
	if err != nil {
		return nil
	}
	_, err = s.AddNewAssessmentStage(entity.MappingStage, "", false)
	if err != nil {
		return nil
	}
	_, err = s.AddNewAssessmentStage(entity.ExploitationStage, "", false)
	if err != nil {
		return nil
	}
	_, err = s.AddNewAssessmentStage(entity.ReportingStage, "", true)

	return err
}

func (s *AssessmentStageService) GetAllAssessmentStages() (*[]entity.AssessmentStage, error) {
	var assessmentStages []entity.AssessmentStage

	err := s.repo.GetAll(&assessmentStages)
	if err != nil {
		return nil, err
	}

	return &assessmentStages, nil
}

func (s *AssessmentStageService) DeleteAssessmentStage(id uint) error {
	return s.repo.Delete(&entity.AssessmentStage{Model: gorm.Model{ID: id}})
}

func (s *AssessmentStageService) AddNewAssessmentStage(name string, description string, completed bool) (*entity.AssessmentStage, error) {
	assessmentStage := &entity.AssessmentStage{
		Name:        name,
		Description: description,
		Completed:   completed,
	}

	err := s.repo.Create(assessmentStage)
	if err != nil {
		return nil, err
	}

	return assessmentStage, nil
}

func (s *AssessmentStageService) UpdateAssessmentStage(id uint, updatedAssessment *entity.AssessmentStage) error {
	assessmentStage, err := s.GetAssessmentStageById(id)
	if err != nil {
		return err
	}

	assessmentStage.Name = updatedAssessment.Name
	assessmentStage.Description = updatedAssessment.Description
	assessmentStage.Completed = updatedAssessment.Completed

	return s.repo.Save(assessmentStage)
}

func (s *AssessmentStageService) GetAssessmentStageById(id uint) (*entity.AssessmentStage, error) {
	var assessment entity.AssessmentStage
	err := s.repo.GetOneByID(&assessment, id)
	if err != nil {
		return nil, err
	}
	return &assessment, nil
}
