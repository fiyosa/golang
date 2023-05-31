package auth_controller

import (
	"tutorial/controller"
	"tutorial/helper"
	"tutorial/lang"
	"tutorial/resource/user_resource"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	user, status := controller.GetUser(c)
	if !status {
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
