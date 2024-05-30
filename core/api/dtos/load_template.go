package dtos

import "github.com/Ein-Framework/Ein-Framework/core/templating"

type LoadTemplateRequest struct {
	Name string `json:"name"`
}

type LoadTemplateResponse struct {
	Name string `json:"name"`
	Meta templating.TemplateMeta
}
