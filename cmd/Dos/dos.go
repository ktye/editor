// Dos writes files with MS-DOS line endings.
package main

//go:generate godocdown -output README.md

import (
	"fmt"
	"os"
	"regexp"

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

	if args := p.Args(); len(args) != 0 {
		return fmt.Errorf("Write does not accept arguments")
	}

	re := regexp.MustCompile("\r?\n")
	p.Text = re.ReplaceAllString(p.Text, "\r\n")

	return p.Forward("Write", nil)
}
