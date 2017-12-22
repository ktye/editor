// Go is an editor command which bundles some go commands.
//
// Go is set as the default command for files with .go extension.
//
// Subcommands
// 	-Fmt runs "goimports -w" on the current file.
// 	-Install runs "go install" in the window's directory.
// 	-Test runs "go test" in the window's directory.
// 	-def runs "godef" (github.com/rogpeppe/godef) using the current selection.
// 	-doc runs "go doc" with the current selection.
//
// Subcommands -Fmt, -Test and -Install also write the file to disk.
//
// Requirements
//
// These external commands must be on the path: go, goimports, godef.
package main

//go:generate godocdown -output README.md

import (
	"fmt"
	"os"

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

	args := p.Args()
	if len(args) == 0 {
		return fmt.Errorf("go command called without arguments")
	} else {
		switch args[0] {
		case "-Fmt":
			return p.fmt()
		case "-Install":
			return p.install()
		case "-Test":
			return p.test(args[1:])
		case "-doc":
			return p.doc()
		case "-def":
			return p.def()
		default:
			return fmt.Errorf("go: unknown arguments: %v", args)
		}
	}
}
