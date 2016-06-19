package data

type Client struct {
	Email        string        `json:"email"`
	Applications []Application `json:"applications"`
}
