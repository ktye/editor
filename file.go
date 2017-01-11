package main

import (
	"bytes"
	"fmt"
	"github.com/ktye/editor/sam"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// fileRead returns the file content.
// the file string may contain an edit command (e.g. address range)
func fileRead(w *bytes.Buffer, file string, rootDir string) (ok bool, addr string) {

	// default return selection range is empty
	addr = ""

	// sam command
	var cmd = ""

	// get address rage part, the part after the first ":"
	idx := strings.Index(file, ":")
	if idx != -1 {
		cmd = file[idx+1:]
		file = file[0:idx]
	}

	// open disc file
	f, err := os.Open(rootDir + file)
	if err != nil {
		fmt.Fprintln(w, err)
		return false, ""
	}
	defer f.Close()
	d, err1 := f.Stat()
	if err1 != nil {
		fmt.Fprintln(w, err)
		return false, ""
	}

	// if it's a directory, return it's content instead
	if d.IsDir() {
		if file[len(file)-1] != '/' {
			file = file + "/"
		}
		return dir(w, file, rootDir), ""
	}

	// don't process large files, it's too slow and mostly a mistake
	if d.Size() > 256*1024 {
		fmt.Fprintln(w, "file is too large:", d.Size())
		return false, ""
	}

	// read file content
	b, err2 := ioutil.ReadAll(f)
	if err2 != nil {
		fmt.Println(w, err2)
		return false, ""
	}

	// edit the content with sam
	if cmd != "" {
		b, addr, err = sam.Edit(b, cmd, "") // on the read request, the initDot is the whole file
		fmt.Println("sam cmd:", cmd, "addr", addr)
	}

	w.Write(b) // ignore errors, what could we do about it anyway?

	return true, addr
}

// fileWrite reads file content from POST request and writes file to disk.
func fileWrite(w *bytes.Buffer, file string, body io.ReadCloser, rootDir string) (ok bool) {

	out, err := os.Create(rootDir + file)
	if err != nil {
		fmt.Fprintf(w, "unable to create file: %s", rootDir+file)
		return false
	}
	defer out.Close()

	_, err = io.Copy(out, body)
	if err != nil {
		fmt.Fprintln(w, err)
		return false
	}

	return true
}
