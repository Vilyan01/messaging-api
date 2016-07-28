package data

import (
	"testing"
)

func TestNewDatabase(t *testing.T) {
	_, err := NewDatabase("messaging_test")
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
}
