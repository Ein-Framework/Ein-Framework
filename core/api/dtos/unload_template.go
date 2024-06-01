package dtos

import "github.com/Ein-Framework/Ein-Framework/core/templating"

type UnloadTemplateRequest struct {
	Name string `json:"name"`
}

type UnloadTemplateResponse struct {
	Name string `json:"name"`
	Meta templating.TemplateMeta
}
