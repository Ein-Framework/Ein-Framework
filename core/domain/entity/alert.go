package entity

import "gorm.io/gorm"

type Alert struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Scope       string `json:"scope"`
	// TaskExecutionResultId uint                    `json:"-"`
	// TaskExecutionResult   TaskExecutionResultType `json:"taskExecutionResult" gorm:"foreignkey:TaskExecutionResultId;association_foreignkey:ID;"`
}
