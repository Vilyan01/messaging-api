package data

import (
	"reflect"
)

type Model struct {
	TableName string
	Id        int `json:"id"`
}

func (m Model) Save() {
}
