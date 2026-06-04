package domain

import "proyecto/src/users/domain/entities"

type IUser interface {
	Register(email, password, name string) error
	Login(email, password string) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
}
