package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// dir reponds with a list of all files and directories.
func dir(w *bytes.Buffer, file string, rootDir string) (ok bool) {
	fullpath := rootDir + file
	files, err := ioutil.ReadDir(fullpath)
	if err != nil {
		fmt.Fprintln(w, err)
		return false
	}
	for _, f := range files {
		filename := f.Name()
		if f.IsDir() {
			filename += "/"
		}
		fmt.Fprintln(w, filename)
	}
	return true
}

// dirNew reads directory from post request and checks if new files or directories should be created.
func dirNew(w *bytes.Buffer, file string, body io.ReadCloser, rootDir string) (ok bool) {
	var f *os.File
	var err error
	f, err = os.Open(rootDir + file)
	if err != nil {
		fmt.Fprintln(w, err)
		return false
	}
	var dirnames []string
	dirnames, err = f.Readdirnames(0)
	if err != nil {
		fmt.Fprintln(w, err)
		return false
	}
	root := rootDir + file
	r := bufio.NewReader(body)
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Fprintln(w, err)
			return false
		}
		// remove trailing newline, may not exist for last line
		if len(line) > 0 && line[len(line)-1] == '\n' {
			line = line[0 : len(line)-1]
		}
		// return on empty lines (often the last)
		if len(line) == 0 {
			break
		}

		// last char is '/', it's a directory
		if line[len(line)-1] == '/' {
			newdir := line[0 : len(line)-1]
			if err = createFile(newdir, true, dirnames, root); err != nil {
				fmt.Fprintln(w, err)
				return false
			}
		} else { // it's a file
			newfile := line
			if err = createFile(newfile, false, dirnames, root); err != nil {
				fmt.Fprintln(w, err)
				return false
			}
		}
	}

	// Respond with the updated dir.
	return dir(w, file, rootDir)
}

// createFile tries to create a new file or directory if they don't exist yet.
// newfile does not contain the tailing / for directories
// createFile is not a good name, because it implies that it creates files, it does so however only when necessary.
func createFile(newfile string, isDir bool, dirnames []string, root string) (err error) {
	for i := 0; i < len(dirnames); i++ {
		if dirnames[i] == newfile {
			return nil
		}
	}

	if strings.IndexByte(newfile, '/') != -1 {
		return errors.New("filename contains a slash") // slash in between is not allowed
	}

	if isDir {
		return os.Mkdir(root+newfile, os.ModeDir)
	} else {
		var f *os.File
		var err error
		if f, err = os.Create(root + newfile); err != nil {
			return err
		}
		return f.Close()
	}
}
