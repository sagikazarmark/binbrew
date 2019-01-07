package provider

import "fmt"

const urlPrefix = "https://github.com/{{ .FullName }}/releases/download/%s"

// NewGithubProvider
func NewGithubProvider() *Provider {
	return &Provider{
		binaryRules: map[string][]BinaryRule{
			"golang/dep": {
				{
					VersionConstraint: ">0.3.0",
					Template: BinaryTemplate{
						URL:  fmt.Sprintf(urlPrefix, "v{{.Version}}/dep-{{ .Os }}-{{ .Arch }}"),
						File: "dep-{{ .Os }}-{{ .Arch }}",
					},
				},
				{
					VersionConstraint: "<=0.3.0",
					Template: BinaryTemplate{
						URL:  fmt.Sprintf(urlPrefix, "v{{.Version}}/dep-{{ .Os }}-{{ .Arch }}.zip"),
						File: "dep-{{ .Os }}-{{ .Arch }}",
					},
				},
			},
		},
	}
}
