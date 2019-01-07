package provider

import (
	"bytes"
	"fmt"
	"net/url"
	"runtime"
	"testing"
	"text/template"

	"github.com/Masterminds/semver"
	"github.com/stretchr/testify/require"
)

func TestNewGithubProvider(t *testing.T) {
	provider := NewGithubProvider()

	for name, rules := range provider.binaryRules {
		name, rules := name, rules

		for _, rule := range rules {
			rule := rule

			t.Run(fmt.Sprintf("%s/%s", name, rule.VersionConstraint), func(t *testing.T) {
				_, err := semver.NewConstraint(rule.VersionConstraint)
				require.NoError(t, err)

				dummyContext := TemplateContext{
					Name:     "repo",
					FullName: "org/repo",
					Version:  "1.0.0",
					Os:       runtime.GOOS,
					Arch:     runtime.GOARCH,
				}

				urlTemplate, err := template.New("").Funcs(funcMap).Parse(rule.Template.URL)
				require.NoError(t, err)

				var buf bytes.Buffer

				err = urlTemplate.Execute(&buf, dummyContext)
				require.NoError(t, err)

				_, err = url.Parse(buf.String())
				require.NoError(t, err)

				fileTemplate, err := template.New("").Funcs(funcMap).Parse(rule.Template.File)
				require.NoError(t, err)

				buf.Reset()

				err = fileTemplate.Execute(&buf, dummyContext)
				require.NoError(t, err)
			})
		}
	}
}
