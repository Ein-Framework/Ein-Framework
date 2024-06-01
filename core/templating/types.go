package templating

import (
	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/core/services"
	"github.com/Ein-Framework/Ein-Framework/pkg/config"
	"github.com/Ein-Framework/Ein-Framework/pkg/plugins"
)

type TemplatingManager struct {
	coreServices    *services.Services
	pluginsManager  plugins.IPluginManager
	config          *config.Config
	loadedTemplates map[string]*TemplateData
}

type ITemplateManager interface {
	ReadTemplate(templatePath string) (*TemplateData, error)
	CanTemplateExecute(templatePath string) error
	ExecuteTemplate(templatePath string, executionContext map[string]interface{}) ([]TemplateExecutionResultType, error)

	GetAllLoadedTemplatesOfType(typ TemplateType) map[string]*TemplateData
	GetAllLoadedTemplatesMeta() map[string]TemplateMeta

	GetTemplatesForJob(jobId uint) ([]TemplateData, error)
	GetAllAvailableTemplates() ([]string, error)

	LoadTemplate(templateFile string) (*TemplateData, error)
	UnloadTemplate(templateFile string) error
	LoadAllTemplates() error
	UnloadAllTemplates() error

	PluginManager() plugins.IPluginManager
}

type TemplateData struct {
	Meta  TemplateMeta `yaml:"meta"`
	Steps []*TemplateStep
}

type TemplateType int

// To add a new type, make sure it's between `UnknownType` & `UndefinedType` (Used for validation)
// `UnknownType` must always be 0, which is a default value for int type.
const (
	UnknownTemplateType        TemplateType = iota
	ReconTemplateType                       // Enumeration phase
	ExploiterTemplateType                   // Scan for a vulnerability/exploit
	ChangeDetectorTemplateType              // Watch for a change detection
	UndefinedTemplateType
)

type TemplateMeta struct {
	Author string       `yaml:"author"`
	Type   TemplateType `yaml:"type"`
}

type TemplateStep struct {
	Protocol string
	Data     map[string]interface{}
}

type TemplateExecutionResultType struct {
	Response  string
	MetaData  map[string]string
	NewAssets []entity.Asset
	Alerts    []entity.Alert
}
