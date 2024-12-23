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
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
)

func Run(command string, args ...string) (int, string, string, error) {
	cmd := exec.Command(command, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	exitCode := 0
	err := cmd.Run()
	if err != nil {
		switch e := err.(type) {
		case *exec.ExitError:
			exitCode = e.ExitCode()
		default:
			exitCode = -1
		}
	}
	return exitCode, stdout.String(), stderr.String(), err
}

func RunTool(command string, args ...string) (int, string, string, error) {
	var err error
	tempDir, err := os.MkdirTemp("", "wintool-")
	if err != nil {
		return -1, "", "", err
	}
	defer os.RemoveAll(tempDir)

	tmpCommand := filepath.Join(tempDir, command)
	err = copyEmbeddedFile(command, tmpCommand)
	if err != nil {
		return -1, "", "", err
	}
	return Run(tmpCommand, args...)
}

func copyEmbeddedFile(src, dst string) error {
	data, err := Toolkit.ReadFile(filepath.Join("toolkit", src))
	if err != nil {
		return err
	}
	err = os.WriteFile(dst, data, 0700)
	return err
}
