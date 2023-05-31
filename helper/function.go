package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"tutorial/config/env"
	"tutorial/lang"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type H map[string]interface{}

func Lang(msg string, args H) string {
	for key, value := range args {
		msg = strings.ReplaceAll(msg, ":"+key, fmt.Sprintf("%v", value))
	}
	return msg
}

func ValidationErrors(c *gin.Context, input interface{}) bool {
	if err := c.ShouldBind(input); err != nil {
		errorMessages := map[string][]string{}
		errorMessage := ""

		switch err.(type) {
		case *json.UnmarshalTypeError:
			field := err.(*json.UnmarshalTypeError).Field
			errorMessages[field] = append(errorMessages[field], "Json binding error: "+field+" type error")
			errorMessage = "Json binding error: " + field + " type error"
			break

		case validator.ValidationErrors:
			for _, e := range err.(validator.ValidationErrors) {
				error := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages[e.Field()] = append(errorMessages[e.Field()], error)
				if errorMessage == "" {
					errorMessage = error
				}
			}
			break

		default:
			if env.ENV_DEBUG {
				errorMessage = err.Error()
			} else {
				errorMessage = Lang(lang.EN["something_went_wrong"], H{})
			}
			break
		}

		SendValidation(c, errorMessage, errorMessages)

		return false
	}

	return true
}

func HandleDeferInit() {
	errorMessage := recover()
	if errorMessage == nil {
		return
	}
	log.Fatalf("Error system: %v", errorMessage)
}
