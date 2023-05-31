package role_controller

import (
	"tutorial/controller"
	"tutorial/database/migration"
	"tutorial/helper"
	"tutorial/lang"
	"tutorial/repository/role_repository"
	"tutorial/resource/role_resource"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	defer controller.HandleDeferRequest(c)

	if !controller.Permission(c, "role_show", true) {
		return
	}

	var role migration.Role

	if !role_repository.Show(c, &role) {
		return
	}

	helper.SendData(
		c,
		helper.Lang(
			lang.EN["retrieved_successfully"],
			helper.H{"operator": lang.EN["role"]}),
		role_resource.Show(role),
	)
}
