package user_controller

import (
	"tutorial/controller"
	"tutorial/database/migration"
	"tutorial/helper"
	"tutorial/lang"
	"tutorial/repository/user_repository"
	"tutorial/resource/user_resource"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	defer controller.HandleDeferRequest(c)

	if !controller.Permission(c, "user_index", true) {
		return
	}

	var users []migration.User

	if !user_repository.Index(c, &users) {
		return
	}

	helper.SendData(
		c,
		helper.Lang(
			lang.EN["retrieved_successfully"],
			helper.H{"operator": lang.EN["user"]}),
		user_resource.Users(users),
	)
}
