package application

import (
	"proyecto/src/pets/domain"
	"proyecto/src/pets/domain/entities"
)

type ViewPets struct {
	db domain.IPet
}

func NewViewPets(db domain.IPet) *ViewPets {
	return &ViewPets{db: db}
}

func (vp *ViewPets) Execute() ([]entities.Pet, error) {
	return vp.db.ViewAll()
}
