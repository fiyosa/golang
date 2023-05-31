package role_repository

import (
	"log"
	"tutorial/config/db"
	"tutorial/controller"
	"tutorial/database/migration"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context, roles *[]migration.Role) bool {
	_, _, keyword := controller.Param(c)
	log.Println(keyword)
	res := db.DB.
		Model(&migration.Role{}).
		Preload("RoleHasPermissions.Permission").
		Order("id ASC").
		Where("name like ?", "%"+keyword+"%").
		Find(&roles)

	if !controller.ResultQuery(c, res) {
		return false
	}

	return true
}
