package templating

import (
	"errors"

	"github.com/Ein-Framework/Ein-Framework/pkg/template"
)

func templateValidationRules() []template.TemplateValidationRule {
	return []template.TemplateValidationRule{
		template.ValidateResultOf(
			template.KeyExists("meta"),
			template.KeyExistsValueString("author"),
			template.KeyExistsValueString("type"),
		),
		template.KeyExists("steps"),
	}
}

func validateTemplate(template *TemplateData) error {
	if template.Meta.Type >= UndefinedTemplateType || template.Meta.Type <= UnknownTemplateType {
		return errors.New("unknown template type")
	}

	if len(template.Steps) == 0 {
		return errors.New("template requires at least 1 step")
	}
	return nil
}
