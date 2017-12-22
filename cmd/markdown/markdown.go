// +build ignore

// Markdown is an editor command which renders a markdown file as html.
package main

//go:generate godocdown -output README.md

import (
	"github.com/ktye/editor/cmd"
	"gopkg.in/russross/blackfriday.v2"
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

	out := blackfriday.Run([]byte(p.Text))
	p.Type = "html"
	p.Name += "+html"
	p.Text = string(out)
	p.Clean = true
	return nil
}
