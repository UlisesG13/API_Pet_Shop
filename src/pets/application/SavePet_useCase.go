package application

import (
	"proyecto/src/pets/domain"
)

type SavePet struct {
	db domain.IPet
}

func NewSavePet(db domain.IPet) *SavePet {
	return &SavePet{db: db}
}
func (sv *SavePet) Execute(name, raza string) error {
	return sv.db.Save(name, raza)
}
