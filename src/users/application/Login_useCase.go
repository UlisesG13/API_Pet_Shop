package application

import (
	"fmt"
	"proyecto/src/users/domain"
	"proyecto/src/users/domain/entities"

	"golang.org/x/crypto/bcrypt"
)

type LoginUser struct {
	db domain.IUser
}

func NewLoginUser(db domain.IUser) *LoginUser {
	return &LoginUser{db: db}
}

func (lu *LoginUser) Execute(email, password string) (*entities.User, error) {
	// Validaciones básicas
	if email == "" || password == "" {
		return nil, fmt.Errorf("email y contraseña son requeridos")
	}

	// Buscar usuario
	user, err := lu.db.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("usuario no encontrado: %w", err)
	}

	// Verificar contraseña
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("contraseña incorrecta")
	}

	// Limpiar la contraseña antes de retornar
	user.Password = ""

	return user, nil
}
