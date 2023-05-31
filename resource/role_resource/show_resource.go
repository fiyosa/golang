package role_resource

import (
	"fmt"
	"tutorial/database/migration"
	"tutorial/helper"
)

func Show(role migration.Role) map[string]interface{} {
	if role.Id == 0 {
		return nil
	}

	var newData map[string]interface{}

	var permission []string

	for _, permissions := range role.RoleHasPermissions {
		permission = append(permission, permissions.Permission.Name)
	}

	newData = map[string]interface{}{
		"id":         helper.EncodeId(role.Id),
		"name":       role.Name,
		"permission": helper.RemoveDuplicateStr(permission),
		"created_at": helper.TimeFormat(fmt.Sprintf("%v", role.CreatedAt), helper.TimeFormatInput, helper.TimeFormatSQL),
		"updated_at": helper.TimeFormat(fmt.Sprintf("%v", role.UpdatedAt), helper.TimeFormatInput, helper.TimeFormatSQL),
	}

	return newData
}
