package entity

import (
	"github.com/Ein-Framework/Ein-Framework/pkg/repository"
	"gorm.io/gorm"
)

type AssessmentStage struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Completed   bool     `json:"completed"`
	Keywords    []string `json:"keywords" gorm:"type:text"`
	Link        string   `json:"link"`
	Tasks       []Task   `json:"tasks" gorm:"many2many:assessment_stage_tasks;"`
}

const (
	ReconnaissanceStage = "reconnaissance"
	ScanningStage       = "scanning"
	AutomatedChecking   = "automated_checking"
	MappingStage        = "mapping"
	ExploitationStage   = "exploitation"
	ReportingStage      = "reporting"
)

var DefaultStages = []string{
	ReconnaissanceStage,
	ScanningStage,
	AutomatedChecking,
	MappingStage,
	ExploitationStage,
	ReportingStage,
}

func GetStageByName(name string, repo repository.Repository) (*AssessmentStage, error) {
	var stage AssessmentStage
	err := repo.GetOneByField(&stage, "name", name)
	if err != nil {
		return nil, err
	}
	return &stage, nil
}
