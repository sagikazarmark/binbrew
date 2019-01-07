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
						File:        "packr_{{ .Version }}_{{ .Os }}_{{ .Arch }}/packr",
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
			"google/protobuf": {
				{
					VersionConstraint: "*",
					Template: BinaryTemplate{
						Name:        "protoc",
						Description: "Protocol Buffers - Google's data interchange format",
						URL:         fmt.Sprintf(urlPrefix, "v{{ .Version }}/protoc-{{ .Version }}-{{ .Os|protobuf_goos }}-{{ .Arch|protobuf_goarch }}.zip"), // nolint: lll
						File:        "bin/protoc",
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
						File:        "goreleaser_{{ .Os | title }}_{{ .Arch | goarch }}/goreleaser",
					},
				},
			},
			"mattes/migrate": {
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
