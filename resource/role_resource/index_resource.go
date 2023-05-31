package role_resource

import (
	"fmt"
	"tutorial/database/migration"
	"tutorial/helper"
)

func Index(datas []migration.Role) []map[string]interface{} {
	if len(datas) == 0 {
		return []map[string]interface{}{}
	}

	var newDatas []map[string]interface{}

	for _, data := range datas {
		var permission []string

		for _, permissions := range data.RoleHasPermissions {
			permission = append(permission, permissions.Permission.Name)
		}

		newDatas = append(newDatas, helper.H{
			"id":         helper.EncodeId(data.Id),
			"name":       data.Name,
			"permission": helper.RemoveDuplicateStr(permission),
			// "permission": data.RoleHasPermissions,
			"created_at": helper.TimeFormat(fmt.Sprintf("%v", data.CreatedAt), helper.TimeFormatInput, helper.TimeFormatSQL),
			"updated_at": helper.TimeFormat(fmt.Sprintf("%v", data.UpdatedAt), helper.TimeFormatInput, helper.TimeFormatSQL),
		})
	}
	return newDatas
}
