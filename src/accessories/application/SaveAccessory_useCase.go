package application

import (
	"proyecto/src/accessories/domain"
)

type SaveAccessory struct {
	db domain.IAccessory
}

func NewSaveAccessory(db domain.IAccessory) *SaveAccessory {
	return &SaveAccessory{db: db}
}
func (sv *SaveAccessory) Execute(name, description string) error {
	return sv.db.Save(name, description)
}
