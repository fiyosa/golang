package user_repository

import (
	"tutorial/config/db"
	"tutorial/controller"
	"tutorial/database/migration"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) bool {
	var user *migration.User
	id := c.Param("id")

	resUser := db.DB.First(&user, id)
	if !controller.ResultQuery(c, resUser) {
		return false
	}

	resCheck := db.DB.Delete(&user, id)
	if !controller.ResultQuery(c, resCheck) {
		return false
	}

	return true
}
