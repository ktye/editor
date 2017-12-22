// Run is an editor command which executes a
// program associated with the extension of the current window Name.
//
// Associations
//
// 	Name       command line
// 	file.awk   awk -f file.awk
// 	file.go    go run file.go
// 	file.sed   sed -f file.sed
// 	file.sh    sh file.sh
//
// Flags
//
// All arguments before the last are appended to the command line.
//
// InOutput
//
// If run is called without arguments, it's stdin is empty and it's output
// is written to an +Errors window.
// If run is called with a file arguemnt (e.g. `file.txt`), the input is read
// from disk from a file build by ArgPath from the directory of the script
// and the given file argument.
// It's output is written to a window with the name of the TargetPath of the
// file argument.
// If the argument is prefixed with `<`, as in `run <file.txt`, only the input
// is read from the file, and the output is written to an `+Errors` window.
// If the argument is prefixed with `>`, as in `run >file.txt`, the input is
// empty.
//
// Usage
//
// Edit the script in it's own window (e.g. `/path/to/script.sed`) and the
// file to be modified in another window (e.g. `/path/to/file.txt`).
// Now execute "run file.txt" in the script window.
package main

//go:generate godocdown -output README.md

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ktye/editor/cmd"
)

type program struct {
	cmd.Cmd
}

func main() {
	var p program
	if err := p.Run(); err != nil {
		p.Fatal(err)
	}
	p.Exit()
}

func (p *program) Run() error {
	if err := p.Parse(); err != nil {
		return err
	}

	if err := os.Chdir(p.Directory()); err != nil {
		return err
	}

	// tap maps from file extension to argv
	ext := filepath.Ext(p.Base())
	tab := map[string][]string{
		".awk": {"awk", "-f"},
		".go":  {"go", "run"},
		".sed": {"sed", "-f"},
		".sh":  {"sh"},
	}

	argv, ok := tab[ext]
	if !ok {
		return fmt.Errorf("file extension '%s' is not associated with any program", ext)
	}

	args := p.Args()
	if len(args) > 1 {
		argv = append(argv, args[:len(args)-1]...)
		args = args[len(args)-2:]
	}
	var stdin io.Reader
	var outname string
	if len(args) > 0 {
		name := args[0]
		if strings.HasPrefix(name, ">") {
			name = name[1:]
			outname = name
		} else {
			if strings.HasPrefix(name, "<") {
				name = name[1:]
			} else {
				outname = name
			}
			argPath, _ := p.ArgPath(name)
			if f, err := os.Open(argPath); err != nil {
				return err
			} else {
				defer f.Close()
				stdin = f
			}
		}
	}
	argv = append(argv, p.Base())

	c := exec.Command(argv[0], argv[1:]...)
	c.Stdin = stdin
	out, err := c.CombinedOutput()
	p.Text = ""
	if outname == "" || err != nil {
		p.Name += "+Errors"
		if err != nil {
			p.Text += err.Error() + "\n"
		}
	} else {
		argPath, _ := p.ArgPath(outname)
		p.Name = p.TargetPath(argPath)
	}
	p.Text += string(out)
	return nil
}
