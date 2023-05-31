package role_repository

import (
	"tutorial/config/db"
	"tutorial/controller"
	"tutorial/database/migration"
	"tutorial/helper"
	"tutorial/validation/role_validation"

	"github.com/gin-gonic/gin"
)

func Store(c *gin.Context, input role_validation.Store) bool {
	var permissions []*migration.Permission
	var roleHasPermission []migration.RoleHasPermission

	resPermission := db.DB.
		Model(&migration.Permission{}).
		Find(&permissions)

	if !controller.ResultQuery(c, resPermission) {
		return false
	}

	dateNow := helper.TimeNow()

	role := &migration.Role{
		Name:      input.Name,
		CreatedAt: dateNow,
		UpdatedAt: dateNow,
	}

	tx := db.DB.Begin()

	if err := tx.Create(&role); err.Error != nil {
		helper.SendCrash(c, err.Error.Error())
		tx.Rollback()
		return false
	}

	newInputPermissions := helper.RemoveDuplicateStr(input.Permission)

	for _, value := range newInputPermissions {
		for _, permissionId := range permissions {
			if permissionId.Name == value {
				roleHasPermission = append(roleHasPermission, migration.RoleHasPermission{
					RoleId:       role.Id,
					PermissionId: permissionId.Id,
					CreatedAt:    dateNow,
					UpdatedAt:    dateNow,
				})
			}
		}
	}

	if err := tx.Create(&roleHasPermission); err.Error != nil {
		helper.SendCrash(c, err.Error.Error())
		tx.Rollback()
		return false
	}

	tx.Commit()

	return true
}
