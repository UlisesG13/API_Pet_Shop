package controllers

import (
	"net/http"
	"proyecto/src/accessories/application"
	"proyecto/src/accessories/domain/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditAccessoriesController struct {
	petEditor *application.EditAccessory
}

func NewEditAccessoryController(editor *application.EditAccessory) *EditAccessoriesController {
	return &EditAccessoriesController{
		petEditor: editor,
	}
}

func (cp *EditAccessoriesController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pet ID"})
		return
	}

	var accessory entities.Accessory
	if err := c.ShouldBindJSON(&accessory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Ejecutar la actualización en la base de datos
	err = cp.petEditor.Execute(id, accessory.Name, accessory.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet updated successfully"})
}
