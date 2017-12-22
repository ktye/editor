package main

import (
	"os/exec"
	"strings"
)

func (p *program) gitStatus() error {
	c := exec.Command("git", "status")
	if out, err := c.CombinedOutput(); err != nil {
		p.Name += "+Errors"
		p.Text = err.Error() + "\n" + string(out)
	} else {
		if strings.HasSuffix(p.Name, "+Git") == false {
			p.Name += "+Git"
		}
		p.Text = string(out)
		p.Tags = "-add -commit"
		p.Default = "git -add"
	}
	return nil
}
