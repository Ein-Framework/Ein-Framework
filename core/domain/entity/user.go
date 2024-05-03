package entity

type User struct {
	Username        string
	Password        string
	Email           string
	Assessments     []Assessment
	CustomTemplates []Template
}
