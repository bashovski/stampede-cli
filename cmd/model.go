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

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "Create a base module for your new model!",
	Long: `By creating a new model, you also create a new module containing routes for the model.
Everything is autowired instantly, hence you don't need to spend additional time on explicitly setting up everything else.
You can also create a migration for the model so it's ready for use with a database.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("model called")
		template.ReplaceFileTemplates(args, "")
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
}
