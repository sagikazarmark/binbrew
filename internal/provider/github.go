package provider

import "fmt"

const urlPrefix = "https://github.com/{{ .Name }}/releases/download/%s"

// NewGithubProvider
func NewGithubProvider() *Provider {
	return &Provider{
		binaryRules: map[string][]BinaryRule{
			"dep": {
				{
					VersionConstraint: ">0.3.0",
					URLTemplate:       fmt.Sprintf(urlPrefix, "v{{.Version}}/dep-{{ .Os }}-{{ .Arch }}.zip"),
					FileTemplate:      "dep-{{ .Os }}-{{ .Arch }}",
				},
				{
					VersionConstraint: "<=0.3.0",
					URLTemplate:       fmt.Sprintf(urlPrefix, "v{{.Version}}/dep-{{ .Os }}-{{ .Arch }}"),
					FileTemplate:      "dep-{{ .Os }}-{{ .Arch }}",
				},
			},
		},
	}
}
