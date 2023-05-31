package user_controller

import (
	"tutorial/controller"
	"tutorial/helper"
	"tutorial/lang"
	"tutorial/policy/user_policy"
	"tutorial/repository/user_repository"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	defer controller.HandleDeferRequest(c)

	if !controller.Permission(c, "user_delete", user_policy.Delete(c)) {
		return
	}

	if !user_repository.Delete(c) {
		return
	}

	helper.SendSuccess(
		c,
		helper.Lang(
			lang.EN["deleted_successfully"],
			helper.H{"operator": lang.EN["user"]},
		),
	)
}
