// Package plugins defines how to load and interact with the framework plugins.
package plugins

import (
	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
)

// Scanner: Scan for a vulnerability/exploit
// ChangeDetector: Watch for a change detection
const (
	Scanner        int = 0
	ChangeDetector int = 1
)

// Metadata Type describes the metadata of the plugin
// Version is the current version of the plugin
type Metadata struct {
	Version     string   `json:"version"`
	Author      string   `json:"author"`
	Tags        []string `json:"tags"`
	ReleaseDate string   `json:"releaseDate"`
	Type        int      `json:"type"`
	Protocol    string   `json:"protocol"`
	SourceLink  string   `json:"sourceLink"`
	Description string   `json:"description"`
}

// PluginInfo ReturnType returns the type of returned data, so we can parse it
// Options is a map where the key is the args name and string is the plugin's type which can be either bytes/rune/int/bool/string
type PluginInfo struct {
	Name       string            `json:"name"`
	Options    map[string]string `json:"options"`
	ReturnType string            `json:"returnType"`
}

// Plugin have a Name which is defined by the author.
type Plugin struct {
	Metadata   `json:"metdata"`
	PluginInfo `json:"pluginInfo"`
}

// IPlugin is the interface that all plugins should implement.
// The CLI will generate a scaffold, and it will make sure that it add this interface on the top
type IPlugin interface {
	MetaInfo() *Metadata
	Info() *PluginInfo
	Options() map[string]string
	Execute(...interface{}) []entity.AlertModel
	SetArgs(map[string]interface{}) error
	IsWaitingForTaskResult() (bool, string)
}