package user_policy

import (
	"tutorial/controller"
	"tutorial/helper"
	"tutorial/resource/user_resource"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) bool {
	id := c.Param("id")

	user, err := controller.GetUser(c)
	if !err {
		return false
	}

	newUser := user_resource.User(user)
	if id != helper.EncodeId(user.Id) && !helper.IncludesStr(newUser["roles"].([]string), []string{"Super Admin"}) {
		return false
	}

	return true
}
