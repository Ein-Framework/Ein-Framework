package templating

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Ein-Framework/Ein-Framework/pkg/log"
	"github.com/Ein-Framework/Ein-Framework/pkg/template"
)

func (manager *TemplatingManager) parseTemplateFile(templatePath string) (interface{}, error) {
	f, err := os.ReadFile(templatePath)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("[-] error: Template file not found: " + templatePath)
		}
		return nil, errors.New("[-] error: Unknown error occured when parsing template")
	}

	template, err := template.ParseTemplate(
		f,
		templateValidationRules(),
	)
	if err != nil {
		return nil, err
	}
	return template, nil
}

func (manager *TemplatingManager) ReadTemplate(templatePath string) (*TemplateData, error) {
	template, err := manager.parseTemplateFile(filepath.Join(manager.config.TemplatesDir, templatePath))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	templateData, err := convertTemplate(template)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	err = validateTemplate(templateData)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return templateData, nil
}

func (manager *TemplatingManager) LoadTemplate(templateFile string) error {
	template, err := manager.ReadTemplate(templateFile)
	if err != nil {
		return err
	}

	manager.loadedTemplates[templateFile] = template
	return nil
}

func (manager *TemplatingManager) UnloadTemplate(templateFile string) error {
	_, ok := manager.loadedTemplates[templateFile]
	if !ok {
		return errors.New("template must be loaded to unload it")
	}
	delete(manager.loadedTemplates, templateFile)
	return nil
}

func (manager *TemplatingManager) LoadAllTemplates() error {
	templates, err := manager.ListAllAvailableTemplates()
	if err != nil {
		return err
	}

	for _, template := range templates {
		err := manager.LoadTemplate(template)
		if err != nil {
			return err
		}
	}
	return nil
}

func (manager *TemplatingManager) UnloadAllTemplates() error {
	for k := range manager.loadedTemplates {
		err := manager.UnloadTemplate(k)
		if err != nil {
			return err
		}
	}
	return nil
}
