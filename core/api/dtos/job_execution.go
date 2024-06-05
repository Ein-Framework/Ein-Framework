package dtos

import "github.com/Ein-Framework/Ein-Framework/core/domain/entity"

type RunJobRequest struct {
	JobId        uint `json:"jobId"`
	AssessmentId uint `json:"assessmentId"`
}

type RunTemplateRequest struct {
	Template     entity.Template `json:"template"`
	AssessmentId uint            `json:"assessmentId"`
}
