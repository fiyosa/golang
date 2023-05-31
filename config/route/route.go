package route

import (
	"tutorial/config/cors"
	"tutorial/controller"
	"tutorial/controller/auth_controller"
	"tutorial/controller/role_controller"
	"tutorial/controller/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(router *gin.Engine) {
	router.Use(cors.CORSMiddleware())
	prefix := router.Group("/api")

	prefix.POST("/auth/login", auth_controller.Login)
	prefix.POST("/auth/register", auth_controller.Register)
	prefix.DELETE("/auth/logout", auth_controller.Logout)

	prefix.GET("/auth/user", auth_controller.User)
	prefix.GET("/user", user_controller.Index)
	prefix.GET("/user/:id", user_controller.Show)
	prefix.PUT("/user/:id", user_controller.Update)
	prefix.DELETE("/user/:id", user_controller.Delete)

	prefix.GET("/role", role_controller.Index)
	prefix.GET("/role/:id", role_controller.Show)
	prefix.POST("/role", role_controller.Store)
	prefix.PUT("/role/:id", role_controller.Update)
	prefix.DELETE("/role/:id", role_controller.Delete)

	router.NoRoute(controller.NotFound)
}
