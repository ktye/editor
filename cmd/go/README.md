# go
--
Go is an editor command which bundles some go commands.

Go is set as the default command for files with .go extension.

### Subcommands

    -Fmt runs "goimports -w" on the current file.
    -Install runs "go install" in the window's directory.
    -Test runs "go test" in the window's directory.
    -def runs "godef" (github.com/rogpeppe/godef) using the current selection.
    -doc runs "go doc" with the current selection.

Subcommands -Fmt, -Test and -Install also write the file to disk.


### Requirements

These external commands must be on the path: go, goimports, godef.
