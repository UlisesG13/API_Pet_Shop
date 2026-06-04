package infrastructure

import (
	"proyecto/src/accessories/application"
	"proyecto/src/accessories/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func InitAccessories(db *MySQL, router *gin.Engine) {
	println("ACCESSORIES")

	// Instanciar casos de uso (Use Cases)
	saveAccessory := application.NewSaveAccessory(db)
	deleteAccessory := application.NewDeleteAccessory(db)
	viewAllAccessories := application.NewViewAccessories(db)
	viewAccessory := application.NewViewAccessory(db)
	editAccessory := application.NewEditAccessory(db)

	// Instanciar controladores (Handlers)
	saveAccessoryController := controllers.NewSaveAccessoryController(saveAccessory)
	deleteAccessoryController := controllers.NewDeleteAccessoryController(deleteAccessory)
	viewAccessoriesController := controllers.NewViewAccessoriesController(viewAllAccessories)
	viewAccessoryController := controllers.NewViewAccessoryController(viewAccessory)
	editAccessoryController := controllers.NewEditAccessoryController(editAccessory)

	// Configurar rutas
	SetupAccessoryRoutes(router, saveAccessoryController, deleteAccessoryController, viewAccessoriesController, viewAccessoryController, editAccessoryController)
}
