package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ktye/editor/cmd"
)

// Execute splits the commandline of the requested editor command
// and calls an external program.
func execute(w io.Writer, commandline string, req io.Reader) error {
	argv, err := cmd.SplitQuoted(commandline)
	if err != nil {
		return err
	}

	// We treat a pipe command specially.
	// It will be splitted by the pipe binary.
	if strings.HasPrefix(commandline, "pipe ") {
		argv = []string{"pipe", commandline[5:]}
	}

	if len(argv) < 1 {
		return fmt.Errorf("request contains no command")
	}

	// Get installation directory of editor binary.
	// All subcommands must be in the same directory.
	var installDir string
	progname := os.Args[0]
	if p, err := filepath.Abs(progname); err != nil {
		return fmt.Errorf("cannot get editor directory")
	} else {
		installDir = filepath.Dir(p)
	}

	var buf bytes.Buffer
	var errbuf bytes.Buffer
	argv[0] = filepath.Join(installDir, argv[0])
	ctx, cancel := context.WithCancel(context.Background())
	c := exec.CommandContext(ctx, argv[0], argv[1:]...)
	c.Stdin = req
	c.Stdout = &buf
	c.Stderr = &errbuf
	if err := c.Start(); err != nil {
		return err
	}
	pid := c.Process.Pid
	ProcessList.Add(pid, argv, cancel)

	err = c.Wait()
	ProcessList.Remove(pid)
	io.Copy(w, &buf)

	// Write stderr of commands to the console.
	if errbuf.Len() > 0 {
		if err != nil {
			errmsg, _ := ioutil.ReadAll(&errbuf)
			err = fmt.Errorf("%s\n%s\n", err.Error(), string(errmsg))
		} else {
			io.Copy(os.Stdout, &errbuf)
		}
	}
	return err
}

// ProcessList keeps a list of all executed processes running in the background.
var ProcessList procList

type procList struct {
	sync.Mutex
	list map[int]proc
}

type proc struct {
	argv   []string
	cancel func()
}

func (p *procList) Add(pid int, args []string, cancel func()) {
	p.Lock()
	if p.list == nil {
		p.list = make(map[int]proc)
	}
	p.list[pid] = proc{args, cancel}
	p.Unlock()
}

func (p *procList) GetCancel(pid int) func() {
	p.Lock()
	defer p.Unlock()
	if proc, ok := p.list[pid]; ok {
		return proc.cancel
	}
	return nil
}

func (p *procList) Remove(pid int) {
	p.Lock()
	delete(p.list, pid)
	p.Unlock()
}
