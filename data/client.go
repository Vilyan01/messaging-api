package data

type Client struct {
	Model                        // Inherit the common fields and methods from Model
	Email          string        `json:"email"`
	HashedPassword string        `json:"-"`
	Applications   []Application `json:"applications"`
}

/*
	Creates a new client object and returns it.
*/
func NewClient(e string) Client {
	return Client{Model: Model{TableName: "clients"}, Email: e}
}
