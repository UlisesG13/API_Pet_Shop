package application

import (
	"fmt"
	"proyecto/src/users/domain"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	db domain.IUser
}

func NewRegisterUser(db domain.IUser) *RegisterUser {
	return &RegisterUser{db: db}
}

func (ru *RegisterUser) Execute(email, password, name string) error {
	// Validaciones básicas
	if email == "" || password == "" || name == "" {
		return fmt.Errorf("email, contraseña y nombre son requeridos")
	}

	if len(password) < 6 {
		return fmt.Errorf("la contraseña debe tener al menos 6 caracteres")
	}

	// Hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error al encriptar la contraseña: %w", err)
	}

	// Guardar usuario
	return ru.db.Register(email, string(hashedPassword), name)
}
