package application

import (
	"proyecto/src/accessories/domain"
	"proyecto/src/accessories/domain/entities"
)

type ViewAccessories struct {
	db domain.IAccessory
}

func NewViewAccessories(db domain.IAccessory) *ViewAccessories {
	return &ViewAccessories{db: db}
}

func (vp *ViewAccessories) Execute() ([]entities.Accessory, error) {
	return vp.db.ViewAll()
}
