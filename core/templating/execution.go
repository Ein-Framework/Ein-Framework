package templating

import (
	"fmt"
)

/*
Returns whether or not a template can be executed.

If false, it'll return the protocol causing it to fail.
*/
func (manager *TemplatingManager) CanTemplateExecute(templatePath string) (bool, string) {
	template := manager.ReadTemplate(templatePath)

	for _, step := range template.Steps {
		_, err := manager.pluginsManager.GetPluginByProtocol(step.Protocol)
		if err != nil {
			return false, step.Protocol
		}
	}
	return true, ""
}

/*
Execute the given template.

TODO(M0ngi): Handle results
*/
func (manager *TemplatingManager) ExecuteTemplate(templatePath string, params ...interface{}) (interface{}, error) {
	template := manager.ReadTemplate(templatePath)

	for _, step := range template.Steps {
		loadedPlugin, err := manager.pluginsManager.GetPluginByProtocol(step.Protocol)
		if err != nil {
			return nil, fmt.Errorf("unable to execute template '%s', missing plugin for protocol '%s'", templatePath, step.Protocol)
		}

		// TODO(M0ngi): Handle plugin execution results
		loadedPlugin.Plugin.Execute(params)
	}
	return nil, nil
}