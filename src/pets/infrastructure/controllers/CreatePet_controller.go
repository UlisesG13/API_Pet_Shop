package controllers

import (
	"net/http"
	"proyecto/src/pets/application"
	"proyecto/src/pets/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreatePetController struct {
	petSaver *application.SavePet
}

func NewSavePetController(useCase *application.SavePet) *CreatePetController {
	return &CreatePetController{petSaver: useCase}
}

func (cp *CreatePetController) Run(c *gin.Context) {
	var pet entities.Pet
	// Validar JSON y enlazar a la estructura
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	// Guardar la mascota
	err := cp.petSaver.Execute(pet.Name, pet.Raza)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save pet: " + err.Error()})
		return
	}

	// Retornar la mascota creada
	c.JSON(http.StatusCreated, gin.H{"message": "Pet saved successfully"})
}
