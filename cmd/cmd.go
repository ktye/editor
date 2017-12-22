// Package cmd contains helper function to implement an editor command.
//
// An editor command is a standalone program which communicates with the editor server
// by stdin and stdout.
package cmd

//go:generate godocdown -output README.md

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// Cmd is a common type which is embedded by editor command types.
type Cmd struct {
	Root       string     // Directory root.
	Name       string     // Window ID.
	Replace    string     // Window ID to replace.
	Tags       string     // New window tags.
	Default    string     // Default command for executed text in the body.
	Selections RuneRanges // Current selections.
	Type       string     // Content type "text" (default), "text/go", or "html".
	Clean      bool       // Mark buffer as clean.
	Text       string     // File content.
	args       []string   // Command arguments.
	in         io.Reader
	out        io.Writer
}

// Parse reads the header and data from stdin.
// It terminates the program with an appropriate error message in case of an error.
func (c *Cmd) Parse() error {
	// When testing, in and out may be set externally.
	// By default is is connected to stdin, and stdout.
	if c.in == nil {
		c.in = bufio.NewReader(os.Stdin)
	}
	if c.out == nil {
		c.out = os.Stdout
	}
	dec := json.NewDecoder(bufio.NewReader(c.in))
	if err := dec.Decode(c); err != nil {
		return fmt.Errorf("cannot decode header: %s", err)
	}
	if c.args == nil {
		c.args = os.Args[1:]
	}
	return nil
}

// Args returns the command line arguments.
// The program name is stripped.
// They are available after a call to parse.
func (c *Cmd) Args() []string {
	return c.args
}

// Encode encodes the command as json to w.
func (c *Cmd) Encode(w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(c)
}

// Exit writes the header with no error including the file content and terminates the program.
func (c *Cmd) Exit() {
	c.Encode(c.out)
	os.Exit(0)
}

// Forward forwards the request to a new process and decodes the response.
func (c *Cmd) Forward(prog string, args []string) error {
	execDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return fmt.Errorf("cannot get executable directory: %s", err)
	}

	cmd := exec.Command(filepath.Join(execDir, prog), args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	go func() {
		defer stdin.Close()
		enc := json.NewEncoder(stdin)
		enc.Encode(c)
	}()

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := json.NewDecoder(stdout).Decode(c); err != nil {
		return err
	}

	return cmd.Wait()
}

// Fatal writes the header containing the error to stdout and terminates the program.
// The function does nothing if the error is nil.
func (c *Cmd) Fatal(err error) {
	if err == nil {
		return
	}
	c.Name = "Error"
	c.Replace = ""
	c.Clean = false
	c.Text = err.Error()
	c.Encode(c.out)
	os.Exit(0)
}

// SetIO should be used only for testing commands.
func (c *Cmd) SetIO(in io.Reader, out io.Writer, args []string) {
	c.in = in
	c.out = out
	c.args = args
}

// relpath returns the relative directory path from the Name.
// It cuts everything after a '/' from the Name field.
func (c *Cmd) relpath() string {
	if idx := strings.LastIndexByte(c.Name, '/'); idx == -1 {
		return ""
	} else {
		return c.Name[:idx+1]
	}
}

// Directory returns the full path of the directory combining the Root and Name fields.
// The path has the format of the os, e.g. "c:\path\to\file.txt" on windows.
func (c *Cmd) Directory() string {
	return filepath.Clean(filepath.Join(c.Root, c.relpath()))
}

// Path return the full path of the file combining the Root and Name fields
// and the optional address after the first ':' in the Name field.
func (c *Cmd) Path() (string, string) {
	s := c.Name
	addr := ""
	if idx := strings.IndexByte(s, ':'); idx != -1 {
		addr = s[idx+1:]
		s = s[:idx]
	}
	return filepath.Join(c.Directory(), path.Base(s)), addr
}

// TargetPath substracts the root directory from the given full path.
// It converts all slashes to forward slashes.
func (c *Cmd) TargetPath(fullPath string) string {
	fullPath = filepath.Clean(fullPath)
	root := filepath.Clean(c.Root)
	if strings.HasPrefix(fullPath, root) {
		fullPath = fullPath[len(root):]
	}
	s := filepath.ToSlash(fullPath)
	if s == "" {
		return "/"
	}
	return s
}

// ArgPath returns the full path build from the Directory and a given relative path
// which may contain backwards slashes on windows.
// It also returns the optional address part after the first ':' in the relPath.
func (c *Cmd) ArgPath(relPath string) (string, string) {
	addr := ""
	if idx := strings.IndexByte(relPath, ':'); idx != -1 {
		addr = relPath[idx+1:]
		relPath = relPath[:idx]
	}
	return filepath.Clean(filepath.Join(c.Directory(), relPath)), addr
}

// Base returns the base name of the file from the Name field.
// It cuts everything before the last '/' and the first ':'.
func (c *Cmd) Base() string {
	s := c.Name
	if idx := strings.IndexByte(s, ':'); idx != -1 {
		s = s[:idx]
	}
	return filepath.Base(path.Base(s))
}

// NewTestRequest returns a Cmd for testing commands with the Root field filled.
func NewTestRequest() *Cmd {
	c := Cmd{
		Root: filepath.ToSlash(filepath.Join(os.Getenv("GOPATH"), "src/github.com/ktye/editor/cmd")),
	}
	return &c
}

// CompareTestResults compares a command output in the reader with the expected Cmd.
func CompareTestResults(got Cmd, expected Cmd) error {
	check := func(err error, name, a, b string) error {
		if err != nil {
			return err
		}
		if a != b {
			return fmt.Errorf("%s: expected: '%s', got '%s'\n", name, a, b)
		}
		return nil
	}
	var err error
	err = check(err, "Root", expected.Root, got.Root)
	err = check(err, "Name", expected.Name, got.Name)
	err = check(err, "Replace", expected.Replace, got.Replace)
	err = check(err, "Default", expected.Replace, got.Replace)
	err = check(err, "Tags", expected.Tags, got.Tags)
	err = check(err, "Type", expected.Type, got.Type)
	err = check(err, "Clean", strconv.FormatBool(expected.Clean), strconv.FormatBool(got.Clean))
	err = check(err, "Text", expected.Text, got.Text)
	if err != nil {
		return err
	}
	if expected.Selections.String() != got.Selections.String() {
		return fmt.Errorf("Selections: expected: '%v', got '%v'\n", expected.Selections, got.Selections)
	}
	return err
}
