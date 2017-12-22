// Sleep is an editor command which sleeps given number of seconds.
//
// It is used to test the built-in `ps` and `kill` commands.
package main

//go:generate godocdown -output README.md

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ktye/editor/cmd"
)

type Prog struct {
	cmd.Cmd
}

func main() {
	var p Prog
	if err := p.Run(); err != nil {
		p.Fatal(err)
	}
	p.Exit()
}

func (p *Prog) Run() error {
	if err := p.Parse(); err != nil {
		return err
	}

	seconds := 10
	args := p.Args()
	if len(args) > 0 {
		if n, err := strconv.Atoi(args[0]); err == nil {
			seconds = n
		}
	}

	time.Sleep(time.Second * time.Duration(seconds))

	p.Name = "+Sleep"
	p.Replace = ""
	p.Tags = ""
	p.Clean = true
	p.Text = fmt.Sprintf("slept for %d seconds.\n", seconds)
	return nil
}
