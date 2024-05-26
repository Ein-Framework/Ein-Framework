package templating

import (
	"github.com/Ein-Framework/Ein-Framework/core/services"
	"github.com/Ein-Framework/Ein-Framework/pkg/plugins"
)

type TemplatingManager struct {
	coreServices   *services.Services
	pluginsManager plugins.IPluginManager
}

type ITemplateManager interface {
	ReadTemplate(templatePath string) *TemplateData
	CanTemplateExecute(templatePath string) (bool, string)
}

type TemplateData struct {
	Meta  *TemplateMeta `yaml:"meta"`
	Steps []TemplateSteps
}

type TemplateMeta struct {
	Author string `yaml:"author"`
	Type   string `yaml:"type"`
}

type TemplateSteps struct {
	Protocol string
	Data     map[string]interface{}
}
