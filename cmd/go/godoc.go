package main

import "os/exec"

func (p *program) doc() error {
	docstr := p.FirstSelectedText()
	c := exec.Command("go", "doc", docstr)
	out, err := c.CombinedOutput()

	p.Name += "+Errors"
	p.Default = ""
	if err != nil {
		p.Text = err.Error() + "\n"
		p.Text += string(out)
	} else {
		p.Text = string(out)
	}
	return nil
}
