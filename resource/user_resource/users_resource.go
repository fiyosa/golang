package user_resource

import (
	"fmt"
	"tutorial/database/migration"
	"tutorial/helper"
)

func Users(datas []migration.User) []map[string]interface{} {
	if len(datas) == 0 {
		return []map[string]interface{}{}
	}

	var newDatas []map[string]interface{}

	for _, data := range datas {
		var roles []interface{}
		var permissions []string

		for _, value := range data.UserHasRoles {
			roles = append(roles, value.Role.Name)
		}

		for _, value := range data.UserHasRoles {
			for _, permission := range value.Role.RoleHasPermissions {
				permissions = append(permissions, permission.Permission.Name)
			}
		}

		newDatas = append(newDatas, helper.H{
			"id":         helper.EncodeId(data.Id),
			"email":      data.Email,
			"name":       data.Name,
			"role":       roles,
			"permission": helper.RemoveDuplicateStr(permissions),
			"created_at": helper.TimeFormat(fmt.Sprintf("%v", data.CreatedAt), helper.TimeFormatInput, helper.TimeFormatSQL),
			"updated_at": helper.TimeFormat(fmt.Sprintf("%v", data.UpdatedAt), helper.TimeFormatInput, helper.TimeFormatSQL),
		})
	}
	return newDatas
}
