package controllers

import (
	"net/http"
	"proyecto/src/users/application"
	"proyecto/src/users/domain/entities"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginUser *application.LoginUser
}

func NewLoginController(useCase *application.LoginUser) *LoginController {
	return &LoginController{loginUser: useCase}
}

func (lc *LoginController) Run(c *gin.Context) {
	var user entities.User

	// Validar JSON y enlazar a la estructura
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	// Ejecutar caso de uso de login
	authenticatedUser, err := lc.loginUser.Execute(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    authenticatedUser,
	})
}
