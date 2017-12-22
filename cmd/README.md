# cmd
--
    import "github.com/ktye/editor/cmd"

Package cmd contains helper function to implement an editor command.

An editor command is a standalone program which communicates with the editor
server by stdin and stdout.

## Usage

#### func  CompareTestResults

```go
func CompareTestResults(got Cmd, expected Cmd) error
```
CompareTestResults compares a command output in the reader with the expected
Cmd.

#### func  SplitQuoted

```go
func SplitQuoted(line string) ([]string, error)
```
splitQuoted splits a line at whitespace, except if it is quoted. Whitespace at
the beginning or end of a string is trimmed.

Example:

    	`alpha  beta ` => {"alpha", "beta"}
    	`alpha " beta "` => {"alpha", " beta "}
    	`alpha "beta \" gamma"` => {"alpha", `beta " gamma`}
         `alpha|beta` => {"alpha","|","beta"},

#### func  SplitQuotedPipe

```go
func SplitQuotedPipe(line string) []string
```
SplitQuotedPipe splits the input line at '|' characters, if they are not quoted.

#### type Cmd

```go
type Cmd struct {
	Root       string     // Directory root.
	Name       string     // Window ID.
	Replace    string     // Window ID to replace.
	Tags       string     // New window tags.
	Default    string     // Default command for executed text in the body.
	Selections RuneRanges // Current selections.
	Type       string     // Content type "text" (default), "text/go", or "html".
	Clean      bool       // Mark buffer as clean.
	Text       string     // File content.
}
```

Cmd is a common type which is embedded by editor command types.

#### func  NewTestRequest

```go
func NewTestRequest() *Cmd
```
NewTestRequest returns a Cmd for testing commands with the Root field filled.

#### func (*Cmd) ArgPath

```go
func (c *Cmd) ArgPath(relPath string) (string, string)
```
ArgPath returns the full path build from the Directory and a given relative path
which may contain backwards slashes on windows. It also returns the optional
address part after the first ':' in the relPath.

#### func (*Cmd) Args

```go
func (c *Cmd) Args() []string
```
Args returns the command line arguments. The program name is stripped. They are
available after a call to parse.

#### func (*Cmd) Base

```go
func (c *Cmd) Base() string
```
Base returns the base name of the file from the Name field. It cuts everything
before the last '/' and the first ':'.

#### func (*Cmd) ByteRangeToRuneRange

```go
func (cmd *Cmd) ByteRangeToRuneRange(from, to int) RuneRange
```
ByteRangeToRuneRange converts a byte range to a rune range.

#### func (*Cmd) CombinedSelectedText

```go
func (c *Cmd) CombinedSelectedText() string
```
CombinedSelectedText joins the selected text strings by newline.

#### func (*Cmd) Directory

```go
func (c *Cmd) Directory() string
```
Directory returns the full path of the directory combining the Root and Name
fields. The path has the format of the os, e.g. "c:\path\to\file.txt" on
windows.

#### func (*Cmd) Encode

```go
func (c *Cmd) Encode(w io.Writer)
```
Encode encodes the command as json to w.

#### func (*Cmd) Exit

```go
func (c *Cmd) Exit()
```
Exit writes the header with no error including the file content and terminates
the program.

#### func (*Cmd) Fatal

```go
func (c *Cmd) Fatal(err error)
```
Fatal writes the header containing the error to stdout and terminates the
program. The function does nothing if the error is nil.

#### func (*Cmd) FirstSelectedText

```go
func (c *Cmd) FirstSelectedText() string
```
FirstSelectedText return the text of the first selection.

#### func (*Cmd) Forward

```go
func (c *Cmd) Forward(prog string, args []string) error
```
Forward forwards the request to a new process and decodes the response.

#### func (*Cmd) Parse

```go
func (c *Cmd) Parse() error
```
Parse reads the header and data from stdin. It terminates the program with an
appropriate error message in case of an error.

#### func (*Cmd) Path

```go
func (c *Cmd) Path() (string, string)
```
Path return the full path of the file combining the Root and Name fields and the
optional address after the first ':' in the Name field.

#### func (*Cmd) ReplaceSelections

```go
func (c *Cmd) ReplaceSelections(repl string)
```
ReplaceSelections replaces all selections with repl. It selects the replaced
strings and sets their selection.

#### func (*Cmd) SelectAddress

```go
func (c *Cmd) SelectAddress(addr string)
```
SelectAddress sets c.Selection from the given address. The address has the form:
N select line N N:M select line N starting at character M /reg/ select the first
match of the regular expression /reg/g select all matches of the regular
expression If the last character is a colon, it is removed.

#### func (*Cmd) SetIO

```go
func (c *Cmd) SetIO(in io.Reader, out io.Writer, args []string)
```
SetIO should be used only for testing commands.

#### func (*Cmd) TargetPath

```go
func (c *Cmd) TargetPath(fullPath string) string
```
TargetPath substracts the root directory from the given full path. It converts
all slashes to forward slashes.

#### func (*Cmd) Write

```go
func (c *Cmd) Write() error
```

#### type RuneRange

```go
type RuneRange [2]int
```

RuneRange defines the start and end address of a selection. The address is the
rune index starting at 0. An empty selection defines a cursor position and has
equal start and end address. The length of the selection is end-start, the same
as slice indexing.

#### func (RuneRange) ByteRange

```go
func (rr RuneRange) ByteRange(text string) (b [2]int)
```
ByteRange converts a RuneRange to a byte range.

#### func (RuneRange) Slice

```go
func (rr RuneRange) Slice(text string) string
```
Slice returns the text defined by the RuneRange.

#### type RuneRanges

```go
type RuneRanges []RuneRange
```


#### func (RuneRanges) String

```go
func (rr RuneRanges) String() string
```

#### func (RuneRanges) Total

```go
func (rr RuneRanges) Total() int
```
Total returns the number of selected runs.
