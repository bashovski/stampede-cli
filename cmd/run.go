package cmd

import (
	"github.com/bashovski/stampede-cli/pkg/tail"
	"github.com/spf13/cobra"
)

// runCmd represents the run command which boots up the Stampede server and its UI.
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs an application which inherits Stampede's structure/codebase.",
	Long: `Runs both the server and the UI of Stampede project.`,
	Run: func(cmd *cobra.Command, args []string) {
		tail.Command("scripts/run")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
