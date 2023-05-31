package auth_controller

import (
	"net/http"
	"tutorial/config/db"
	"tutorial/controller"
	"tutorial/database/migration"
	"tutorial/helper"
	"tutorial/lang"
	"tutorial/validation/auth_validation"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	defer controller.HandleDeferRequest(c)

	var input auth_validation.Login
	if !helper.ValidationErrors(c, &input) {
		return
	}

	var user migration.User
	resUser := db.DB.Model(&migration.User{}).First(&user, "email = ?", input.Username)
	if !controller.ResultQuery(c, resUser) {
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		helper.SendError(c, http.StatusBadRequest, helper.Lang(lang.EN["auth_failed"], helper.H{}))
		return
	}

	dateNow := helper.TimeNow()
	token := migration.Oauth{
		UserId:    user.Id,
		Token:     controller.GenerateToken(helper.EncodeId(user.Id)),
		Revoked:   false,
		CreatedAt: dateNow,
		UpdatedAt: dateNow,
	}

	resToken := db.DB.Create(&token)
	if !controller.ResultQuery(c, resToken) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"access_token": token.Token,
		"message": helper.Lang(
			lang.EN["retrieved_successfully"],
			helper.H{"operator": lang.EN["user"]},
		),
	})
}
