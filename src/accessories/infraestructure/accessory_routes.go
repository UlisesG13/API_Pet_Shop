package infrastructure

import (
	"proyecto/src/accessories/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAccessoryRoutes(
	router *gin.Engine,
	saveAccessoryController *controllers.CreateAccessoryController,
	deleteAccessoryController *controllers.DeleteAccessoryController,
	viewAccessoriesController *controllers.ViewAccessoriesController,
	viewAccessoryController *controllers.ViewAccessoryController,
	editAccessoryController *controllers.EditAccessoriesController,
) {
	accessoryGroup := router.Group("/accessories")
	{
		accessoryGroup.POST("", saveAccessoryController.Run)
		accessoryGroup.GET("", viewAccessoriesController.Run)
		accessoryGroup.GET("/:id", viewAccessoryController.Run)
		accessoryGroup.DELETE("/:id", deleteAccessoryController.Run)
		accessoryGroup.PUT("/:id", editAccessoryController.Run)
	}
}
