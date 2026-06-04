package entities

type Accessory struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var increment = 0

func NewAccessory(name, description string) *Accessory {
	increment++
	accessory := Accessory{Id: increment, Name: name, Description: description}
	return &accessory
}
