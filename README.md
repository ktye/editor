# Editor

An editor inspired by acme.

![screenshot](editor.png)

## Design
- Editor server (./editor/...)
	- go implementation with the executable editor(.exe)
	- serves the front-end
	- handles editor requests on complete window content
	- executes commands or handles built-ins
	- does not follow the state of windows or their layout
- Front-end: (./editor/html/...)
	- mainly javascript served by the editor server on initial requests.
	- does the layout of windows
	- uses codemirror library for the actual editing
- Editor commands (./cmd/...)
	- each editor command is an executable program and may be written in any language
	- is has to follow the protocol (./editor/cmd/cmd.go: Cmd) encoded in json.
	- new commands can be added or replaced, which does not require changes to the server or front-end
	- even read or write are their own commands
	
## Compared to acme
- runs on windows, uses the browser as frontend
- supports images and html
- their is no file system access to the editor state
- current state and layout is only known to the front-end
- lacks the top row, windows are added by pulling down the top window and columns are added by moving a window to the right by more than 50%
