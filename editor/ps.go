package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/ktye/editor/cmd"
)

func init() {
	builtins["ps"] = ps
	builtins["kill"] = kill
}

// Ps is a builtin command which prints a list of all child processes.
func ps(w io.Writer, args []string, r io.Reader) error {
	ProcessList.Lock()
	defer ProcessList.Unlock()

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "child processes (double-click pid to kill)\n\n")
	for pid, proc := range ProcessList.list {
		fmt.Fprintf(&buf, "[%d] %q\n", pid, proc.argv)
	}

	res := cmd.Cmd{
		Name:    "ps",
		Default: "kill",
		Clean:   true,
		Text:    string(buf.Bytes()),
	}
	enc := json.NewEncoder(w)
	return enc.Encode(res)
}

// Kill terminates a running external process.
// It is invoced form a ps window by clicking on a pid.
// It kills only processes that have been started before.
// It returns a new ps window.
func kill(w io.Writer, args []string, r io.Reader) error {
	if len(args) != 1 {
		return fmt.Errorf("kill needs 1 argument, got %d", len(args))
	}
	arg := args[0]

	// The pid could be encoded by brackets.
	if len(arg) > 0 && arg[0] == '[' {
		arg = arg[1:]
	}
	if len(arg) > 0 && arg[len(arg)-1] == ']' {
		arg = arg[:len(arg)-1]
	}
	if len(arg) < 1 {
		return fmt.Errorf("kill: argument is empty")
	}

	pid, err := strconv.Atoi(arg)
	if err != nil {
		return fmt.Errorf("kill: no valid pid: %s", arg)
	}

	if cancel := ProcessList.GetCancel(pid); cancel != nil {
		cancel()
		time.Sleep(500 * time.Millisecond)
		return ps(w, []string{}, r)
	}
	return fmt.Errorf("kill: %d: process does not exist", pid)
}
