package application

import "proyecto/src/accessories/domain"

type DeleteAccessory struct {
	db domain.IAccessory
}

func NewDeleteAccessory(db domain.IAccessory) *DeleteAccessory {
	return &DeleteAccessory{db: db}
}

func (dp *DeleteAccessory) Execute(id int) error {
	return dp.db.Delete(id)
}
