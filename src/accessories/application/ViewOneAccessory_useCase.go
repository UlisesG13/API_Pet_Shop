package application

import (
	"proyecto/src/accessories/domain"
	"proyecto/src/accessories/domain/entities"
)

type ViewAccessory struct {
	db domain.IAccessory
}

func NewViewAccessory(db domain.IAccessory) *ViewAccessory {
	return &ViewAccessory{db: db}
}

func (vp *ViewAccessory) Execute(id int) (*entities.Accessory, error) {
	return vp.db.ViewOne(id)
}
