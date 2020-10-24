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
	"fmt"
	"github.com/bashovski/stampede-cli/pkg/template"

	"github.com/spf13/cobra"
)

// ctrlCmd represents the ctrl command
var ctrlCmd = &cobra.Command{
	Use:   "ctrl",
	Short: "Create a new controller using a model's name",
	Long: `Creates a base controller module which is ready to be used in your application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ctrl called: ", args)
		for i := 0; i < len(args); i++ {
			fmt.Println("arg ", i, ": ", args[i])
			template.ReplaceFileTemplate(args[i], "Controller")
		}
	},
}

func init() {
	rootCmd.AddCommand(ctrlCmd)
}
