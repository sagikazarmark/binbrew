package provider

import "errors"

var (
	ErrBinaryNotFound = errors.New("binary not found")
)

type Provider struct {
	binaries map[string]Binary
}

type Binary struct {
	Name    string
	Version string
}

func (r *Provider) Resolve(name string, version string) (Binary, error) {
	binary, ok := r.binaries[name]
	if !ok {
		return binary, ErrBinaryNotFound
	}

	return binary, nil
}
