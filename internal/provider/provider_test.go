package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProvider_ResolvesABinary(t *testing.T) {
	expected := Binary{
		Name:    "name",
		Version: "1.0.0",
	}

	repository := &Provider{
		binaries: map[string]Binary{
			"name": expected,
		},
	}

	actual, err := repository.Resolve("name", "1.0.0")
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestProvider_BinaryNotFound(t *testing.T) {
	repository := &Provider{}

	_, err := repository.Resolve("not_found", "1.0.0")

	assert.Equal(t, err, ErrBinaryNotFound)
}
