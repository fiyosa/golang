package user_policy

import (
	"tutorial/controller"
	"tutorial/helper"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) bool {
	id := c.Param("id")

	user, status := controller.GetUser(c)
	if !status {
		return false
	}

	if id != helper.EncodeId(user.Id) {
		return false
	}

	return true
}
