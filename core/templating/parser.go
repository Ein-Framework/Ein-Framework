package templating

import (
	"errors"
	"fmt"
)

func getStringField(dict map[string]interface{}, field string) (string, error) {
	val, ok := dict[field]
	if !ok {
		return "", errors.New(fmt.Sprintf("`%s` is missing in template content", field))
	}

	valStr, ok := val.(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("`%s` is missing in template content", field))
	}

	return valStr, nil
}

func parseTemplateMeta(meta map[string]interface{}) (*TemplateMeta, error) {
	author, err := getStringField(meta, "author")
	if err != nil {
		return nil, err
	}

	return &TemplateMeta{
		Author: author,
	}, nil
}

func parseTemplateContent(data interface{}) (*TemplateData, error) {
	json, ok := data.(map[string]interface{})
	if !ok {
		return nil, errors.New("unable to convert template content to map")
	}

	meta, ok := json["meta"]
	if !ok {
		return nil, errors.New("`meta` is missing in template content")
	}

	metaMap, ok := meta.(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid `meta`")
	}

	metaParsed, err := parseTemplateMeta(metaMap)
	if err != nil {
		return nil, err
	}

	return &TemplateData{
		Meta: metaParsed,
	}, nil
}
