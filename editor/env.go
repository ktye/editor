package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ktye/editor/cmd"
)

func init() {
	builtins["env"] = env
}

// Env is a builtin command which prints or sets the environment for external commands.
// Env by itself prints the environmen to a window with the Name "env".
//
// To change or set a variable, edit it and double click the line.
// This will call "env" with the line "var=value" as an argument.
// The new environment is returned.
func env(w io.Writer, args []string, r io.Reader) error {
	if len(args) > 1 {
		return fmt.Errorf("too many arguments")
	}
	if len(args) == 1 {
		if idx := strings.Index(args[0], "="); idx == -1 {
			return fmt.Errorf("env: syntax must be key=value")
		} else {
			key := args[0][:idx]
			value := args[0][idx+1:]
			if err := os.Setenv(key, value); err != nil {
				return err
			}
		}
	}
	res := cmd.Cmd{
		Name:    "env",
		Default: "env",
		Clean:   true,
		Text:    strings.Join(os.Environ(), "\n"),
	}
	enc := json.NewEncoder(w)
	return enc.Encode(res)
}
