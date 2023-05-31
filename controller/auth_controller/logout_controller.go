package auth_controller

import (
	"tutorial/config/db"
	"tutorial/controller"
	"tutorial/database/migration"
	"tutorial/helper"
	"tutorial/lang"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	_, status := controller.GetUser(c)
	if !status {
		return
	}

	var oauth migration.Oauth
	db.DB.Model(migration.Oauth{}).First(&oauth, "token = ?", controller.GetToken(c))
	oauth.Revoked = true
	db.DB.Save(&oauth)

	helper.SendSuccess(
		c,
		helper.Lang(
			lang.EN["logout"],
			helper.H{},
		),
	)
}
