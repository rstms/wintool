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
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon [FILENAME]",
	Short: "output start menu configuration",
	Long: `
Lists the start menu configuration using the nirsoft shman.exe tool. Optionally
filters the output by presence of FILENAME
`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {

		tempDir, err := os.MkdirTemp("", "daemon-")
		cobra.CheckErr(err)
		if KeepTempDirs {
			fmt.Println(tempDir)
		} else {
			defer os.RemoveAll(tempDir)
		}

		_, _, _, err = RunTool("daemon", "--help")
		cobra.CheckErr(err)

	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
