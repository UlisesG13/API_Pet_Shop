package application

import (
	"proyecto/src/pets/domain"
	"proyecto/src/pets/domain/entities"
)

type ViewPet struct {
	db domain.IPet
}

func NewViewPet(db domain.IPet) *ViewPet {
	return &ViewPet{db: db}
}

func (vp *ViewPet) Execute(id int) (*entities.Pet, error) {
	return vp.db.ViewOne(id)
}
