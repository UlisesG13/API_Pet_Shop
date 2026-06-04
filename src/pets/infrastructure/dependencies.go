package infrastructure

import (
	"proyecto/src/pets/application"
	"proyecto/src/pets/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func InitPets(db *MySQL, router *gin.Engine) {
	println("PETS")

	// Instanciar casos de uso (Use Cases)
	savePet := application.NewSavePet(db)
	deletePet := application.NewDeletePet(db)
	viewAllPets := application.NewViewPets(db)
	viewPet := application.NewViewPet(db)
	editPet := application.NewEditPet(db)

	// Instanciar controladores (Handlers)
	savePetController := controllers.NewSavePetController(savePet)
	deletePetController := controllers.NewDeletePetController(deletePet)
	viewPetsController := controllers.NewViewPetsController(viewAllPets)
	viewPetController := controllers.NewViewPetController(viewPet)
	editPetController := controllers.NewEditPetController(editPet)

	// Configurar rutas
	SetupPetRoutes(router, savePetController, deletePetController, viewPetsController, viewPetController, editPetController)

}
