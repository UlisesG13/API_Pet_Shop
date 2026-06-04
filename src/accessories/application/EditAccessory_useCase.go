package application

import "proyecto/src/accessories/domain"

type EditAccessory struct {
	db domain.IAccessory
}

func NewEditAccessory(db domain.IAccessory) *EditAccessory {
	return &EditAccessory{db: db}
}

func (ep *EditAccessory) Execute(id int, name string, raza string) error {
	return ep.db.Edit(id, name, raza)
}
