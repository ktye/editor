package main

import "os/exec"

func (p *program) test(args []string) error {
	if err := p.Forward("Write", nil); err != nil {
		return err
	}

	if err, errtxt := goTest(args); err != nil && errtxt == "" {
		return err
	} else if errtxt != "" {
		p.Name += "+Errors"
		p.Default = ""
		p.Text = errtxt
		p.Clean = false
		return nil
	}
	p.Clean = true
	return nil
}

func goTest(args []string) (error, string) {
	args = append([]string{"test"}, args...)
	cmd := exec.Command("go", args...)
	b, err := cmd.CombinedOutput()
	return err, string(b)
}
