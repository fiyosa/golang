package user_repository

import (
	"tutorial/config/db"
	"tutorial/controller"
	"tutorial/database/migration"
	"tutorial/helper"
	"tutorial/validation/user_validation"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context, input user_validation.Update) bool {
	var user migration.User
	id := helper.DecodeId(c.Param("id"))

	resUser := db.DB.
		Model(&migration.User{}).
		First(&user, "id = ?", id)
	if !controller.ResultQuery(c, resUser) {
		return false
	}

	user.Name = input.Name
	user.UpdatedAt = helper.TimeNow()
	resCheck := db.DB.Save(&user)
	if !controller.ResultQuery(c, resCheck) {
		return false
	}

	return true
}
