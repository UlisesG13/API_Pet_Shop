package infrastructure

import (
	"proyecto/src/pets/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupPetRoutes(
	router *gin.Engine,
	savePetController *controllers.CreatePetController,
	deletePetController *controllers.DeletePetController,
	viewPetsController *controllers.ViewPetsController,
	viewPetController *controllers.ViewPetController,
	editPetController *controllers.EditPetController,
) {
	petGroup := router.Group("/pets")
	{
		petGroup.POST("", savePetController.Run)
		petGroup.GET("", viewPetsController.Run)
		petGroup.GET("/:id", viewPetController.Run)
		petGroup.DELETE("/:id", deletePetController.Run)
		petGroup.PUT("/:id", editPetController.Run)
	}
}
