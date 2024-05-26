package templating

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
