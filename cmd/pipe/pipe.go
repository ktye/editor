// Pipe is an editor command which executes a pipe line of external commands
// and handles in and output depending on the first character of the argument.
//
// 	First arugment  input      output
//
// 	!               none       Name+Errors
// 	|               Selection  Selection
// 	<               none       Selection
// 	>               Selection  Name+Errors
//
// If the Selection is empty on input the complete file is used.
//
// Pipe expects a single argument which will be splitted following the
// quoting rules.
// If the arguments contains a pipe, it connects multiple programs.
//
// Example
//      pipe "!grep -n alpha file | wc -l"
package main

//go:generate godocdown -output README.md

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
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

	args := p.Args()
	if n := len(args); n != 1 {
		return fmt.Errorf("pipe expects a single argument, got %d", n)
	}
	if len(args[0]) < 2 {
		return fmt.Errorf("pipe argument is too short")
	}
	mode := args[0][0]
	line := args[0][1:]

	if err := os.Chdir(p.Directory()); err != nil {
		return err
	}

	var stdin io.Reader
	var selReader io.Reader
	var out bytes.Buffer
	if p.Selections.Total() == 0 {
		selReader = strings.NewReader(p.Text)
	} else {
		selReader = strings.NewReader(p.CombinedSelectedText())
	}
	substitute := false
	switch mode {
	case '!':
	case '|':
		stdin = selReader
		substitute = true
	case '>':
		stdin = selReader
	case '<':
		substitute = true
	default:
		return fmt.Errorf("pipe: fist argument must be !, |, > or <")
	}

	if err := execPipe(line, stdin, &out); err != nil {
		return err
	}

	if substitute == true {
		p.ReplaceSelections(string(out.Bytes()))
	} else {
		s := string(out.Bytes())

		// Dont's create a new window, if there is no output.
		if len(s) == 0 {
			return nil
		} else {
			p.Name += "+Errors"
			p.Text = s
		}
	}
	return nil
}

func execPipe(line string, stdin io.Reader, stdout io.Writer) error {
	var err error
	pipes := cmd.SplitQuotedPipe(line)
	var cmds []*exec.Cmd
	for _, arg := range pipes {
		if argv, err := cmd.SplitQuoted(arg); err != nil {
			return err
		} else if len(argv) > 0 {
			cmd := exec.Command(argv[0], argv[1:]...)
			cmds = append(cmds, cmd)
		}
	}
	if len(cmds) == 0 {
		return fmt.Errorf("pipe does not contain any commands")
	}

	cmds[0].Stdin = stdin
	last := len(cmds) - 1
	for i, cmd := range cmds[:last] {
		if cmds[i+1].Stdin, err = cmd.StdoutPipe(); err != nil {
			return err
		}
	}
	cmds[last].Stdout = stdout

	for _, cmd := range cmds {
		if err := cmd.Start(); err != nil {
			return err
		}
	}
	if err := cmds[last].Wait(); err != nil {
		return err
	}
	return nil
}
