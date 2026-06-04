package controllers

import (
	"net/http"
	"proyecto/src/pets/application"
	"proyecto/src/pets/domain/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditPetController struct {
	petEditer *application.EditPet
}

func NewEditPetController(editer *application.EditPet) *EditPetController {
	return &EditPetController{
		petEditer: editer,
	}
}

func (cp *EditPetController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pet ID"})
		return
	}
	var pet entities.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	err = cp.petEditer.Execute(id, pet.Name, pet.Raza)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet updated successfully"})
}
