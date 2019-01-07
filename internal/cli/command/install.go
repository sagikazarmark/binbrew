package command

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-getter" // nolint: goimports
	"github.com/sagikazarmark/binbrew/internal/provider"
	"github.com/spf13/cobra"
)

type installOptions struct {
	name    string
	version string
}

// NewInstallCommand returns a cobra command for `binbrew install`.
func NewInstallCommand() *cobra.Command {
	options := installOptions{}

	cmd := &cobra.Command{
		Use:   "install NAME VERSION",
		Short: "Install a binary",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			options.name = args[0]
			options.version = args[1]

			return runInstall(options)
		},
	}

	return cmd
}

func runInstall(options installOptions) error {
	p := provider.NewGithubProvider()

	binary, err := p.Resolve(options.name, options.version)
	if err != nil {
		return err
	}

	tmp := filepath.Join("bin/tmp", binary.Name)

	err = os.MkdirAll(tmp, 0744)
	if err != nil {
		return err
	}

	err = getter.GetAny(tmp, binary.URL)
	if err != nil {
		return err
	}

	input, err := ioutil.ReadFile(filepath.Join(tmp, binary.File))
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join("bin", binary.Name), input, 0644)
	if err != nil {
		return err
	}

	return nil
}
