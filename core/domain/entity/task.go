package entity

import "time"

type Template interface{}

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
	Type JobType
	Name string
}

type PeriodConfiguration struct {
	Tasks []Task
}

type CronJob struct {
	Job
	PeriodConfiguration PeriodConfiguration
}

type Task struct {
	ID              uint `gorm:"primaryKey"`
	Template        Template
	CreatedAt       time.Time
	State           TaskState
	Output          string
	OutputFormat    OutputFormat
	Args            map[string]string
	AssessmentStage AssessmentStage
}
