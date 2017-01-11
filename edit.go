package main

import (
	"bytes"
	"fmt"
	"github.com/ktye/editor/sam"
	"io"
	"io/ioutil"
)

// fileEdit uses the sam package to edit the posted content
func fileEdit(w *bytes.Buffer, cmd string, body io.ReadCloser, initDot string) (ok bool, addr string) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Fprintln(w, err)
		return false, ""
	}

	b, addr, err = sam.Edit(b, cmd, initDot)
	if err != nil {
		fmt.Fprintln(w, err)
		return false, ""
	}
	fmt.Fprint(w, string(b))

	return true, addr
}
