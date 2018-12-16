// Apl is an editor command which interpretes the selection as APL.
//
// Lines that start with a TAB are repeated and interpreted as APL.
// Lines that do not start with a TAB are ignored.
//
// Apl interpretes the current selection, or the complete file if nothing
// is selected.
//
// Example:
//		f←{(2=+⌿0=X∘.|X)⌿X←⍳⍵}
//		f 42
//	2 3 5 7 11 13 17 19 23 29 31 37 41
package main

//go:generate godocdown -output README.txt

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/ktye/editor/cmd"
	"github.com/ktye/iv/apl"
	"github.com/ktye/iv/apl/numbers"
	"github.com/ktye/iv/apl/operators"
	"github.com/ktye/iv/apl/primitives"
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

	all := false
	in := p.FirstSelectedText()
	if len(in) == 0 {
		in = p.Text
		all = true
	}
	var out bytes.Buffer

	//var out bytes.NewBuffer
	a := apl.New(&out)
	numbers.Register(a)
	primitives.Register(a)
	operators.Register(a)

	run := true
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) > 0 && s[0] == '\t' {
			fmt.Fprintln(&out, s)
			if run {
				if err := a.ParseAndEval(s); err != nil {
					fmt.Fprintln(&out, err)
					run = false
				}
			}
		}
	}

	if all {
		p.Text = string(out.Bytes())
	} else {
		p.ReplaceSelections(string(out.Bytes()))
	}
	return nil
}
