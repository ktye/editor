// Header is an editor command which displays the protocol header.
//
// It is used for debugging.
// It does not return the Text, but only the size of it.
package main

//go:generate godocdown -output README.md

import (
	"bytes"
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

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Args: %v\n", p.Args())
	fmt.Fprintf(&buf, "Root: %s\n", p.Root)
	fmt.Fprintf(&buf, "Name: %s\n", p.Name)
	fmt.Fprintf(&buf, "Replace: %s\n", p.Replace)
	fmt.Fprintf(&buf, "Tags: %s\n", p.Tags)
	fmt.Fprintf(&buf, "Default: %s\n", p.Default)
	fmt.Fprintf(&buf, "Selections: %s\n", p.Selections)
	fmt.Fprintf(&buf, "Type: %s\n", p.Type)
	fmt.Fprintf(&buf, "Clean: %v\n", p.Clean)
	fmt.Fprintf(&buf, "Text: %d bytes\n", len(p.Text))

	p.Name = "header"
	p.Text = string(buf.Bytes())
	return nil
}
