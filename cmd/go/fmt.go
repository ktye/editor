package main

import (
	"bytes"
	"os/exec"
)

func (p *program) fmt() error {
	if err := p.Forward("Write", nil); err != nil {
		return err
	}

	if err, errtxt := goImports(p.Base()); err != nil {
		return err
	} else if errtxt != "" {
		p.Name += "+Errors"
		p.Default = ""
		p.Text = errtxt
		p.Clean = false
		return nil
	}
	return p.Forward("read", []string{p.Base()})
}

func goImports(filename string) (error, string) {
	var stderr bytes.Buffer
	cmd := exec.Command("goimports", "-w", filename)
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return err, ""
	}

	if err := cmd.Wait(); err != nil {
		return nil, string(stderr.Bytes())
	}
	return nil, ""
}
