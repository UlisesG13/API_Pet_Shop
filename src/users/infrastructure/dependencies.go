package infrastructure

import (
	"proyecto/src/users/application"
	"proyecto/src/users/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func InitUsers(db *MySQL, router *gin.Engine) {
	println("USERS")

	// Instanciar casos de uso (Use Cases)
	registerUser := application.NewRegisterUser(db)
	loginUser := application.NewLoginUser(db)

	// Instanciar controladores (Handlers)
	registerController := controllers.NewRegisterController(registerUser)
	loginController := controllers.NewLoginController(loginUser)

	// Configurar rutas
	SetupUserRoutes(router, registerController, loginController)
}
