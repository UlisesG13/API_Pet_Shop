package application

import "proyecto/src/pets/domain"

type DeletePet struct {
	db domain.IPet
}

func NewDeletePet(db domain.IPet) *DeletePet {
	return &DeletePet{db: db}
}

func (dp *DeletePet) Execute(id int) error {
	return dp.db.Delete(id)
}
