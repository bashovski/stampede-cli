/*
	MIT License

	Copyright (c) 2020 Anur Bašić

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.
*/
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
