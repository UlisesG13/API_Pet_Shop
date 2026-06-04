package application

import "proyecto/src/pets/domain"

type EditPet struct {
	db domain.IPet
}

func NewEditPet(db domain.IPet) *EditPet {
	return &EditPet{db: db}
}

func (ep *EditPet) Execute(id int, name, raza string) error {
	return ep.db.Edit(id, name, raza)
}
