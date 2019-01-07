package provider

import (
	"fmt"
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

				_, err = template.New("").Parse(rule.Template.URL)
				require.NoError(t, err)

				_, err = template.New("").Parse(rule.Template.File)
				require.NoError(t, err)
			})
		}
	}
}
