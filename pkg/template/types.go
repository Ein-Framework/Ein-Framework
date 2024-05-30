package template

type TemplateValidationRule func(interface{}) (interface{}, error)
