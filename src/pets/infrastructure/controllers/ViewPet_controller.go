package controllers

import (
	"net/http"
	"strconv"

	"proyecto/src/pets/application"

	"github.com/gin-gonic/gin"
)

type ViewPetController struct {
	vp *application.ViewPet
}

func NewViewPetController(useCase *application.ViewPet) *ViewPetController {
	return &ViewPetController{vp: useCase}
}

func (vpc *ViewPetController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pet ID"})
		return
	}

	pet, err := vpc.vp.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pet)
}
