package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func (p *program) gitCommit(args []string) error {
	if len(args) > 0 {
		msg := strings.Join(args, "\n")
		return p.doCommit(msg)
	} else if msg := p.CombinedSelectedText(); msg != "" {
		return p.doCommit(msg)
	} else {
		return fmt.Errorf("no commit message")
	}
}

func (p *program) doCommit(msg string) error {
	c := exec.Command("git", "commit", "-m", msg)
	if out, err := c.CombinedOutput(); err != nil {
		return fmt.Errorf("%s\n%s", err, out)
	} else {
		p.Text = string(out)
		return nil
	}
}
