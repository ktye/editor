package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func (p *program) gitAdd(args []string) error {
	if len(args) != 0 {
		if err := gitadd(args); err != nil {
			return err
		}
	} else {
		var files []string
		text := p.CombinedSelectedText()
		lines := strings.Split(text, "\n")
		prefixes := []string{"modified:"}
		for _, line := range lines {
			file := strings.TrimSpace(line)
			for _, prefix := range prefixes {
				if strings.HasPrefix(file, prefix) {
					file = strings.TrimSpace(file[len(prefix):])
				}
			}
			if file != "" {
				files = append(files, file)
			}
		}
		if len(files) > 0 {
			if err := gitadd(files); err != nil {
				return err
			}
		}
	}
	return p.gitStatus()
}

func gitadd(files []string) error {
	args := append([]string{"add"}, files...)
	c := exec.Command("git", args...)
	if out, err := c.CombinedOutput(); err != nil {
		return fmt.Errorf("%s\n%s", err, out)
	} else {
		return nil
	}
}
