package templating

import (
	"os"
	"strings"
)

func (manager *TemplatingManager) GetTemplatesForJob(jobId uint) ([]TemplateData, error) {
	return nil, nil
}

func (manager *TemplatingManager) GetAllAvailableTemplates() ([]string, error) {
	files, err := os.ReadDir(manager.config.TemplatesDir)
	if os.IsNotExist(err) {
		os.Mkdir(manager.config.TemplatesDir, os.ModePerm)
		files, err = os.ReadDir(manager.config.TemplatesDir)
	}

	if err != nil {
		return nil, err
	}

	templates := make([]string, 0)
	for idx := range files {
		file := files[idx]

		if file.IsDir() || !strings.Contains(file.Name(), ".yaml") {
			continue
		}

		templates = append(templates, file.Name())
	}
	return templates, nil
}

func (manager *TemplatingManager) GetAllLoadedTemplatesOfType(typ TemplateType) map[string]*TemplateData {
	res := make(map[string]*TemplateData)
	for key, value := range manager.loadedTemplates {
		if value.Meta.Type != typ {
			continue
		}

		res[key] = value
	}

	return res
}

func (manager *TemplatingManager) GetAllLoadedTemplatesMeta() map[string]TemplateMeta {
	templates := make(map[string]TemplateMeta)
	for k, v := range manager.loadedTemplates {
		templates[k] = v.Meta
	}

	return templates
}
