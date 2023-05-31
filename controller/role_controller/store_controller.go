package role_controller

import (
	"tutorial/controller"
	"tutorial/helper"
	"tutorial/lang"
	"tutorial/repository/role_repository"
	"tutorial/validation/role_validation"

	"github.com/gin-gonic/gin"
)

func Store(c *gin.Context) {
	defer controller.HandleDeferRequest(c)

	if !controller.Permission(c, "role_store", true) {
		return
	}

	var input role_validation.Store
	if !helper.ValidationErrors(c, &input) {
		return
	}

	if !role_repository.Store(c, input) {
		return
	}

	helper.SendSuccess(
		c,
		helper.Lang(
			lang.EN["saved_successfully"],
			helper.H{"operator": lang.EN["role"]},
		),
	)
}
