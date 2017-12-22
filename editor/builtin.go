package main

import (
	"fmt"
	"io"

	"github.com/ktye/editor/cmd"
)

var builtins map[string]Builtin = map[string]Builtin{}

type Builtin func(io.Writer, []string, io.Reader) error

func isBuiltin(command string) bool {
	argv, err := cmd.SplitQuoted(command)
	if err != nil || len(argv) < 1 {
		return false
	}
	if _, ok := builtins[argv[0]]; ok {
		return true
	}
	return false
}

func execBuiltin(w io.Writer, command string, r io.Reader) error {
	argv, err := cmd.SplitQuoted(command)
	if err != nil {
		return err
	}

	if len(argv) < 1 {
		return fmt.Errorf("request contains no command")
	}

	if b, ok := builtins[argv[0]]; ok == false {
		return fmt.Errorf("builtin command %s does not exist", argv[0])
	} else {
		return b(w, argv[1:], r)
	}
}
