// Read is an editor command which returns files from disk or directory listings.
//

// Read is called for a new window when the Name in the title bar is changed.
// In this case it has no arguments and behaves like that:
//
// 	Name
// 	/path/to/dir/       return directory listing
// 	/path/to/file       return file content
//
//
// If the current window is a directory listing or a command output, and read is executed with an argument,
// it opens the requested file or directory:
//
// 	Name            Argument
// 	/a/b/           c/ 	       return the directory listing /a/b/c/
// 	/a/b/+Errors    alpha.go    return file /a/b/alpha.go
//
// Adresses
// A file may have an address appended to it's name.
// The address follows a colon and has the following syntax:
// 	file:N             return file and select line N
// 	file:N:M           return file and select line N starting at character M
//
// By default each directory or file will be shown in a new window, or an already
// opened window for the target Name.
// This can be changed by passing the -r flag to read, in which case the source
// window will be replaced.
// To enable this, call "read -r /" manually from the Tag bar.
//
// It somehow serves the purpose of plan9's plumber.
// There should be more configurability. Currently each editor command which is bound to a file type
// is linked in ./file.go.
//
package main

//go:generate godocdown -output README.md

import (
	"fmt"
	"strings"

	"github.com/ktye/editor/cmd"
)

type program struct {
	cmd.Cmd
	replace string
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

	if len(p.Name) == 0 {
		return fmt.Errorf("Name is empty")
	} else if strings.HasPrefix(p.Name, "http") {
		return p.web()
	} else if p.Name[0] != '/' {
		return fmt.Errorf("Name must start with a '/'")
	}

	args := p.Args()
	if len(args) > 0 && args[0] == "-r" {
		p.replace = p.Name
		args = args[1:]
	}
	if len(args) == 0 {
		// A directory ends with a slash, but this may also be a file
		// with a regexp address part: file.txt:/alpha/
		if strings.HasPrefix(p.Name, "http") {
			return p.web()
		}
		if strings.IndexByte(p.Name, ':') == -1 && p.Name[len(p.Name)-1] == '/' {
			return p.handleDirectory(p.Directory())
		} else {
			path, addr := p.Path()
			return p.handleFile(path, addr)
		}
	} else {
		arg := args[0]
		if strings.HasPrefix(arg, "http") {
			p.Name = arg
			return p.web()
		}
		path, addr := p.ArgPath(arg)
		if strings.IndexByte(arg, ':') == -1 && arg[len(arg)-1] == '/' {
			return p.handleDirectory(path)
		} else {
			return p.handleFile(path, addr)
		}
	}
}
