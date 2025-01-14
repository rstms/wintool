/*
Copyright © 2024 Matt Krueger <mkrueger@rstms.net>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"embed"
	"os"

	"github.com/spf13/cobra"
)

//go:embed toolkit
var Toolkit embed.FS
var KeepTempDirs bool
var Debug bool
var Verbose bool
var Quiet bool
var Version string = "0.0.16"

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wintool",
	Short: "windows desktop management tool",
	Long:  "wintool v" + Version + `  Utility functions for windows desktop management.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&KeepTempDirs, "keep", "k", false, "keep temp directories")
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "enable debug output")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "be chatty")
	rootCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "be taciturn")
}
