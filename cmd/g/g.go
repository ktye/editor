// G is an editor command which runs a grep like command.
//
// It prints files and lines of the matching regular expression (like grep -n)
// and is recursive (like grep -r) starting on the window's directory.
//
// The regular expression is passed as an argument or the FirstSelectedText is used.
// If a second argument is given, it is interpreted as a regular expression matching
// against the file name, e.g. "\.go$" to match only go files.
package main

//go:generate godocdown -output README.md

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/ktye/editor/cmd"
)

type program struct {
	cmd.Cmd
}

func main() {
	var p program
	if err := p.Run(); err != nil {
		p.Fatal(err)
	}
	p.Exit()
}

func (p *program) Run() error {
	if err := p.Parse(); err != nil {
		return err
	}

	if err := os.Chdir(p.Directory()); err != nil {
		return err
	}

	args := p.Args()
	find := ""
	filter := ""
	switch len(args) {
	case 0:
		find = p.FirstSelectedText()
	case 1:
		find = args[0]
	case 2:
		find = args[0]
		filter = args[1]
	default:
		return fmt.Errorf("g is called with too many arguments")
	}
	if find == "" {
		return fmt.Errorf("regular expression is empty")
	}

	re, err := regexp.Compile(find)
	if err != nil {
		return err
	}

	match := func(string) bool { return true }
	if filter != "" {
		re, err := regexp.Compile(filter)
		if err != nil {
			return err
		}
		match = re.MatchString
	}

	var out bytes.Buffer
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() == false {
			return nil
		}
		if info.IsDir() == false {
			if match(info.Name()) {
				grep(&out, path, re)
			}
			return nil
		}
		return nil
	})
	if err != nil {
		return err
	}

	p.Name = p.TargetPath(p.Directory()) + "/+Errors"
	p.Default = "read"
	p.Text = string(out.Bytes())
	return nil
}

func grep(w io.Writer, path string, re *regexp.Regexp) {
	r, err := os.Open(path)
	if err != nil {
		return
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	line := 0
	for scanner.Scan() {
		line++
		if t := scanner.Text(); re.MatchString(scanner.Text()) == true {
			fmt.Fprintf(w, "%s:%d: %s\n", filepath.ToSlash(path), line, t)
		}
	}
}
