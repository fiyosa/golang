package role_repository

import (
	"net/http"
	"tutorial/config/db"
	"tutorial/controller"
	"tutorial/database/migration"
	"tutorial/helper"
	"tutorial/lang"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) bool {
	id := helper.DecodeId(c.Param("id"))

	if id == -1 {
		helper.SendError(c, http.StatusBadRequest, helper.Lang(
			lang.EN["not_found"],
			helper.H{"operator": "Data"},
		))
		return false
	}

	resRole := db.DB.Delete(&migration.Role{}, "id = ?", id)

	if !controller.ResultQuery(c, resRole) {
		return false
	}

	if resRole.RowsAffected > 0 {
		helper.SendError(c, http.StatusBadRequest, helper.Lang(
			lang.EN["not_found"],
			helper.H{"operator": "Data"},
		))
		return false
	}

	return true
}
