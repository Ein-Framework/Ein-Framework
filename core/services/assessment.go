package services

import (
	"github.com/Ein-Framework/Ein-Framework/core/domain"
	"github.com/Ein-Framework/Ein-Framework/pkg/repository"
)

type AssessmentService struct {
	ORMConnection *domain.ORMConnection
	repo          repository.TransactionRepository
}

func NewAssessmentService(ctx Context) *AssessmentService {

	repo := repository.NewGormRepository(ctx.Db, ctx.Logger)
	return &AssessmentService{
		repo:          repo,
		ORMConnection: ctx.OrmConnection,
	}
}

func AddNewAssessment() {

}

func DeleteAssessment() {

}

func ModifyAssessment() {

}
