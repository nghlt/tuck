package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Set by goreleaser via ldflags for tagged releases; defaults to the
// current development version otherwise.
var (
	Version = "0.1.1"
	Commit  = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("tuck %s (%s)\n", Version, Commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Enable `tuck --version` / `tuck -v` in addition to `tuck version`,
	// using the same output format.
	rootCmd.Version = Version
	rootCmd.SetVersionTemplate(fmt.Sprintf("tuck %s (%s)\n", Version, Commit))
}
