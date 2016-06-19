package data

type Application struct {
	Name   string `json:"name"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
	Users  []User `json:"users"`
}
