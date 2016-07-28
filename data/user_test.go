package data

import (
	"testing"
)

func TestCreate(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}

func TestRemove(t *testing.T) {

}

func TestFind(t *testing.T) {
	_, err := FindUserByID(24124124)
	// Error should be nil
	if err == nil {
		t.Error("Should throw error for a user with ID of 24124124 but it didn't")
	}
}
