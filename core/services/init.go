package services

import (
	"github.com/Ein-Framework/Ein-Framework/core/domain"
	"github.com/Ein-Framework/Ein-Framework/pkg/config"
	"go.uber.org/zap"
)

func InitServices(db *domain.ORMConnection, logger *zap.Logger, config *config.Config) *Services {
	context := BuildContext(db, logger, config)

	return &Services{
		AssessmentService: NewAssessmentService(context),
		TaskService:       NewTaskService(context),
		JobService:        NewJobService(context),
	}
}
