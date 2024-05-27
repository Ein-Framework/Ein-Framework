package templating

import (
	"errors"
	"fmt"
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
func (manager *TemplatingManager) ExecuteTemplate(templatePath string, executionContext ...interface{}) ([]TemplateExecutionResultType, error) {
	err := manager.CanTemplateExecute(templatePath)
	if err != nil {
		return nil, err
	}

	template := manager.loadedTemplates[templatePath]
	results := make([]TemplateExecutionResultType, len(template.Steps))

	for idx, step := range template.Steps {
		loadedPlugin, _ := manager.pluginsManager.GetPluginByProtocol(step.Protocol)

		execRes := loadedPlugin.Plugin.Execute(executionContext)
		res, ok := execRes.(TemplateExecutionResultType)
		if !ok {
			return nil, fmt.Errorf("unable to parse template '%s' results for protocol '%s'", templatePath, step.Protocol)
		}
		results[idx] = res
	}
	return results, nil
}
