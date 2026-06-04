package controllers

import (
	"net/http"
	"proyecto/src/accessories/application"

	"github.com/gin-gonic/gin"
)

type CreateAccessoryController struct {
	saver *application.SaveAccessory
}

func NewSaveAccessoryController(useCase *application.SaveAccessory) *CreateAccessoryController {
	return &CreateAccessoryController{saver: useCase}
}

func (ca *CreateAccessoryController) Run(c *gin.Context) {
	var json struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ca.saver.Execute(json.Name, json.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Accessory saved successfully"})
}
