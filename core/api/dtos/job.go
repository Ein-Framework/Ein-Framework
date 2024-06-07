package dtos

import "github.com/Ein-Framework/Ein-Framework/core/domain/entity"

type CreateJobRequest struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Templates   []entity.Template `json:"templates"`
}

type UpdateJobRequest struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Templates   []entity.Template `json:"templates"`
}
