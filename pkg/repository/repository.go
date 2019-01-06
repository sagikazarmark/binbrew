package repository

import "errors"

var (
	ErrBinaryNotFound = errors.New("binary not found")
)

type Repository struct {
}

type Binary struct {
	Name    string
	Version string
}

func (r *Repository) Resolve(name string, version string) (Binary, error) {
	if name == "not_found" {
		return Binary{}, ErrBinaryNotFound
	}

	return Binary{
		Name:    name,
		Version: version,
	}, nil
}
