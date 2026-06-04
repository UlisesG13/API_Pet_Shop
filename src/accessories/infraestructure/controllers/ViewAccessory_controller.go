package controllers

import (
	"net/http"
	"strconv"

	"proyecto/src/accessories/application"

	"github.com/gin-gonic/gin"
)

type ViewAccessoryController struct {
	vp *application.ViewAccessory
}

func NewViewAccessoryController(useCase *application.ViewAccessory) *ViewAccessoryController {
	return &ViewAccessoryController{vp: useCase}
}

func (vpc *ViewAccessoryController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Accessory ID"})
		return
	}

	accessory, err := vpc.vp.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, accessory)
}
