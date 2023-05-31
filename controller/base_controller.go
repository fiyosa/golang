package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"tutorial/config/db"
	"tutorial/config/env"
	"tutorial/database/migration"
	"tutorial/helper"
	"tutorial/lang"
	"tutorial/resource/user_resource"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const secret = "secret"

func GetToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func GenerateToken(user_id interface{}) string {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	res, err := token.SignedString([]byte(secret))
	if err != nil {
		return ""
	}
	return res
}

func IsTokenValid(tokenCheck string) bool {
	_, err := jwt.Parse(tokenCheck, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false
	}
	return true
}

func Param(c *gin.Context) (page int, limit int, keyword string) {
	_page := c.Query("page")
	_limit := c.Query("limit")
	keyword = c.Query("keyword")

	if _page == "" {
		_page = "1"
	}
	if _limit == "" {
		_limit = "10"
	}

	newPage, errpage := strconv.Atoi(_page)
	if errpage != nil {
		newPage = 1
	}
	if newPage < 1 {
		newPage = 1
	}

	newLimit, errLimit := strconv.Atoi(_limit)
	if errLimit != nil {
		newLimit = 10
	}
	if newLimit < 0 {
		newLimit = 10
	}

	page = newPage - 1
	limit = newLimit

	return
}

func GetUser(c *gin.Context) (migration.User, bool) {
	token := GetToken(c)
	if token == "" {
		helper.SendError(c, http.StatusBadRequest, helper.Lang(lang.EN["unauthorized_access"], helper.H{}))
		return migration.User{}, false
	}

	var oauth migration.Oauth
	if res := db.DB.Model(&migration.Oauth{}).Preload("User").First(&oauth, "token = ?", token); res.Error != nil {
		helper.SendError(c, http.StatusBadRequest, helper.Lang(lang.EN["unauthorized_access"], helper.H{}))
		return migration.User{}, false
	}

	if oauth.Revoked == true {
		helper.SendError(c, http.StatusBadRequest, helper.Lang(lang.EN["unauthorized_access"], helper.H{}))
		return migration.User{}, false
	}

	if !IsTokenValid(token) {
		helper.SendError(c, http.StatusBadRequest, helper.Lang(lang.EN["unauthorized_access"], helper.H{}))
		return migration.User{}, false
	}

	var user migration.User

	db.DB.
		Model(&migration.User{}).
		Preload("UserHasRoles.Role").
		Preload("UserHasRoles.Role.RoleHasPermissions.Permission").
		First(&user, "id = ?", oauth.User.Id)

	return user, true
}

func HandleDeferRequest(c *gin.Context) {
	errorMessage := recover()
	if errorMessage != nil {
		helper.SendCrash(
			c,
			fmt.Sprintf("%v", errorMessage),
		)
	}
}

func Permission(c *gin.Context, permission string, policy bool) bool {
	if !policy {
		helper.SendError(c, http.StatusBadRequest, helper.Lang(lang.EN["unauthorized_access"], helper.H{}))
		return false
	}

	user, status := GetUser(c)
	if !status {
		return false
	}

	newUser := user_resource.User(user)
	checkPermission := false

	for _, value := range newUser["permission"].([]string) {
		if value == permission || permission == "" {
			checkPermission = true
			break
		}
	}
	if !checkPermission {
		helper.SendError(c, http.StatusBadRequest, helper.Lang(lang.EN["unauthorized_access"], helper.H{}))
		return false
	}

	return true
}

func ResultQuery(c *gin.Context, res *gorm.DB) bool {
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			helper.SendError(
				c,
				http.StatusNotFound,
				helper.Lang(
					lang.EN["not_found"],
					helper.H{"operator": "Data"},
				))
			return false
		}

		if strings.Contains(res.Error.Error(), "duplicate key value violates unique constraint") {
			helper.SendError(c, http.StatusBadRequest, "Duplicate data entry")
			return false
		}

		resError := "Invalid query"
		if env.ENV_DEBUG {
			resError = res.Error.Error()
		}
		helper.SendError(
			c,
			http.StatusInternalServerError,
			resError,
		)
		return false

	}

	return true
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Request not found",
	})
}
