// Write is the editor command which writes files to disk.
//
// The program writes the command's Text to the file name combining Root and Name.
// On success it returns it's input and marks the file as clean.
package main

//go:generate godocdown -output README.md

import "github.com/ktye/editor/cmd"

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

	return p.Write()
}
