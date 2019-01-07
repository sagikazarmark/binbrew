package provider

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProvider_ResolvesABinary(t *testing.T) {
	expected := &Binary{
		Name:     "repo",
		FullName: "org/repo",
		Version:  "1.0.0",
		URL: fmt.Sprintf(
			"https://github.com/org/repo/releases/download/1.0.0/repo_1.0.0_%s_%s.tar.gz",
			runtime.GOOS,
			runtime.GOARCH,
		),
		File: fmt.Sprintf(
			"repo_1.0.0_%s_%s",
			runtime.GOOS,
			runtime.GOARCH,
		),
	}

	repository := &Provider{
		binaryRules: map[string][]BinaryRule{
			"org/repo": {
				BinaryRule{
					VersionConstraint: "*",
					Template: BinaryTemplate{
						URL:  "https://github.com/{{ .FullName }}/releases/download/{{ .Version }}/repo_{{ .Version }}_{{ .Os }}_{{ .Arch }}.tar.gz", // nolint: lll
						File: "repo_{{ .Version }}_{{ .Os }}_{{ .Arch }}",
					},
				},
			},
		},
	}

	actual, err := repository.Resolve("org/repo", "1.0.0")
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestProvider_ResolvesABinaryWithACustomName(t *testing.T) {
	expected := &Binary{
		Name:     "other_name",
		FullName: "org/repo",
		Version:  "1.0.0",
		URL: fmt.Sprintf(
			"https://github.com/org/repo/releases/download/1.0.0/repo_1.0.0_%s_%s.tar.gz",
			runtime.GOOS,
			runtime.GOARCH,
		),
		File: fmt.Sprintf(
			"repo_1.0.0_%s_%s",
			runtime.GOOS,
			runtime.GOARCH,
		),
	}

	repository := &Provider{
		binaryRules: map[string][]BinaryRule{
			"org/repo": {
				BinaryRule{
					VersionConstraint: "*",
					Template: BinaryTemplate{
						Name: "other_name",
						URL:  "https://github.com/{{ .FullName }}/releases/download/{{ .Version }}/repo_{{ .Version }}_{{ .Os }}_{{ .Arch }}.tar.gz", // nolint: lll
						File: "repo_{{ .Version }}_{{ .Os }}_{{ .Arch }}",
					},
				},
			},
		},
	}

	actual, err := repository.Resolve("org/repo", "1.0.0")
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestProvider_BinaryNotFound(t *testing.T) {
	repository := &Provider{}

	_, err := repository.Resolve("not_found", "1.0.0")

	assert.Equal(t, err, ErrBinaryNotFound)
}

func TestProvider_LatestVersionNotSupported(t *testing.T) {
	repository := &Provider{}

	_, err := repository.Resolve("org/repo", "latest")

	assert.Equal(t, err, ErrLatestVersionNotSupported)
}
