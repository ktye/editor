package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func (p *program) def() error {
	file := p.Base()
	if p.Selections == nil {
		return fmt.Errorf("no text is selected")
	}
	br := p.Selections[0].ByteRange(p.Text)
	offset := br[0]

	argv := []string{"godef", "-o", strconv.Itoa(offset), "-f", file, "-i"}
	c := exec.Command(argv[0], argv[1:]...)
	c.Stdin = strings.NewReader(p.Text)
	out, err := c.CombinedOutput()

	p.Name += "+Errors"
	p.Default = ""
	p.Text = ""
	if err != nil {
		p.Text += err.Error() + "\n"
		p.Text += string(out)
		return nil
	}

	// Filter absolute path names in the output.
	var b bytes.Buffer
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		t := scanner.Text()
		if rel, err := filepath.Rel(p.Directory(), t); err == nil {
			b.Write([]byte(rel))
		} else {
			b.Write([]byte(t))
		}
		b.Write([]byte("\n"))
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	p.Text = string(b.Bytes())
	return nil
}
