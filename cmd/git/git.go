// Git is an editor command which wraps git.
//
// Called with no argument, it executes "git status" and presents
// the output in a new window.
// The window has it's default command set to "git -add".
//
// Double-clicking any files in the status window will run git add on these files.
// Files or blocks of files may also be selected and executing "-add" on the tag bar
// will add these files. Git -add strips prefixes such as "modified:" and any whitespace.
//
// To commit changes, type a commit message in the window, select it and execute
// "-commit" on the tag bar.
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
		return p.gitStatus()
	} else if args[0] == "-add" {
		// This allows to execute "-commit" from the status window,
		// which has it's default command set to "git -add".
		if len(args) > 1 && args[1] == "-commit" {
			return p.gitCommit(args[2:])
		}
		// This allows to execute "-add" with selected files
		// from the status window, which has it's default command set
		// to "git -add".
		if len(args) > 1 && args[1] == "-add" {
			return p.gitAdd(args[2:])
		}
		return p.gitAdd(args[1:])
	} else if args[0] == "-commit" {
		return p.gitCommit(args[1:])
	} else {
		return fmt.Errorf("unknown arguments to git: %v", args)
	}
}
