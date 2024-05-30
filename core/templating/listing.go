package templating

import (
	"os"
	"strings"
)

func (manager *TemplatingManager) FindTemplatesForJob(jobId uint) []TemplateData {
	return nil
}

func (manager *TemplatingManager) ListAllTemplates() ([]string, error) {
	var (
		templates []string
	)

	files, err := os.ReadDir(manager.config.TemplatesDir)
	if os.IsNotExist(err) {
		os.Mkdir(manager.config.TemplatesDir, os.ModePerm)
		files, err = os.ReadDir(manager.config.TemplatesDir)
	}

	if err != nil {
		return nil, err
	}

	for idx := range files {
		file := files[idx]

		if file.IsDir() || !strings.Contains(file.Name(), ".yaml") {
			continue
		}

		templates = append(templates, file.Name())
	}
	return templates, nil
}

func (manager *TemplatingManager) GetAllTemplatesOfType(typ TemplateType) map[string]*TemplateData {
	res := make(map[string]*TemplateData)
	for key, value := range manager.loadedTemplates {
		if value.Meta.Type != typ {
			continue
		}

		res[key] = value
	}

	return res
}
