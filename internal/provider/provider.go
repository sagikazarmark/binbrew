package provider

import (
	"bytes"
	"errors"
	"runtime"
	"strings"
	"text/template"

	"github.com/Masterminds/semver"
)

var (
	ErrBinaryNotFound            = errors.New("binary not found")
	ErrNoMatchingVersionFound    = errors.New("no matching version found")
	ErrLatestVersionNotSupported = errors.New("latest version is not supported yet")
)

// Provider contains a set of binary rules.
type Provider struct {
	binaryRules map[string][]BinaryRule
	vanityNames map[string]string
}

// BinaryRule contains all information to resolve a binary.
type BinaryRule struct {
	VersionConstraint string
	Template          BinaryTemplate
}

// BinaryTemplate contains the binary information.
type BinaryTemplate struct {
	Name        string
	Homepage    string
	Description string
	URL         string
	File        string
}

// Binary is a result of a binary resolution.
type Binary struct {
	Name     string
	FullName string
	Version  string
	URL      string
	File     string
}

// TemplateContext is passed to URL and File templates.
type TemplateContext struct {
	Name     string
	FullName string
	Version  string
	Os       string
	Arch     string
}

// Resolve resolves a binary.
func (r *Provider) Resolve(name string, version string) (*Binary, error) {
	if strings.ToLower(version) == "latest" {
		return nil, ErrLatestVersionNotSupported
	}

	// Resolve a vanity name (if any)
	if originalName, ok := r.vanityNames[name]; ok {
		name = originalName
	}

	v, err := semver.NewVersion(version)
	if err != nil {
		return nil, err
	}

	binaryRules, ok := r.binaryRules[name]
	if !ok {
		return nil, ErrBinaryNotFound
	}

	for _, binaryRule := range binaryRules {
		constraint, err := semver.NewConstraint(binaryRule.VersionConstraint)
		if err != nil {
			return nil, err // This should never happen
		}

		if constraint.Check(v) {
			binaryName := binaryRule.Template.Name
			if binaryName == "" {
				binaryName = name
				nameSegments := strings.SplitN(name, "/", 2)
				if len(nameSegments) > 1 {
					binaryName = nameSegments[1]
				}
			}

			tplCtx := TemplateContext{
				Name:     binaryName,
				FullName: name,
				Version:  version,
				Os:       runtime.GOOS,
				Arch:     runtime.GOARCH,
			}

			urlTemplate, err := template.New("").Funcs(funcMap).Parse(binaryRule.Template.URL)
			if err != nil {
				return nil, err
			}

			var buf bytes.Buffer

			err = urlTemplate.Execute(&buf, tplCtx)
			if err != nil {
				return nil, err
			}

			binary := &Binary{
				Name:     binaryName,
				FullName: name,
				Version:  version,
				URL:      buf.String(),
				File:     binaryName,
			}

			if binaryRule.Template.File != "" {
				fileTemplate, err := template.New("").Funcs(funcMap).Parse(binaryRule.Template.File)
				if err != nil {
					return nil, err
				}

				buf.Reset()

				err = fileTemplate.Execute(&buf, tplCtx)
				if err != nil {
					return nil, err
				}

				binary.File = buf.String()
			}

			return binary, nil
		}
	}

	return nil, ErrNoMatchingVersionFound
}
