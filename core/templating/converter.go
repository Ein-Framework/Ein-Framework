package templating

import "fmt"

func stringToTemplateType(value string) TemplateType {
	mapper := map[string]TemplateType{
		"recon":           ReconTemplateType,
		"exploiter":       ExploiterTemplateType,
		"change_detector": ChangeDetectorTemplateType,
	}
	t, ok := mapper[value]
	if !ok {
		return UnknownTemplateType
	}
	return t
}

func parseTemplateMeta(template map[string]interface{}) (TemplateMeta, error) {
	meta := template["meta"].(map[string]interface{})
	return TemplateMeta{
		Author:      meta["author"].(string),
		Remediation: meta["remediation"].(string),
		Description: meta["description"].(string),
		Severity:    meta["severity"].(string),
		Type:        stringToTemplateType(meta["type"].(string)),
	}, nil
}

func parseTemplateSteps(template map[string]interface{}) ([]*TemplateStep, error) {
	steps := template["steps"].(map[string]interface{})

	stepsArr := make([]*TemplateStep, len(steps))
	i := 0
	for k := range steps {
		data, ok := steps[k].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("unable to convert %s to map", k)
		}

		stepsArr[i] = &TemplateStep{
			Protocol: k,
			Data:     data,
		}
		i++
	}
	return stepsArr, nil
}

func convertTemplate(template interface{}) (*TemplateData, error) {
	mapData := template.(map[string]interface{})

	meta, err := parseTemplateMeta(mapData)
	if err != nil {
		return nil, err
	}

	steps, err := parseTemplateSteps(mapData)
	if err != nil {
		return nil, err
	}

	return &TemplateData{
		Meta:  meta,
		Steps: steps,
	}, nil
}
