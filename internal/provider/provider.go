package provider

import (
	"errors"

	"github.com/Masterminds/semver"
)

var (
	ErrBinaryNotFound = errors.New("binary not found")
)

type BinaryRule struct {
	Constraint string
	URL        string
	File       string
}

type Provider struct {
	binaryRules map[string][]BinaryRule
}

type Binary struct {
	Name    string
	Version string
}

func (r *Provider) Resolve(name string, version string) (Binary, error) {
	v := semver.MustParse(version)
	binaryRules, ok := r.binaryRules[name]
	if !ok {
		return Binary{}, ErrBinaryNotFound
	}

	for _, binaryRule := range binaryRules {
		constraint, err := semver.NewConstraint(binaryRule.Constraint)
		if err != nil {
			panic(err) // This should never happen
		}

		if constraint.Check(v) {
			return Binary{
				Name:    name,
				Version: version,
			}, nil
		}
	}

	return Binary{}, errors.New("constraint not found")
}
