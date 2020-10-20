package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

// HandleFile calls a handler depending on the file name.
func (p *program) handleFile(file string, addr string) error {
	p.Name = p.TargetPath(file)
	p.Default = ""
	p.Replace = p.replace
	p.Clean = false
	tab := []struct {
		re   *regexp.Regexp
		def  string
		tags string
		typ  string
		fn   func(string, string) error
	}{
		{regexp.MustCompile(`_test\.go$`), "go", "-def -doc -Fmt -Test", "text/go", p.read},
		{regexp.MustCompile(`\.go$`), "go", "-def -doc -Fmt -Install", "text/go", p.read},
		{regexp.MustCompile(`\.txt$`), "", "Dos Write", "text", p.read},
		{regexp.MustCompile(`\.md$`), "", "markdown Write", "text", p.read},
		{regexp.MustCompile(`\.pdf$`), "", "", "", p.pdf},
		{regexp.MustCompile(`(?i)\.jpe?g$`), "", "", "", p.forward("image", nil)},
		{regexp.MustCompile(`(?i)\.png$`), "", "", "", p.forward("image", nil)},
		{regexp.MustCompile(`.zip$`), "", "", "", p.zip},
		{regexp.MustCompile(`.j$`), "ked", "-draw -run Write", "text/k", p.read},
		{regexp.MustCompile(`.k$`), "ked", "-draw -run Write", "text/k", p.read},
	}
	for _, t := range tab {
		if t.re.MatchString(file) {
			p.Default = t.def
			p.Tags = t.tags
			p.Type = t.typ
			return t.fn(file, addr)
		}
	}
	p.Tags = "Write"
	if err := p.read(file, addr); err != nil {
		return err
	}
	return nil
}

// Read is the default file handler.
// It returns the file content.
// It does not return large files.
func (p *program) read(file, addr string) error {
	if f, err := os.Open(file); err != nil {
		return err
	} else {
		defer f.Close()
		if fi, err := f.Stat(); err != nil {
			return err
		} else {
			if fi.IsDir() == true {
				return fmt.Errorf("%s: expected file, but it is a directory", file)
			}
			if size := fi.Size(); size > 1E6 {
				return fmt.Errorf("%s: file too large: %d", file, size)
			}
		}
		if all, err := ioutil.ReadAll(f); err != nil {
			return err
		} else {
			p.Text = string(all)
		}
	}

	p.SelectAddress(addr)
	p.Clean = true
	return nil
}

func (p *program) forward(fwdprog string, args []string) func(string, string) error {
	return func(file, addr string) error {
		return p.Forward(fwdprog, args)
	}
}
