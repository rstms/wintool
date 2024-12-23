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

var outputAll bool
var outputRaw bool
var outputField string
var outputFile string

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

		oFile := os.Stdout

		if outputFile != "" {
			oFile, err = os.Create(outputFile)
			cobra.CheckErr(err)
			defer oFile.Close()
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if len(args) < 1 || strings.Index(line, args[0]) != -1 {
				formatted, err := formatOutput(line)
				cobra.CheckErr(err)
				fmt.Fprintln(oFile, formatted)
				if !outputAll {
					break
				}
			}
		}
		cobra.CheckErr(scanner.Err())
	},
}

func formatOutput(line string) (string, error) {
	if outputRaw {
		return line, nil
	}
	fields := strings.Split(line, "\t")
	record := make(map[string]string)
	record["name"] = fields[0]
	record["broken"] = fields[1]
	record["exe"] = fields[2]
	record["args"] = fields[3]
	record["dir"] = fields[4]
	record["window"] = fields[5]
	record["hotkey"] = fields[6]
	record["comment"] = fields[7]
	record["location"] = fields[8]
	record["modified"] = fields[9]
	record["icon"] = fields[10]
	record["shortcut"] = fields[11]

	field, ok := record[outputField]
	if ok {
		return field, nil
	}
	out := []string{}
	for k, v := range record {
		out = append(out, k+": "+v)
	}
	return strings.Join(out, "\n"), nil
}

func init() {
	startmenuCmd.Flags().BoolVarP(&outputAll, "all", "a", false, "output all matches")
	startmenuCmd.Flags().BoolVarP(&outputRaw, "raw", "r", false, "disable output formatting")
	startmenuCmd.Flags().StringVarP(&outputField, "field", "f", "", "single field")
	startmenuCmd.Flags().StringVarP(&outputFile, "output", "o", "", "output filename")
	rootCmd.AddCommand(startmenuCmd)
}
