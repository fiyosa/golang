package seeder

import (
	"log"
	"tutorial/config/db"
	"tutorial/database/migration"
	"tutorial/helper"
)

var ListRoles = []string{
	"Super Admin",
	"Admin",
	"User",
}

func RolePermissionSeeder() bool {
	listPermissions := []string{
		"user_index",
		"user_show",
		"user_update",
		"user_delete",
		"role_index",
		"role_store",
		"role_update",
		"role_show",
		"role_delete",
	}

	listRolePermission := map[string][]string{
		"Super Admin": {
			"user_index",
			"user_show",
			"user_update",
			"user_delete",
			"role_index",
			"role_store",
			"role_update",
			"role_show",
			"role_delete",
		},
		"Admin": {
			"user_index",
			"user_show",
			"user_update",
			"user_delete",
		},
		"User": {
			"user_update",
		},
	}

	dateNow := helper.TimeNow()
	var newListRoles []migration.Role
	var newListPermissions []migration.Permission
	var newListRoleHasPermissions []migration.RoleHasPermission

	for _, value := range ListRoles {
		newListRoles = append(newListRoles, migration.Role{
			Name:      value,
			CreatedAt: dateNow,
			UpdatedAt: dateNow,
		})
	}

	for _, value := range listPermissions {
		newListPermissions = append(newListPermissions, migration.Permission{
			Name:      value,
			CreatedAt: dateNow,
			UpdatedAt: dateNow,
		})
	}

	for role, permissions := range listRolePermission {
		for _, permission := range permissions {
			newListRoleHasPermissions = append(newListRoleHasPermissions, migration.RoleHasPermission{
				RoleId:       helper.SliceIndexStr(ListRoles, role) + 1,
				PermissionId: helper.SliceIndexStr(listPermissions, permission) + 1,
				CreatedAt:    dateNow,
				UpdatedAt:    dateNow,
			})
		}
	}

	tx := db.DB.Begin()

	if err := tx.Create(newListRoles); err.Error != nil {
		log.Printf("Seeder failed to add table role: %v", err.Error.Error())
		tx.Rollback()
		return false
	}
	if err := tx.Create(newListPermissions); err.Error != nil {
		log.Print("Seeder failed to add table permission")
		tx.Rollback()
		return false
	}
	if err := tx.Create(newListRoleHasPermissions); err.Error != nil {
		log.Print("Seeder failed to add table role has permission")
		tx.Rollback()
		return false
	}

	tx.Commit()
	log.Print("Seeder table role permission successfully")
	return true
}
