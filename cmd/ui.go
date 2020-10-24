package cmd

import (
	"github.com/bashovski/stampede-cli/pkg/tail"
	"github.com/spf13/cobra"
)

// uiCmd represents the ui command
var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Boots up and serves Stampede UI",
	Long: `Serves pre-made UI app to showcase Stampede's possibilities. Also could be easily used as a very good starting point
once you want to easily continue building a very stable application as it offers tons of features already.'`,
	Run: func(cmd *cobra.Command, args []string) {
		tail.Command("scripts/ui")
	},
}

func init() {
	rootCmd.AddCommand(uiCmd)
}
