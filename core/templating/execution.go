package templating

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"

	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"gopkg.in/yaml.v3"
)

/*
Returns whether or not a template can be executed.
*/
func (manager *TemplatingManager) CanTemplateExecute(templatePath string) error {
	template, ok := manager.loadedTemplates[templatePath]
	if !ok {
		return errors.New("template not loaded")
	}

	for _, step := range template.Steps {
		_, err := manager.pluginsManager.GetPluginByProtocol(step.Protocol)
		if err != nil {
			return fmt.Errorf("unknown protocol: '%s'", step.Protocol)
		}
	}
	return nil
}

/*
Execute the given template.

Calls CanTemplateExecute before execution.
*/
func (manager *TemplatingManager) ExecuteTemplate(templatePath string, executionContext map[string]interface{}) ([]entity.TaskExecutionResultType, error) {
	err := manager.CanTemplateExecute(templatePath)
	if err != nil {
		return nil, err
	}

	template := manager.loadedTemplates[templatePath]
	results := make([]entity.TaskExecutionResultType, 0)

	for _, step := range template.Steps {
		loadedPlugin, _ := manager.pluginsManager.GetPluginByProtocol(step.Protocol)

		// step.Data -> string -> replace vars -> map
		newStep, err := step.applyExecutionContext(executionContext)
		if err != nil {
			return nil, err
		}

		execRes := loadedPlugin.Plugin.Execute(newStep.Data)
		if execRes == nil {
			// No result
			continue
		}

		res, ok := execRes.(entity.TaskExecutionResultType)
		if !ok {
			return nil, fmt.Errorf("unable to parse template '%s' results for protocol '%s'", templatePath, step.Protocol)
		}
		// extend executionContext
		results = append(results, res)
	}
	return results, nil
}

func (step *TemplateStep) applyExecutionContext(context map[string]interface{}) (*TemplateStep, error) {
	toParse, err := yaml.Marshal(step.Data)
	if err != nil {
		return nil, fmt.Errorf("error (1) parsing protocol '%s' variables", step.Protocol)
	}

	temp, err := template.New(step.Protocol).Parse(string(toParse))
	if err != nil {
		return nil, fmt.Errorf("error (2) parsing protocol '%s' variables", step.Protocol)
	}

	var buff bytes.Buffer
	err = temp.Execute(&buff, context)
	if err != nil {
		return nil, fmt.Errorf("error (3) parsing protocol '%s' variables", step.Protocol)
	}

	var parsedData map[string]interface{}
	err = yaml.Unmarshal(buff.Bytes(), parsedData)
	if err != nil {
		return nil, fmt.Errorf("error (4) parsing protocol '%s' variables", step.Protocol)
	}

	return &TemplateStep{
		Protocol: step.Protocol,
		Data:     parsedData,
	}, nil
}
