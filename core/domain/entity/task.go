package entity

import (
	"gorm.io/gorm"
)

type Template string

type OutputFormat string

const (
	CsvOutput  OutputFormat = "csv"
	JsonOutput OutputFormat = "json"
	HTMLOutput OutputFormat = "html"
	TextOutout OutputFormat = "txt"
)

type TaskState string

const (
	Running  TaskState = "Running"
	Stopped  TaskState = "Stopped"
	Suspened TaskState = "Suspeneded"
	Canceled TaskState = "Canceled"
	Queued   TaskState = "Queued"
)

type JobType string

const (
	Once     JobType = "once"
	Periodic JobType = "periodic"
)

type Job struct {
	gorm.Model
	Type JobType `gorm:"type:text"`
	Name string
}

type PeriodConfiguration struct {
	gorm.Model
	Tasks []Task `json:"tasks" gorm:"many2many:period_configuration_tasks;"`
}

type CronJob struct {
	Job
	PeriodConfiguration   PeriodConfiguration `json:"periodConfiguration" gorm:"foreignkey:PeriodConfigurationID;association_foreignkey:ID;"`
	PeriodConfigurationID uint                `json:"-"`
}

type Task struct {
	gorm.Model
	// Template        Template
	State             TaskState `gorm:"type:text"`
	Output            string
	OutputFormat      OutputFormat      `gorm:"type:text"`
	Args              map[string]string `gorm:"type:text"`
	AssessmentStageId uint              `json:"-"`
	AssessmentStage   AssessmentStage   `json:"assessmentStage" gorm:"foreignkey:AssessmentStageId;association_foreignkey:ID;"`
}
