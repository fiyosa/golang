package user_controller

import (
	"tutorial/config/db"
	"tutorial/controller"
	"tutorial/database/migration"
	"tutorial/helper"
	"tutorial/lang"
	"tutorial/resource/user_resource"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	defer controller.HandleDeferRequest(c)

	if !controller.Permission(c, "user_show", true) {
		return
	}

	var user migration.User
	id := c.Param("id")

	res := db.DB.
		Model(&migration.User{}).
		Preload("UserHasRoles.Role").
		Preload("UserHasRoles.Role.RoleHasPermissions.Permission").
		First(&user, "id = ?", helper.DecodeId(id))

	if !controller.ResultQuery(c, res) {
		return
	}

	helper.SendData(
		c,
		helper.Lang(
			lang.EN["retrieved_successfully"],
			helper.H{"operator": lang.EN["user"]},
		),
		user_resource.User(user),
	)
}
