package server

import (
	"tutorial/config/env"
	"tutorial/config/route"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.Default()

	route.InitRoute(router)

	router.Run(env.ENV_PORT)
}
