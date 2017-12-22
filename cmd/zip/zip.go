// Zip is an editor command which allows to browse zip files.
//
// Zip is called from read and exposes the contents of a zip file
// as a directory structure.
// The name will have the form `/path/to/file.zip/path/to/file`.
package main

//go:generate godocdown -output README.md

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ktye/editor/cmd"
)

type program struct {
	cmd.Cmd
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

	if args := p.Args(); len(args) > 0 {
		p.Name += args[0]
	}

	zipName, rel := p.splitZipPath()
	if zipName == "" {
		return fmt.Errorf("zip could not find zip path in Name")
	}

	z := cmd.Cmd{
		Root: p.Root,
		Name: zipName,
	}
	path, _ := z.Path()

	if r, err := zip.OpenReader(path); err != nil {
		return err
	} else {
		defer r.Close()

		if rel[len(rel)-1] == '/' {
			return p.zipListDir(r, rel)
		} else {
			if f, err := zipFindFile(r, rel); err != nil {
				return err
			} else {
				return p.zipReadFile(f)
			}
		}
	}
}

func (p *program) zipListDir(r *zip.ReadCloser, relpath string) error {
	var out bytes.Buffer
	exists := map[string]bool{}
	for _, f := range r.File {
		name := "/" + f.Name
		if strings.HasPrefix(name, relpath) {
			name = name[len(relpath):]
			if idx := strings.IndexByte(name, '/'); idx != -1 {
				name = name[:idx+1]
			}
			if exists[name] == false && name != "" {
				fmt.Fprintf(&out, "%s\n", name)
			}
			exists[name] = true
		}
	}
	p.Text = string(out.Bytes())
	p.Default = "zip"
	p.Clean = true
	return nil
}

func (p *program) zipReadFile(f *zip.File) error {
	if r, err := f.Open(); err != nil {
		return err
	} else {
		defer r.Close()
		if b, err := ioutil.ReadAll(r); err != nil {
			return err
		} else {
			p.Text = string(b)
			p.Default = ""
			p.Clean = true
			return nil
		}
	}
}

func zipFindFile(r *zip.ReadCloser, relpath string) (*zip.File, error) {
	for _, f := range r.File {
		name := "/" + f.Name
		if name == relpath {
			return f, nil
		}
	}
	return nil, fmt.Errorf("%s: file does not exist in zip", relpath)
}

// splitZipPath splits the Name field after the zip file.
// It returns the zip file name realtive to Root and the
// path within the zip file.
// If the path inside the zip file is empty, "/" is returned.
// /path/to/file.zip/path/to/a.txt is split to
// /path/to/file.zip and /path/to/a.txt .
// On error both paths are empty.
func (p *program) splitZipPath() (string, string) {
	idx := strings.Index(p.Name, ".zip")
	if idx == -1 {
		return "", ""
	}
	zipName := p.Name[:idx+4]
	relpath := p.Name[idx+4:]
	if relpath == "" {
		relpath = "/"
	}
	return zipName, relpath
}
