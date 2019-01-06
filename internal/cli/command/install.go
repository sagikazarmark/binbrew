package command

import (
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
	return nil
}
