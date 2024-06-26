package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

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

// type JobType string

// const (
// 	Once     JobType = "once"
// 	Periodic JobType = "periodic"
// )

type Job struct {
	gorm.Model
	// Type JobType `gorm:"type:text"`
	Name          string         `json:"name" gorm:"type:text"`
	Description   string         `json:"description" gorm:"type:text"`
	Templates     []TemplateData `json:"templates" gorm:"many2many:job_templates;"`
	JobExecutions []JobExecution `json:"jobExecutions" gorm:"foreignKey:JobID"`
}

type TemplateData struct {
	Template  Template `json:"template" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type JobExecution struct {
	gorm.Model
	Job          Job
	JobID        uint       `json:"-"`
	Assessment   Assessment `json:"assessment" gorm:"foreignkey:AssessmentId;association_foreignkey:ID;"`
	AssessmentId uint       `json:"-"`
	Tasks        []Task     `json:"tasks" gorm:"foreignKey:JobExecutionId"` // gorm:"many2many:period_configuration_tasks;"
	Status       TaskState  `json:"status"`
}

type PeriodConfiguration struct {
	gorm.Model
	Month uint
	Day   uint
	Hour  uint
	// Tasks []Task `json:"tasks" gorm:"many2many:period_configuration_tasks;"`
}

type CronJob struct {
	Job
	PeriodConfiguration   PeriodConfiguration `json:"periodConfiguration" gorm:"foreignkey:PeriodConfigurationID;association_foreignkey:ID;"`
	PeriodConfigurationID uint                `json:"-"`
}

type Task struct {
	gorm.Model
	Template          Template
	State             TaskState `gorm:"type:text"`
	Output            string
	OutputFormat      OutputFormat      `gorm:"type:text"`
	Args              map[string]string `gorm:"type:text"`
	AssessmentId      uint              `json:"-"`
	Assessment        Assessment        // `json:"assessment" gorm:"foreignkey:AssessmentId;association_foreignkey:ID;"`
	AssessmentStageId uint              `json:"-"`
	AssessmentStage   AssessmentStage   // `json:"assessmentStage" gorm:"foreignkey:AssessmentStageId;association_foreignkey:ID;"`
	JobExecutionId    *uint             `json:"-"`
	JobExecution      JobExecution      // `json:"assessmentStage" gorm:"foreignkey:AssessmentStageId;association_foreignkey:ID;"`
}

type HttpResponse struct {
	gorm.Model
	ResponseBody    string
	ResponseHeaders JSONMap
}

type TaskExecutionResultType struct {
	gorm.Model
	ResponseId       uint         `json:"-"`
	Response         HttpResponse `json:"httpResponse" gorm:"foreignkey:ResponseId;association_foreignkey:ID;"`
	MetaData         JSONMap
	NewAssets        []Asset `json:"tasks" gorm:"many2many:task_execution_assets;"`
	Alerts           []Alert `gorm:"foreignKey:TaskExecutionResultId"`
	TaskOutputFormat OutputFormat
	TaskOutput       string // For display
	TaskId           uint   `json:"-"`
	Task             Task   `json:"task" gorm:"foreignkey:TaskId;association_foreignkey:ID;"`
}

type JSONMap map[string]string

func (j *JSONMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}

	return json.Unmarshal(bytes, j)
}

// Value implements the driver Valuer interface for JSONMap
func (j JSONMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}
