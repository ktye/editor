package main

import (
	"bytes"
	"os/exec"
)

func (p *program) install() error {
	if err := p.Forward("Write", nil); err != nil {
		return err
	}

	if err, errtxt := goInstall(); err != nil {
		return err
	} else if errtxt != "" {
		p.Name += "+Errors"
		p.Default = ""
		p.Text = errtxt
		p.Clean = false
		return nil
	}

	if err := p.Forward("read", nil); err != nil {
		return err
	}
	return nil
}

func goInstall() (error, string) {
	var stderr bytes.Buffer
	cmd := exec.Command("go", "install")
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return err, ""
	}

	if err := cmd.Wait(); err != nil {
		return nil, string(stderr.Bytes())
	}
	return nil, ""
}
