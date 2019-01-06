package repository

import "errors"

var (
	ErrBinaryNotFound = errors.New("binary not found")
)

type Repository struct {
	binaries map[string]Binary
}

type Binary struct {
	Name    string
	Version string
}

func (r *Repository) Resolve(name string, version string) (Binary, error) {
	binary, ok := r.binaries[name]
	if !ok {
		return binary, ErrBinaryNotFound
	}

	return binary, nil
}
