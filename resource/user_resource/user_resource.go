package user_resource

import (
	"fmt"
	"tutorial/database/migration"
	"tutorial/helper"
)

func User(data migration.User) map[string]interface{} {
	if data.Email == "" {
		return nil
	}

	var newData map[string]interface{}
	var roles []string
	var permissions []string

	for _, value := range data.UserHasRoles {
		roles = append(roles, value.Role.Name)
	}

	for _, value := range data.UserHasRoles {
		for _, permission := range value.Role.RoleHasPermissions {
			permissions = append(permissions, permission.Permission.Name)
		}
	}

	newData = map[string]interface{}{
		"id":         helper.EncodeId(data.Id),
		"email":      data.Email,
		"name":       data.Name,
		"role":       roles,
		"permission": helper.RemoveDuplicateStr(permissions),
		"created_at": helper.TimeFormat(fmt.Sprintf("%v", data.CreatedAt), helper.TimeFormatInput, helper.TimeFormatSQL),
		"updated_at": helper.TimeFormat(fmt.Sprintf("%v", data.UpdatedAt), helper.TimeFormatInput, helper.TimeFormatSQL),
	}
	return newData
}
