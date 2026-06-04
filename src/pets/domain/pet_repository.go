package domain

import "proyecto/src/pets/domain/entities"

type IPet interface {
	Save(name, raza string) error
	ViewOne(id int) (*entities.Pet, error)
	ViewAll() ([]entities.Pet, error)
	Delete(id int) error
	Edit(id int, name, raza string) error
}
