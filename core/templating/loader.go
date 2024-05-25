package templating

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func (manager *TemplatingManager) parseTemplateFile(templatePath string) (interface{}, error) {
	f, err := os.ReadFile(templatePath)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("[-] error: Template file not found: " + templatePath)
		}
		return nil, errors.New("[-] error: Unknown error occured when parsing template")
	}

	var body interface{}
	if err = yaml.Unmarshal(f, &body); err != nil {
		return nil, errors.New("[-] error: Parsing template file has failed")
	}
	return body, nil
}

func (manager *TemplatingManager) ReadTemplate(templatePath string) *TemplateData {
	template, err := manager.parseTemplateFile(templatePath)
	if err != nil {
		log.Panicf(err.Error())
	}

	templateData, err := parseTemplateContent(template)
	if err != nil {
		log.Panicf(err.Error())
	}

	return templateData
}
