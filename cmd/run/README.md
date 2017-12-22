# run
--
Run is an editor command which executes a program associated with the extension
of the current window Name.

### Associations

    Name       command line
    file.awk   awk -f file.awk
    file.go    go run file.go
    file.sed   sed -f file.sed
    file.sh    sh file.sh


### Flags

All arguments before the last are appended to the command line.


### InOutput

If run is called without arguments, it's stdin is empty and it's output is
written to an +Errors window. If run is called with a file arguemnt (e.g.
`file.txt`), the input is read from disk from a file build by ArgPath from the
directory of the script and the given file argument. It's output is written to a
window with the name of the TargetPath of the file argument. If the argument is
prefixed with `<`, as in `run <file.txt`, only the input is read from the file,
and the output is written to an `+Errors` window. If the argument is prefixed
with `>`, as in `run >file.txt`, the input is empty.


### Usage

Edit the script in it's own window (e.g. `/path/to/script.sed`) and the file to
be modified in another window (e.g. `/path/to/file.txt`). Now execute "run
file.txt" in the script window.
