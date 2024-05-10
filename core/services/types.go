package services

import (
	"github.com/Ein-Framework/Ein-Framework/core/domain"
	"go.uber.org/zap"

	"gorm.io/gorm"
)

type Context struct {
	Logger        *zap.SugaredLogger
	Db            *gorm.DB
	OrmConnection *domain.ORMConnection
}
