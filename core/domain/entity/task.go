package entity

import "time"

type OutputFormat string

const (
	CsvOutput  = "csv"
	JsonOutput = "json"
	HTMLOutput = "html"
	TextOutout = "txt"
)

// TODO: Change to concrete type when Template module is ready

type TaskState string

type Args map[string]string

const (
	Running  TaskState = "Running"
	Stopped  TaskState = "Stopped"
	Suspened TaskState = "Suspeneded"
	Canceled TaskState = "Canceled"
	Queued   TaskState = "Queued"
)

type Template interface{}

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
	Args            Args
	AssessmentStage AssessmentStage
}
