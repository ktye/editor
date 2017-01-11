package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ktye/editor/sam"
	"io"
	"os/exec"
	"path"
	"strconv"
)

// shell executes a shell command in the directory of the file.
func shell(w *bytes.Buffer, file string, cmd string, rootDir string) (ok bool) {
	com := exec.Command(*shellPath, "-c", cmd)
	com.Dir = rootDir + path.Dir(file)
	// TODO: write a custom CombinedOutput which prints the pid
	// to the console, so we can kill hanging processes
	out, err := com.CombinedOutput()
	if err != nil {
		fmt.Fprintln(w, string(out))
		fmt.Fprintln(w, err)
		return false
	}

	fmt.Fprint(w, string(out))
	return true
}

// shellFilter filteres the posted content with a shell command executes a shell command.
// If a single selection region is given in addr, only this part is replaced by the shell filter,
// otherwise the hole file is used.
// Multiple selections are not allowed.
func shellFilter(w *bytes.Buffer, file string, cmd string, body io.ReadCloser, addr string, rootDir string) (ok bool, pos string) {

	// Copy the posted body in a buffer
	var in bytes.Buffer
	_, err := io.Copy(&in, body)
	if err != nil {
		fmt.Fprintln(w, err)
		return false, ""
	}

	// Use sam to check the selected range.
	// var f sam.Samfile
	var initDots []sam.Address

	// f, initDots, err = sam.InitSamFile(in.Bytes(), "", addr)
	_, initDots, err = sam.InitSamFile(in.Bytes(), "", addr)
	if err != nil {
		fmt.Fprintln(w, err)
		return false, ""
	}

	var fullFile = true
	var origContent = in.Bytes()
	var from, to int
	if len(initDots) == 1 && initDots[0].Length() != 0 {
		fullFile = false
		from, to = initDots[0].GetRange()
		in = *bytes.NewBuffer(origContent[from:to])
	} else if len(initDots) > 1 {
		fmt.Fprintln(w, errors.New("multiple selections are not allowed for |-filter"))
		return false, ""
	}

	com := exec.Command(*shellPath, "-c", cmd)
	com.Dir = rootDir
	com.Stdin = &in

	var out bytes.Buffer
	com.Stdout = &out

	var errbuf bytes.Buffer
	com.Stderr = &errbuf

	err = com.Run()
	if err != nil {
		fmt.Fprintln(w, err)
		w.Write(errbuf.Bytes())
		return false, ""
	}

	if fullFile == true {
		w.Write(out.Bytes())
	} else {
		// Replace the filtered content at the selected position.
		newContent := out.Bytes()
		var b []byte
		b = append(b, origContent[0:from]...)
		b = append(b, newContent...)
		b = append(b, origContent[to:]...)

		// Calculate the new selection range.
		lineAddresses := sam.GetLineAddresses(b)

		line, ch := sam.LineAddr(from, true, lineAddresses, b)
		pos += strconv.Itoa(line) + ":" + strconv.Itoa(ch)
		line, ch = sam.LineAddr(from+len(newContent), false, lineAddresses, b)
		pos += "," + strconv.Itoa(line) + ":" + strconv.Itoa(ch)
		w.Write(b)
		return true, pos
	}
	return true, ""
}
