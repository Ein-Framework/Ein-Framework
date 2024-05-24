package services

import (
	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/pkg/repository"
	"gorm.io/gorm"
)

type AssessmentService struct {
	Service
}

func NewAssessmentService(ctx Context) *AssessmentService {

	repo := repository.NewGormRepository(ctx.OrmConnection.Db, ctx.Logger.Sugar())
	return &AssessmentService{
		Service{
			repo:          repo,
			ormConnection: ctx.OrmConnection,
			logger:        ctx.Logger,
		},
	}
}
func (s *AssessmentService) AddNewAssessment(name string, assessmentType entity.AssessmentType, scope entity.Scope) (*entity.Assessment, error) {
	assessment, err := entity.NewAssessment(name, assessmentType, scope, s.repo)
	if err != nil {
		return nil, err
	}

	err = s.repo.Create(assessment)
	if err != nil {
		return nil, err
	}

	return assessment, nil
}

func (s *AssessmentService) DeleteAssessment(id uint) error {
	assessment := &entity.Assessment{Model: gorm.Model{ID: id}}
	return s.repo.Delete(assessment)
}

func (s *AssessmentService) UpdateAssessment(id uint, updatedAssessment *entity.Assessment) error {
	assessment, err := s.GetAssessmentById(id)
	if err != nil {
		return err
	}

	assessment.Name = updatedAssessment.Name
	assessment.Type = updatedAssessment.Type
	assessment.Scope = updatedAssessment.Scope
	assessment.Assets = updatedAssessment.Assets
	assessment.Stage = updatedAssessment.Stage
	assessment.EngagementRules = updatedAssessment.EngagementRules
	assessment.Jobs = updatedAssessment.Jobs
	assessment.Reports = updatedAssessment.Reports

	return s.repo.Save(assessment)
}

func (s *AssessmentService) GetAssessmentById(id uint) (*entity.Assessment, error) {
	var assessment entity.Assessment
	err := s.repo.GetOneByID(&assessment, id)
	if err != nil {
		return nil, err
	}
	return &assessment, nil
}
