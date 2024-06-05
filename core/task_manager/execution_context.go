package taskmanager

import "github.com/Ein-Framework/Ein-Framework/core/domain/entity"

func CreateExecutionContext(task entity.Task, asset entity.Asset) map[string]interface{} {
	return map[string]interface{}{
		"asset":     asset.Value,
		"assetType": asset.Type,
	}
}
