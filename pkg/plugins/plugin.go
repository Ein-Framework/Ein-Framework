// Package plugins defines how to load and interact with the framework plugins.
package plugins

import (
	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
)

type PluginType int

// To add a new type, make sure it's between `UnknownType` & `UndefinedType` (Used for validation)
// `UnknownType` must always be 0, which is a default value for int type.
const (
	UnknownType    PluginType = iota
	Recon                     // Enumeration phase
	Exploiter                 // Scan for a vulnerability/exploit
	ChangeDetector            // Watch for a change detection
	UndefinedType
)

// Metadata Type describes the metadata of the plugin
// Version is the current version of the plugin
type Metadata struct {
	Version     string     `json:"version"`
	Author      string     `json:"author"`
	Tags        []string   `json:"tags"`
	ReleaseDate string     `json:"releaseDate"`
	Type        PluginType `json:"type"`
	SourceLink  string     `json:"sourceLink"`
	Description string     `json:"description"`
}

// PluginInfo ReturnType returns the type of returned data, so we can parse it
// Options is a map where the key is the args name and string is the plugin's type which can be either bytes/rune/int/bool/string
type PluginInfo struct {
	Name       string            `json:"name"`
	Options    map[string]string `json:"options"`
	Protocol   string            `json:"protocol"`
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
