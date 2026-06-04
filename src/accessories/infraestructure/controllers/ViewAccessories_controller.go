package controllers

import (
	"net/http"
	"proyecto/src/accessories/application"

	"github.com/gin-gonic/gin"
)

type ViewAccessoriesController struct {
	vas *application.ViewAccessories
}

func NewViewAccessoriesController(useCase *application.ViewAccessories) *ViewAccessoriesController {
	return &ViewAccessoriesController{vas: useCase}
}

func (vpc *ViewAccessoriesController) Run(c *gin.Context) {
	accessories, err := vpc.vas.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if accessories == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No accessories found"})
		return
	}
	c.JSON(http.StatusOK, accessories)
}
