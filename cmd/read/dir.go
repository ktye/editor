package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// handleDirectory returns a directory listing.
func (p *program) handleDirectory(dir string) error {
	p.Name = p.TargetPath(dir)
	if len(p.Name) > 0 && strings.HasSuffix(p.Name, "/") == false {
		p.Name = p.Name + "/"
	}
	p.Type = "text"
	p.Replace = p.replace
	if p.replace == "" {
		p.Default = "read"
	} else {
		p.Default = "read -r"
	}
	p.Tags = ""
	p.Selections = nil
	p.Clean = false
	if f, err := os.Open(dir); err != nil {
		return err
	} else {
		defer f.Close()

		if fi, err := f.Readdir(-1); err != nil {
			return err
		} else {
			var buf bytes.Buffer
			if p.Name != "/" {
				fmt.Fprintf(&buf, "../\n")
			}
			for _, info := range fi {
				if info.IsDir() {
					fmt.Fprintf(&buf, "%s/\n", info.Name())
				} else {
					fmt.Fprintf(&buf, "%s\n", info.Name())
				}
			}
			p.Text = string(buf.Bytes())
		}
	}
	p.Clean = true
	return nil
}
