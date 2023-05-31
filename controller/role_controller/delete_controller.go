package role_controller

import (
	"tutorial/controller"
	"tutorial/helper"
	"tutorial/lang"
	"tutorial/repository/role_repository"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	defer controller.HandleDeferRequest(c)

	if !controller.Permission(c, "role_delete", true) {
		return
	}

	if !role_repository.Delete(c) {
		return
	}

	helper.SendSuccess(
		c,
		helper.Lang(
			lang.EN["deleted_successfully"],
			helper.H{"operator": lang.EN["role"]},
		),
	)
}
