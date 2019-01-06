package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository_ResolvesABinary(t *testing.T) {
	expected := Binary{
		Name:    "name",
		Version: "1.0.0",
	}

	repository := &Repository{
		binaries: map[string]Binary{
			"name": expected,
		},
	}

	actual, err := repository.Resolve("name", "1.0.0")
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestRepository_BinaryNotFound(t *testing.T) {
	repository := &Repository{}

	_, err := repository.Resolve("not_found", "1.0.0")

	assert.Equal(t, err, ErrBinaryNotFound)
}
