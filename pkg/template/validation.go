package template

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

func KeyExists(key string) TemplateValidationRule {
	return func(data interface{}) (interface{}, error) {
		casted, ok := data.(map[string]interface{})
		if !ok {
			return nil, errors.New("unable to cast interface to map")
		}

		res, ok := casted[key]
		if !ok {
			return nil, fmt.Errorf("key '%s' does not exist", key)
		}
		return res, nil
	}
}

func KeyExistsValueString(key string) TemplateValidationRule {
	existsValidator := KeyExists(key)
	return func(data interface{}) (interface{}, error) {
		obj, err := existsValidator(data)
		if err != nil {
			return nil, err
		}
		str, ok := obj.(string)
		if !ok {
			return nil, fmt.Errorf("key '%s' exists but is not string", key)
		}

		return str, nil
	}
}

func KeyExistsValueEnumString(key string, values []string) TemplateValidationRule {
	existsValidator := KeyExists(key)
	return func(data interface{}) (interface{}, error) {
		obj, err := existsValidator(data)
		if err != nil {
			return nil, err
		}

		str, ok := obj.(string)
		if !ok {
			return nil, fmt.Errorf("key '%s' exists but is not string", key)
		}

		ok = slices.Contains(values, str)
		if !ok {
			return nil, fmt.Errorf("key '%s' exists but is not allowed. allowed values are: %s", key, strings.Join(values, ", "))
		}
		return str, nil
	}
}

func ValidateResultOf(rootValidation TemplateValidationRule, validations ...TemplateValidationRule) TemplateValidationRule {
	return func(i interface{}) (interface{}, error) {
		res, err := rootValidation(i)
		if err != nil {
			return nil, err
		}

		if len(validations) == 0 {
			return res, nil
		}

		for _, validation := range validations {
			_, err = validation(res)
			if err != nil {
				return nil, err
			}
		}

		return res, nil
	}
}

func ChainValidations(validations ...TemplateValidationRule) TemplateValidationRule {
	return func(i interface{}) (interface{}, error) {
		if len(validations) == 0 {
			return i, nil
		}
		if len(validations) == 1 {
			return validations[0](i)
		}

		res, err := validations[0](i)
		if err != nil {
			return nil, err
		}

		for idx := range len(validations) - 1 {
			res, err = validations[idx+1](res)
			if err != nil {
				return nil, err
			}
		}

		return res, nil
	}
}
