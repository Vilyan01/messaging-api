package data

type Client struct {
	Model                 // Inherit the common fields and methods from Model
	Email          string `json:"email",gorethink:"email"`
	HashedPassword string `json:"-",gorethink:"hashed_password"`
}

/*
	Creates a new client object and returns it.
*/
func NewClient(e string, h string) Client {
	return Client{Model: Model{TableName: "clients"}, Email: e, HashedPassword: h}
}
