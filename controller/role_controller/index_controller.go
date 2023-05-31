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

func Index(c *gin.Context) {
	defer controller.HandleDeferRequest(c)

	if !controller.Permission(c, "role_index", true) {
		return
	}

	var roles []migration.Role

	if !role_repository.Index(c, &roles) {
		return
	}

	helper.SendData(
		c,
		helper.Lang(
			lang.EN["retrieved_successfully"],
			helper.H{"operator": lang.EN["role"]}),
		role_resource.Index(roles),
	)
}
