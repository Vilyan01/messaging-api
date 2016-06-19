package data

type User struct {
	Model
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
}

func NewUser(e string, d string) User {
	return User{Model: Model{"users"}, Email: e, DisplayName: d}
}
