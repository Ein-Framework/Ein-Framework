package entity

import (
	"github.com/Ein-Framework/Ein-Framework/pkg/repository"
	"gorm.io/gorm"
)

type AssessmentType string

const (
	VDP AssessmentType = "vdp"
	BB  AssessmentType = "bb"
)

type Attachement struct {
	gorm.Model
	Type string `json:"type"`
	Link string `json:"link"`
}

type Report struct {
	gorm.Model
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Attachements []Attachement `json:"attachements" gorm:"many2many:report_attachements;"`
	Severity     uint          `json:"severity"`
}

type Assessment struct {
	gorm.Model
	Name              string          `json:"name"`
	Type              AssessmentType  `json:"type" gorm:"type:text"`
	Scope             Scope           `json:"scope" gorm:"foreignkey:ScopeID;association_foreignkey:ID;"`
	ScopeID           uint            `json:"-"`
	Assets            []Asset         `json:"assets" gorm:"many2many:assessment_assets;"`
	Stage             AssessmentStage `json:"assessmentStage" gorm:"foreignkey:StageID;association_foreignkey:ID;"`
	StageID           uint            `json:"-"`
	EngagementRules   EngagementRules `json:"engagementRules" gorm:"foreignkey:EngagementRulesID;association_foreignkey:ID;"`
	EngagementRulesID uint            `json:"-"`
	// Jobs              []Job           `json:"jobs" gorm:"many2many:assessment_jobs;"`
	Reports []Report `json:"reports" gorm:"many2many:assessment_reports;"`
	Tasks   []Task   `json:"tasks" gorm:"foreignKey:AssessmentId"`
}

func NewAssessment(name string, assessmentType AssessmentType, scope Scope, engagementRules EngagementRules, repo repository.Repository) (*Assessment, error) {
	reconStage, err := GetStageByName(ReconnaissanceStage, repo)
	if err != nil {
		return nil, err
	}

	return &Assessment{
		Name:            name,
		Type:            assessmentType,
		Scope:           scope,
		Stage:           *reconStage,
		Assets:          []Asset{},
		EngagementRules: engagementRules,
		Reports:         []Report{},
		// Jobs:    []Job{},
	}, nil
}

func (a *Assessment) ToggleStageCompletion(repo repository.Repository) error {
	a.Stage.Completed = !a.Stage.Completed
	return repo.Save(&a.Stage)
}

func (a *Assessment) SetCurrentStage(stageName string, repo repository.Repository) error {
	stage, err := GetStageByName(stageName, repo)
	if err != nil {
		return err
	}
	a.Stage = *stage
	return nil
}

func (a *Assessment) CheckStagePlugins(repo repository.Repository) error {
	// Implement the logic to check stage plugins
	// For example, you can check if the required plugins are installed or up-to-date
	// based on the current stage's keywords or other properties.
	return nil
}

func (a *Assessment) RunTasks(repo repository.Repository) error {
	// Implement the logic to run tasks for the current stage
	// You can iterate through a.Stage.Tasks and execute each task using a task runner.
	return nil
}

// func (a *Assessment) ListStageTasks() []Task {
// 	return a.Stage.Tasks
// }

func (a *Assessment) ViewStageTasksQueue() []Task {
	// Implement the logic to view the tasks queue for the current stage
	// You can filter the tasks based on their status or other properties.
	return []Task{}
}
