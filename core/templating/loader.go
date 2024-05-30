package templating

import (
	"errors"
	"os"

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
	template, err := manager.parseTemplateFile(templatePath)
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

func (manager *TemplatingManager) LoadAllTemplates() error {
	templates, err := manager.ListAllTemplates()
	if err != nil {
		return err
	}

	for _, template := range templates {
		manager.LoadTemplate(template)
	}
	return nil
}
