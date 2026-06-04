package controllers

import (
	"net/http"
	"proyecto/src/users/application"
	"proyecto/src/users/domain/entities"

	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	registerUser *application.RegisterUser
}

func NewRegisterController(useCase *application.RegisterUser) *RegisterController {
	return &RegisterController{registerUser: useCase}
}

func (rc *RegisterController) Run(c *gin.Context) {
	var user entities.User

	// Validar JSON y enlazar a la estructura
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	// Ejecutar caso de uso de registro
	err := rc.registerUser.Execute(user.Email, user.Password, user.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
