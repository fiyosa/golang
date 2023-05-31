package auth_controller

import (
	"tutorial/config/db"
	"tutorial/controller"
	"tutorial/database/migration"
	"tutorial/helper"
	"tutorial/lang"
	"tutorial/validation/auth_validation"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	defer controller.HandleDeferRequest(c)

	var input auth_validation.Register
	if !helper.ValidationErrors(c, &input) {
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.SendCrash(c, err.Error())
		return
	}
	input.Password = string(hashedPassword)

	dateNow := helper.TimeNow()
	regiter := migration.User{
		Email:     input.Email,
		Password:  input.Password,
		Name:      input.Name,
		CreatedAt: dateNow,
		UpdatedAt: dateNow,
	}

	resQuery := db.DB.Create(&regiter)
	if !controller.ResultQuery(c, resQuery) {
		return
	}

	helper.SendSuccess(
		c,
		helper.Lang(
			lang.EN["saved_successfully"],
			helper.H{"operator": lang.EN["user"]},
		),
	)
}
