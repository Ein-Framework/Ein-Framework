package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string       `json:"username"`
	Password    string       `json:"password"`
	Email       string       `json:"email"`
	Assessments []Assessment `json:"assessments" gorm:"many2many:user_assessments;"`
	// CustomTemplates []Template   `json:"customTemplates"`
}
