package entities

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name"`
}

func NewUser(email, password, name string) *User {
	return &User{
		Email:    email,
		Password: password,
		Name:     name,
	}
}
