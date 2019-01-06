package repository

type Repository struct {
}

type Binary struct {
	Name    string
	Version string
}

func (r Repository) Resolve(name string, version string) Binary {
	return Binary{
		Name:    name,
		Version: version,
	}
}
