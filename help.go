package main

import (
	"bytes"
)

// help sends the help
// fileRead returns the file content.
func help(w *bytes.Buffer) (ok bool) {
	w.WriteString(`EDITOR HELP
===========
double click the line number to jump to the section
@/^WINDOW LAYOUT
@/^SEARCH
@/^REPLACE
@/^SHELL COMMANDS
@/^SHELL FILTERS
@/^LOG FILE
===========

WINDOW LAYOUT
SEARCH
REPLACE
SHELL COMMANDS
SHELL FILTERS
LOG FILE
	`)
	return true
}
