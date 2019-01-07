package main

import (
	"fmt"
	"os"

	"github.com/sagikazarmark/binbrew/internal/cli/command"
	"github.com/spf13/cobra"
)

// Provisioned by ldflags
// nolint: gochecknoglobals
var (
	version    string
	commitHash string
	buildDate  string
)

// rootCmd represents the base command when called without any subcommands
// nolint: gochecknoglobals
var rootCmd = &cobra.Command{
	Use:     "binbrew",
	Short:   "Binary installer",
	Version: version,
}

// nolint: gochecknoinits
func init() {
	rootCmd.AddCommand(
		command.NewInstallCommand(),
	)

	rootCmd.SetVersionTemplate(fmt.Sprintf("Binbrew version %s (%s) built on %s\n", version, commitHash, buildDate))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
