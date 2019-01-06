package provider

import "fmt"

const urlPrefix = "https://github.com/{{.FullName}}/releases/download/%s"

// NewGithubProvider
func NewGithubProvider() *Provider {
	return &Provider{
		binaryRules: map[string][]BinaryRule{
			"dep": {
				{
					Constraint: ">0.3.0",
					URL:        fmt.Sprintf(urlPrefix, "v{{.Version}}/dep-{{.Os}}-{{.Arch}}.zip"),
					File:       "dep-{{.Os}}-{{.Arch}}",
				},
				{
					Constraint: "<=0.3.0",
					URL:        fmt.Sprintf(urlPrefix, "v{{.Version}}/dep-{{.Os}}-{{.Arch}}"),
					File:       "dep-{{.Os}}-{{.Arch}}",
				},
			},
		},
	}
}
