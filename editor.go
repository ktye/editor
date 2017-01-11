/*
Editor is a backend server which does file system operations and executes shell commands.

The frontend communicates with the backend with XMLHttpRequest with an URL in the form
	ADDR:PORT/path/to/root?file=/relpath/to/file&cmd1=command&cmd2=command2

To initialize the page, the browser requests the editor in the "root directory":
	ADDR:PORT/path/to/root
These requests are also answered by the editorHandler function.

Each session (open page in a browser) has an associated "root-directory".
On windows this could be
	127.0.0.1:1978/c:/Temp

All paths in the front-end are relative to this root directory, even if they
look like absolute paths (starting with /), same style on windows and unix.


The file can be a filename (starting with a slash and ending without a slash)
	file=/path/to/file
or a directory (starting and ending with a slash)
	file=/path/to/dir/
or a special command (starting without a slash)


The request method can be GET or POST.

For GET requests, on success the server writes a file and set's headers:
	content-type: to text/plain, image/png or text/html
	file: file name
	linenumber: current number in the file
	charposition: current character position on the line

The frontend checks if the file is already already open, updates it's content and sets
the cursor position.
If it does not yet exist, a new window is opened with it's content.

BASIC COMMANDS
	ftype	method	command	function
	log	GET		readLog()
		GET	Help	help()
	/dir/	GET		dir()
	/dir/	GET	Read	dir()
	/dir/	GET	/.../	dir()
	/dir/	GET	/...	fileRead()
	/dir/	GET	!...	shell()
	/dir/	POST	New	dirNew()
	/file	POST	diff	diffFile()
	/file	GET		fileRead()
	/file:.	GET	Read	fileRead()
	/file	GET	!...	shell()
	/file	POST	Read	fileRead()
	/file	POST	Write	fileWrite()
	/file	POST	!...	shell()
	/file	POST	|...	shellFilter()
	/file	POST	:...	fileEdit()
	http:.. GET	Read	webRead()

*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func editorHandler(w http.ResponseWriter, r *http.Request) {

	// allow only localhost connections
	remoteHost, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Fprintln(log, err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	if remoteHost != "127.0.0.1" { // TODO this might fail for IPv6, may be updated to "::1" ?
		fmt.Fprintln(log, "request from", remoteHost, "denied")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	method := r.Method

	file := r.FormValue("file")

	// If the file is empty, this is an initial request from the user.
	// Respond with the welcome page, which builds the front-end.
	if file == "" {
		// http.ServeFile(w, r, staticDir + "/editor.html")
		fmt.Fprintf(log, "new editor on %s\n", r.URL.Path)
		welcome(w, r)
		return
	}

	rootDir := r.URL.Path
	rootDir = windowsRoot(rootDir)

	addr := r.FormValue("selections")
	var cmds []string
	for i := 1; ; i++ {
		s := r.FormValue("cmd" + strconv.Itoa(i))
		if s == "" {
			break
		}
		cmds = append(cmds, s)
	}

	fmt.Fprintf(log, "method=%s root=%s file=%s cmds=%s\n", method, rootDir, file, strings.Join(cmds, "; "))

	// set response header
	w.Header().Set("Content-type", "text/plain") // this is the default type // TODO add html/png....

	// output is buffered because the error code is not known until somewhere undefined
	// but it must be written first to w
	var b bytes.Buffer

	if len(file) == 0 {
		file = "/"
	}

	var ok bool
	if method == "GET" {
		if file == "log" {
			ok = readLog(&b)
		} else if strings.HasPrefix(file, "http://") {
			w.Header().Set("Content-type", "text/html")
			ok = webRead(&b, file)
		} else if len(cmds) == 1 && cmds[0] == "Help" {
			ok = help(&b)
		} else if isDir(file) {
			ok = dirGet(&b, file, cmds, w, rootDir)
		} else if file[0] == '!' { // file name starts with !: execute shell command.
			ok = shell(&b, "", file[1:], rootDir)
		} else {
			ok = fileGet(&b, file, cmds, w, rootDir)
		}
	} else if method == "POST" {
		if isDir(file) {
			ok = dirPost(&b, file, cmds, r.Body, w, addr, rootDir)
		} else {
			ok = filePost(&b, file, cmds, r.Body, w, addr, rootDir)
		}
	} else {
		// TODO how to signal error?
		ok = false
		fmt.Println(&b, "unknown http method")
	}

	// errors are indicated with StatusTeapot
	if ok == false {
		w.WriteHeader(http.StatusTeapot)
	}

	// copy from the buffer to the response writer
	io.Copy(w, &b) // errors are ignored. What should I do with them?
}

// GET requests on a directory
func dirGet(w *bytes.Buffer, file string, cmds []string, rw http.ResponseWriter, rootDir string) (ok bool) {
	for i := 0; i < len(cmds); i++ {
		c := cmds[i]

		// /dir/	GET		dir()
		if c == "" {
			if dir(w, file, rootDir) == false {
				return false
			}

			// /dir/	GET	Read	dir()
		} else if c == "Read" {
			if dir(w, file, rootDir) == false {
				return false
			}

			// /dir/	GET	/.../	dir()
		} else if isDir(c) {
			if dir(w, c, rootDir) == false {
				return false
			}

			// /dir/	GET	/...	fileRead()
		} else if c[0] == '/' {
			var pos string
			var ok bool
			if ok, pos = fileRead(w, c, rootDir); ok == false {
				return false
			}
			rw.Header().Set("Selections", pos)

			// /dir/	GET	!...	shell()
		} else if c[0] == '!' {
			if shell(w, file, c[1:], rootDir) == false {
				return false
			}

		} else {
			fmt.Fprintln(w, "dirGet: unknown command:", c)
			return false
		}
	}
	return true
}

// GET requests on a file
func fileGet(w *bytes.Buffer, file string, cmds []string, rw http.ResponseWriter, rootDir string) (ok bool) {
	for i := 0; i < len(cmds); i++ {
		c := cmds[i]

		if (c == "Read") || len(c) == 0 {
			var ok bool
			var pos string
			if ok, pos = fileRead(w, file, rootDir); ok == false {
				return false
			}
			rw.Header().Set("Selections", pos)
		} else if (len(c) > 1) && c[0] == '!' {
			if shell(w, file, c[1:], rootDir) == false {
				return false
			}
		} else {
			fmt.Fprintln(w, "unknown command for fileGet:", c)
			return false
		}
	}
	return true
}

// POST request on a directory.
func dirPost(w *bytes.Buffer, file string, cmds []string, body io.ReadCloser, rw http.ResponseWriter, addr string, rootDir string) (ok bool) {
	for i := 0; i < len(cmds); i++ {
		c := cmds[i]

		if c == "New" {
			if dirNew(w, file, body, rootDir) == false {
				return false
			}
		} else if len(c) > 1 && c[0] == ':' {
			var ok bool
			var pos string
			if ok, pos = fileEdit(w, c[1:], body, addr); ok == false {
				return false
			}
			rw.Header().Set("Selections", pos)
		} else {
			fmt.Fprintln(w, "unknown command for dirPost:", c)
			return false
		}
	}
	return true
}

// POST request on a file.
func filePost(w *bytes.Buffer, file string, cmds []string, body io.ReadCloser, rw http.ResponseWriter, addr string, rootDir string) (ok bool) {
	for i := 0; i < len(cmds); i++ {
		c := cmds[i]

		if c == "Read" {
			if ok, _ := fileRead(w, file, rootDir); ok == false {
				return false
			}
		} else if c == "Write" {
			if fileWrite(w, file, body, rootDir) == false {
				return false
			}
		} else if (len(c) > 1) && (c[0] == '!') {
			if shell(w, file, c[1:], rootDir) == false {
				return false
			}
		} else if (len(c) > 1) && (c[0] == '|') {
			var ok bool
			var pos string
			if ok, pos = shellFilter(w, file, c[1:], body, addr, rootDir); ok == false {
				return false
			}
			rw.Header().Set("Selections", pos)
		} else if (len(c) > 1) && (c[0] == ':') {
			var ok bool
			var pos string
			if ok, pos = fileEdit(w, c[1:], body, addr); ok == false {
				return false
			}
			rw.Header().Set("Selections", pos)
		} else if c == "diff" {
			return diffFile(w, file, body, rootDir)
		} else {
			fmt.Fprintln(w, "filePost: unknown command:", c)
			return false
		}
	}
	return true
}

func readLog(w *bytes.Buffer) (ok bool) {
	// dont Read() from logBuf (or io.Copy) because that would clear the log buffer
	fmt.Fprint(w, logBuf.String())
	return true
}

// isDir checks if a file name should be treated as a directory.
// A directory is a filename, which starts and end with a slash,
// with the exception if is contains a colon (sam pattern: file.go:/^func main/).
func isDir(file string) bool {
	idx := strings.Index(file, ":")
	if idx == -1 {
		if len(file) > 0 && file[0] == '/' && file[len(file)-1] == '/' {
			return true
		}
	}
	return false
}

// windowsRoot magically detects if the root directory is a windows path and strips the leading slash.
// If the root dir is only "/" on windows, for requests: HOST:PORT without any additional path,
// this will not work.
func windowsRoot(rootDir string) string {
	if len(rootDir) > 2 && rootDir[2] == ':' { // this checks for /c:/...
		return rootDir[1:]
	}
	return rootDir
}
