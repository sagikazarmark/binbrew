// nolint: lll
package provider

import "fmt"

const urlPrefix = "https://github.com/{{ .FullName }}/releases/download/%s"

// NewGithubProvider
func NewGithubProvider() *Provider {
	return &Provider{
		vanityNames: map[string]string{
			"dep":           "golang/dep",
			"golangci-lint": "golangci/golangci-lint",
			"protobuf":      "google/protobuf",
			"protoc":        "google/protobuf",
			"protolock":     "nilslice/protolock",
			"prototool":     "uber/prototool",
			"goreleaser":    "goreleaser/goreleaser",
			"gotestsum":     "gotestyourself/gotestsum",
			"jq":            "stedolan/jq",
		},
		binaryRules: map[string][]BinaryRule{
			"gobuffalo/packr": {
				{
					VersionConstraint: "*",
					Template: BinaryTemplate{
						Description: "The simple and easy way to embed static files into Go binaries.",
						URL:         fmt.Sprintf(urlPrefix, "v{{ .Version }}/packr_{{ .Version }}_{{ .Os }}_{{ .Arch }}.tar.gz"),
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
			"nilslice/protolock": {
				{
					VersionConstraint: "0.10.0",
					Template: BinaryTemplate{
						Description: "Protocol Buffer companion tool. Track your .proto files and prevent changes to messages and services which impact API compatibility.",
						URL:         fmt.Sprintf(urlPrefix, "v{{ .Version }}/protolock.20190101T225741Z.{{ .Os }}-{{ .Arch }}.tgz"),
					},
				},
			},
			"stedolan/jq": {
				{
					VersionConstraint: ">=1.5",
					Template: BinaryTemplate{
						Name:        "jq",
						Description: "Command-line JSON processor",
						URL:         fmt.Sprintf(urlPrefix, "jq-{{ .Version }}/jq-{{ jq_osarch .Os .Arch }}"),
						File:        "jq-{{ jq_osarch .Os .Arch }}",
					},
				},
				{
					VersionConstraint: "<1.5,>=1.3",
					Template: BinaryTemplate{
						Name:        "jq",
						Description: "Command-line JSON processor",
						URL:         fmt.Sprintf(urlPrefix, "jq-{{ .Version }}/jq-{{ .Os|protobuf_goos }}-{{ .Arch|jq_goarch }}"),
						File:        "jq-{{ .Os|protobuf_goos }}-{{ .Arch|jq_goarch }}",
					},
				},
			},
			"uber/prototool": {
				{
					VersionConstraint: "*",
					Template: BinaryTemplate{
						Description: "Your Swiss Army Knife for Protocol Buffers",
						URL:         fmt.Sprintf(urlPrefix, "v{{ .Version }}/prototool-{{ .Os | title }}-{{ .Arch | goarch }}"),
						File:        "prototool-{{ .Os | title }}-{{ .Arch | goarch }}",
					},
				},
			},
		},
	}
}
