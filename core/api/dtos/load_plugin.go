package dtos

import (
	"github.com/Ein-Framework/Ein-Framework/pkg/plugins"
)

type LoadPluginRequest struct {
	Name string `json:"name"`
}

type LoadPluginResponse struct {
	Name string `json:"name"`
	Meta *plugins.Metadata
	Info *plugins.PluginInfo
}
