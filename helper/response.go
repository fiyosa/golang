package helper

import (
	"net/http"
	"tutorial/config/env"
	"tutorial/lang"

	"github.com/gin-gonic/gin"
)

func SendSuccess(c *gin.Context, message interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message,
	})
}

func SendData(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
		"message": message,
	})
}

func SendValidation(c *gin.Context, message string, errorMessage interface{}) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   errorMessage,
		"message": message,
	})
}

func SendError(c *gin.Context, status int, message interface{}) {
	c.AbortWithStatusJSON(status, gin.H{
		"success": false,
		"message": message,
	})
}

func SendCrash(c *gin.Context, message interface{}) {
	if env.ENV_DEBUG {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": message,
		})

	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": Lang(lang.EN["something_went_wrong"], H{}),
		})
	}
}
