package role_repository

import (
	"tutorial/config/db"
	"tutorial/controller"
	"tutorial/database/migration"
	"tutorial/helper"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context, role *migration.Role) bool {
	id := c.Param("id")

	resRole := db.DB.
		Model(&migration.Role{}).
		Preload("RoleHasPermissions.Permission").
		First(&role, "id = ?", helper.DecodeId(id))

	if !controller.ResultQuery(c, resRole) {
		return false
	}

	return true
}
