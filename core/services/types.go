package services

import (
	"github.com/Ein-Framework/Ein-Framework/core/domain"
	"github.com/Ein-Framework/Ein-Framework/pkg/repository"
	"go.uber.org/zap"
)

type Context struct {
	Logger        *zap.Logger
	OrmConnection *domain.ORMConnection
}

func BuildContext(orm *domain.ORMConnection, logger *zap.Logger) Context {

	return Context{
		Logger:        logger,
		OrmConnection: orm,
	}

}

type Service struct {
	ormConnection *domain.ORMConnection
	repo          repository.TransactionRepository
	logger        *zap.Logger
}

