package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_ResolvesABinary(t *testing.T) {
	repository := Repository{}

	actual := repository.Resolve("name", "1.0.0")

	expected := Binary{
		Name:    "name",
		Version: "1.0.0",
	}

	assert.Equal(t, expected, actual)
}
