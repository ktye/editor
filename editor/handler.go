package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/ktye/editor/cmd"
)

// Request is a json encoded post request from the front end.
type Request struct {
	cmd.Cmd
	Command string
}

// Handler is the single request handler for the http server.
// It serves the frontend as a single-page application for every GET request,
// and handles editor commands which are sent by POST requests.
// POST requests sent by the editor frontend are of type Request encoded in json.
//
// The response has follows the same protocol as the request, excluding the
// Command string.
//
// Commands fall into 3 categories and are matched in the order:
//	- built-in commands from the Command field
//	- built-in commands from the Name field
//	- external commands (standalone executables, following the protocol).
func handler(w http.ResponseWriter, r *http.Request) {

	// Send a new front-end for all GET Requests
	// and set the client's root variable to the path given in the URL
	// e.g. 127.0.0.1:2017/d:/path/to/root .
	if method := r.Method; method == "GET" {
		w.Write([]byte(strings.Replace(ClientApplication, "__ROOT__", r.URL.Path, -1)))
		return
	} else if method != "POST" {
		log.Print("unknown request method: " + method)
		http.Error(w, "unknown request method", http.StatusExpectationFailed)
		return
	}

	// Handle editor commands which are sent as POST requests.
	var req Request
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		log.Printf("cannot decode request: %s", err)
		http.Error(w, "cannot decode request", http.StatusBadRequest)
		return
	}
	req.Root = cleanRoot(req.Root)
	log.Printf("Name: %s Command: %q", req.Name, req.Command)

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(req.Cmd); err != nil {
		log.Printf("could not encode command: %s", err)
		http.Error(w, "could not encode command", http.StatusExpectationFailed)
		return
	}

	// Execute built-in commands.
	if isBuiltin(req.Command) {
		if err := execBuiltin(w, req.Command, &buf); err != nil {
			sendError(w, err)
		}
		return
	}

	// Execute built-in commands from Name field.
	if isBuiltin(req.Name) {
		if err := execBuiltin(w, req.Name, &buf); err != nil {
			sendError(w, err)
		}
		return
	}

	// Prefix leading ! | > < on commands with pipe.
	if len(req.Command) > 0 {
		switch c := req.Command[0]; c {
		case '!', '|', '>', '<':
			req.Command = "pipe " + string(c) + " " + req.Command[1:]
		}
	}

	// Execute external command.
	// If the command start with a hyphen, append it to the request's default command.
	if len(req.Command) > 0 && req.Command[0] == '-' {
		req.Command = req.Default + " " + req.Command
	}
	if err := execute(w, req.Command, &buf); err != nil {
		log.Println("error", err)
		sendError(w, err)
	}
}

func sendError(w http.ResponseWriter, err error) {
	c := cmd.Cmd{
		Name:  "+Errors",
		Tags:  "",
		Clean: false,
		Text:  err.Error(),
	}
	enc := json.NewEncoder(w)
	enc.Encode(c)
}
