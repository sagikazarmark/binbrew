package main

import (
	"fmt"
	"os"

	"github.com/sagikazarmark/binbrew/internal/cli/command"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
// nolint: gochecknoglobals
var rootCmd = &cobra.Command{
	Use:   "binbrew",
	Short: "Binary installer",
}

// nolint: gochecknoinits
func init() {
	rootCmd.AddCommand(
		command.NewInstallCommand(),
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
