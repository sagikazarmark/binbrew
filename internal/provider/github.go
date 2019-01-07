// nolint: lll
package provider

import "fmt"

const urlPrefix = "https://github.com/{{ .FullName }}/releases/download/%s"

// NewGithubProvider
func NewGithubProvider() *Provider {
	return &Provider{
		binaryRules: map[string][]BinaryRule{
			"gobuffalo/packr": {
				{
					VersionConstraint: "*",
					Template: BinaryTemplate{
						Description: "The simple and easy way to embed static files into Go binaries.",
						URL:         fmt.Sprintf(urlPrefix, "v{{ .Version }}/packr_{{ .Version }}_{{ .Os }}_{{ .Arch }}.tar.gz"),
						File:        "packr",
					},
				},
			},
			"golang/dep": {
				{
					VersionConstraint: ">0.3.0",
					Template: BinaryTemplate{
						Homepage:    "https://github.com/golang/dep",
						Description: "Go dependency management tool",
						URL:         fmt.Sprintf(urlPrefix, "v{{ .Version }}/dep-{{ .Os }}-{{ .Arch }}"),
						File:        "dep-{{ .Os }}-{{ .Arch }}",
					},
				},
				{
					VersionConstraint: "<=0.3.0",
					Template: BinaryTemplate{
						Homepage:    "https://github.com/golang/dep",
						Description: "Go dependency management tool",
						URL:         fmt.Sprintf(urlPrefix, "v{{ .Version }}/dep-{{ .Os }}-{{ .Arch }}.zip"),
						File:        "dep-{{ .Os }}-{{ .Arch }}",
					},
				},
			},
			"golangci/golangci-lint": {
				{
					VersionConstraint: "*",
					Template: BinaryTemplate{
						Homepage:    "http://golangci.com",
						Description: "Linters Runner for Go",
						URL:         fmt.Sprintf(urlPrefix, "v{{ .Version }}/golangci-lint-{{ .Version }}-{{ .Os }}-{{ .Arch }}.tar.gz"),
						File:        "golangci-lint-{{ .Version }}-{{ .Os }}-{{ .Arch }}/golangci-lint",
					},
				},
			},
			"google/protobuf": {
				{
					VersionConstraint: "*",
					Template: BinaryTemplate{
						Name:        "protoc",
						Description: "Protocol Buffers - Google's data interchange format",
						URL:         fmt.Sprintf(urlPrefix, "v{{ .Version }}/protoc-{{ .Version }}-{{ .Os|protobuf_goos }}-{{ .Arch|protobuf_goarch }}.zip"),
						File:        "bin/protoc",
					},
				},
			},
			"goph/licensei": {
				{
					VersionConstraint: "*",
					Template: BinaryTemplate{
						Homepage:    "https://github.com/goph/licensei",
						Description: "Library and various tools for working with project licenses",
						URL:         fmt.Sprintf(urlPrefix, "v{{ .Version }}/licensei_{{ .Os }}_{{ .Arch }}.tar.gz"),
						File:        "licensei",
					},
				},
			},
			"goreleaser/goreleaser": {
				{
					VersionConstraint: "*",
					Template: BinaryTemplate{
						Homepage:    "https://goreleaser.github.io/",
						Description: "Deliver Go binaries as fast and easily as possible",
						URL:         fmt.Sprintf(urlPrefix, "v{{ .Version }}/goreleaser_{{ .Os | title }}_{{ .Arch | goarch }}.tar.gz"),
						File:        "goreleaser",
					},
				},
			},
			"gotestyourself/gotestsum": {
				{
					VersionConstraint: "*",
					Template: BinaryTemplate{
						Homepage:    "https://github.com/gotestyourself/gotestsum",
						Description: "The go test runner you never knew you always wanted. Human readable test output, JUnit XML for CI integration, summary of the test run results",
						URL:         fmt.Sprintf(urlPrefix, "v{{ .Version }}/gotestsum_{{ .Version }}_{{ .Os }}_{{ .Arch }}.tar.gz"),
						File:        "gotestsum",
					},
				},
			},
			"golang-migrate/migrate": {
				{
					VersionConstraint: "*",
					Template: BinaryTemplate{
						Description: "Database migrations. CLI and Golang library.",
						URL:         fmt.Sprintf(urlPrefix, "v{{ .Version }}/migrate.{{ .Os }}-{{ .Arch }}.tar.gz"),
						File:        "migrate.{{ .Os }}-{{ .Arch }}",
					},
				},
			},
		},
	}
}
