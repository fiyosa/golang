package user_controller

import (
	"tutorial/controller"
	"tutorial/helper"
	"tutorial/lang"
	"tutorial/policy/user_policy"
	"tutorial/repository/user_repository"
	"tutorial/validation/user_validation"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	defer controller.HandleDeferRequest(c)

	if !controller.Permission(c, "user_update", user_policy.Update(c)) {
		return
	}

	var input user_validation.Update
	if !helper.ValidationErrors(c, &input) {
		return
	}

	if !user_repository.Update(c, input) {
		return
	}

	helper.SendSuccess(
		c,
		helper.Lang(
			lang.EN["updated_successfully"],
			helper.H{"operator": lang.EN["user"]},
		),
	)
}
