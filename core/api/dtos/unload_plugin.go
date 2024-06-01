package dtos

import "github.com/Ein-Framework/Ein-Framework/core/templating"

type UnloadPluginRequest struct {
	Name string `json:"name"`
}

type UnloadPluginResponse struct {
	Name string `json:"name"`
	Meta templating.TemplateMeta
}
