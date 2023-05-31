package role_repository

import (
	"tutorial/config/db"
	"tutorial/controller"
	"tutorial/database/migration"
	"tutorial/helper"
	"tutorial/validation/role_validation"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context, input role_validation.Update) bool {
	var roleHasPermission []*migration.RoleHasPermission
	var permissions []*migration.Permission
	var role *migration.Role

	id := c.Param("id")

	dateNow := helper.TimeNow()

	tx := db.DB.Begin()

	resPermission := tx.
		Model(&migration.Permission{}).
		Find(&permissions)

	if !controller.ResultQuery(c, resPermission) {
		tx.Rollback()
		return false
	}

	resRole := tx.
		Model(&migration.Role{}).
		First(&role, "id = ?", helper.DecodeId(id))

	if !controller.ResultQuery(c, resRole) {
		tx.Rollback()
		return false
	}

	role.Name = input.Name

	if resRoleUpdate := tx.Save(&role); !controller.ResultQuery(c, resRoleUpdate) {
		tx.Rollback()
		return false
	}

	resRoleHasPermission := tx.Delete(&migration.RoleHasPermission{}, "role_id = ?", role.Id)

	if !controller.ResultQuery(c, resRoleHasPermission) {
		tx.Rollback()
		return false
	}

	newInputPermissions := helper.RemoveDuplicateStr(input.Permission)

	for _, value := range newInputPermissions {
		for _, permissionId := range permissions {
			if permissionId.Name == value {
				roleHasPermission = append(roleHasPermission, &migration.RoleHasPermission{
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
