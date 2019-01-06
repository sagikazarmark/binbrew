package repository

import "testing"

func TestRepository_ResolvesABinary(t *testing.T) {
	repository := Repository{}

	repository.Resolve("name", "1.0.0")
}
