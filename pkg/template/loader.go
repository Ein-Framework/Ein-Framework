package template

import (
	"errors"

	"github.com/Ein-Framework/Ein-Framework/pkg/log"
	"gopkg.in/yaml.v3"
)

func parseTemplateBody(templateBody []byte) (interface{}, error) {
	var body interface{}
	if err := yaml.Unmarshal(templateBody, &body); err != nil {
		return nil, errors.New("[-] error: Parsing template data has failed")
	}
	return body, nil
}

func ParseTemplate(templateBody []byte, rules []TemplateValidationRule) (interface{}, error) {
	template, err := parseTemplateBody(templateBody)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	for _, rule := range rules {
		_, err := rule(template)
		if err != nil {
			return nil, err
		}
	}

	return template, nil
}
