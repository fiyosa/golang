package role_controller

import (
	"tutorial/controller"
	"tutorial/helper"
	"tutorial/lang"
	"tutorial/repository/role_repository"
	"tutorial/validation/role_validation"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	defer controller.HandleDeferRequest(c)

	if !controller.Permission(c, "role_update", true) {
		return
	}

	var input role_validation.Update
	if !helper.ValidationErrors(c, &input) {
		return
	}

	if !role_repository.Update(c, input) {
		return
	}

	helper.SendSuccess(
		c,
		helper.Lang(
			lang.EN["updated_successfully"],
			helper.H{"operator": lang.EN["role"]},
		),
	)
}
