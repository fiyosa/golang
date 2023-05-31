package user_repository

import (
	"tutorial/config/db"
	"tutorial/controller"
	"tutorial/database/migration"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context, users *[]migration.User) bool {
	page, limit, keyword := controller.Param(c)

	res := db.DB.
		Model(&migration.User{}).
		Preload("UserHasRoles.Role").
		Preload("UserHasRoles.Role.RoleHasPermissions.Permission").
		Offset(page).
		Limit(limit).
		Where("email like ?", "%"+keyword+"%").
		Where("name like ?", "%"+keyword+"%").
		Find(&users)

	if !controller.ResultQuery(c, res) {
		return false
	}

	return true
}
