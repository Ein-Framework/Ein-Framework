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
	fmt.Println(cfg)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.LogError("[-] error starting database")
		return nil, err
	}
	if err = db.AutoMigrate(
		&entity.Attachement{},
		&entity.Asset{},
		&entity.Alert{},
		&entity.Vulnerability{},
		&entity.Scope{},
		&entity.Report{},
		&entity.Job{},
		&entity.JobExecution{},
		&entity.Task{},
		&entity.TestCredentials{},
		&entity.HeaderInjection{},
		&entity.AssessmentStage{},
		&entity.EngagementRules{},
		&entity.Assessment{},
		&entity.User{},
		// &entity.TaskExecutionResultType{},
		&entity.HttpResponse{},
		&entity.TemplateData{},
		// &entity.JobTemplate{},
	); err != nil {
		log.LogError("[-] error: migrating database")
		return nil, err
	}

	return &ORMConnection{Db: db}, nil
}
