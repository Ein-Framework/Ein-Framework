package domain

import (
	"fmt"

	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/pkg/config"
	"github.com/Ein-Framework/Ein-Framework/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ORMConnection struct {
	Db *gorm.DB
}

func (orm *ORMConnection) Seed(db *gorm.DB) {

}

func NewDatabase(cfg config.DatabaseConfig) (*ORMConnection, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.LogError("[-] error starting database")
		return nil, err
	}
	if err = db.AutoMigrate(
		&entity.Alert{},
		&entity.Task{},
		&entity.Assessment{},
		&entity.AssessmentStage{},
		&entity.EngagementRules{},
		&entity.TestCredentials{},
		&entity.Job{},
	); err != nil {
		log.LogError("[-] error: migrating database")
		return nil, err
	}

	return &ORMConnection{Db: db}, nil
}
