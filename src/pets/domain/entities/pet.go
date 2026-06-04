package entities

type Pet struct {
	Id   int
	Name string `json:"name"`
	Raza string `json:"raza"`
}

var increment = 0

func NewPet(name, raza string) *Pet {
	increment++
	pet := Pet{Id: increment, Name: name, Raza: raza}
	return &pet
}
