package provider

import (
	"fmt"
	"testing"

	"github.com/Masterminds/semver"
	"github.com/stretchr/testify/require"
)

func TestNewGithubProvider(t *testing.T) {
	provider := NewGithubProvider()

	for name, rules := range provider.binaryRules {
		name, rules := name, rules

		for _, rule := range rules {
			rule := rule

			t.Run(fmt.Sprintf("%s/%s", name, rule.Constraint), func(t *testing.T) {
				_, err := semver.NewConstraint(rule.Constraint)
				require.NoError(t, err)
			})
		}
	}
}
