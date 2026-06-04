package domain

import "proyecto/src/accessories/domain/entities"

type IAccessory interface {
	Save(name, description string) error
	ViewOne(id int) (*entities.Accessory, error)
	ViewAll() ([]entities.Accessory, error)
	Delete(id int) error
	Edit(id int, name, description string) error
}
