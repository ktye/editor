// Editor is a code editor with a web frontend and a go server backend.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/tabwriter"
)

var (
	staticDir string       // path to static content (INSTALLDIR/static or CWD/static)
	shellPath *string      // path of a shell command the server exectues for !... and |... escapes
	logBuf    bytes.Buffer // logBuf must be accessible by readLog(), never Read() from logBuf
	log       io.Writer    // log messages are multiplexed to stdout and to a log buffer for the client
)

func main() {
	var err error
	httpAddr := flag.String("http", "127.0.0.1:1978", "HTTP service address (e.g. 127.0.0.1:1978)")
	shellPath = flag.String("shell", "/bin/sh", "path to shell for shell escapes.") // on windows: "c:/path/to/busybox/sh.exe"
	installPath := flag.String("install", "", "path to install directory (parent of static/)")
	flag.Parse()

	if *installPath == "" {
		*installPath, err = os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	staticDir = *installPath + "/static"

	log = io.MultiWriter(os.Stdout, &logBuf)

	w := new(tabwriter.Writer)
	w.Init(log, 0, 8, 0, '\t', 0)
	fmt.Fprintf(w, "install dir\t%s\n", *installPath)
	fmt.Fprintf(w, "shell\t%s\n", *shellPath)
	fmt.Fprintf(w, "address\t%s\n", *httpAddr)
	fmt.Fprintf(w, "gopath\t%s\n", os.Getenv("GOPATH"))
	w.Flush()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	http.HandleFunc("/", editorHandler)

	err = http.ListenAndServe(*httpAddr, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
