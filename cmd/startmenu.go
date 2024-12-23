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
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var startmenuCmd = &cobra.Command{
	Use:   "startmenu [FILENAME]",
	Short: "output start menu configuration",
	Long: `
Lists the start menu configuration using the nirsoft shman.exe tool. Optionally
filters the output by presence of FILENAME
`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {

		tempDir, err := os.MkdirTemp("", "startmenu-")
		cobra.CheckErr(err)
		if KeepTempDirs {
			if Debug {
				log.Printf("outputDir: %s", tempDir)
			}
		} else {
			defer os.RemoveAll(tempDir)
		}

		filename := filepath.Join(tempDir, "links")
		_, _, _, err = RunTool("shman.exe", "/stab", filename)
		cobra.CheckErr(err)

		file, err := os.Open(filename)
		cobra.CheckErr(err)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if len(args) < 1 || strings.Index(line, args[0]) != -1 {
				fmt.Println(line)
			}
		}
		cobra.CheckErr(scanner.Err())
	},
}

func init() {
	rootCmd.AddCommand(startmenuCmd)
}
