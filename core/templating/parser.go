package templating

import (
	"errors"
	"fmt"
)

func getStringField(dict map[string]interface{}, field string) (string, error) {
	val, ok := dict[field]
	if !ok {
		return "", fmt.Errorf("`%s` is missing in template content", field)
	}

	valStr, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("`%s` is missing in template content", field)
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
	mapData, ok := data.(map[string]interface{})
	if !ok {
		return nil, errors.New("unable to convert template content to map")
	}

	meta, ok := mapData["meta"]
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

	steps, ok := mapData["steps"]
	if !ok {
		return nil, errors.New("`steps` is missing in template content")
	}

	stepsMap, ok := steps.(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid `steps`, must be a map")
	}

	stepsArr := make([]TemplateSteps, len(stepsMap))
	i := 0
	for k := range stepsMap {
		data, ok := stepsMap[k].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("unable to convert %s to map", k)
		}

		stepsArr[i].Protocol = k
		stepsArr[i].Data = data
		i++
	}

	return &TemplateData{
		Meta:  metaParsed,
		Steps: stepsArr,
	}, nil
}
