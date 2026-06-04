package infrastructure

import (
	"proyecto/src/users/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(
	router *gin.Engine,
	registerController *controllers.RegisterController,
	loginController *controllers.LoginController,
) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", registerController.Run)
		authGroup.POST("/login", loginController.Run)
	}
}
