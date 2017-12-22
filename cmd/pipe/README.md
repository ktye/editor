# pipe
--
Pipe is an editor command which executes a pipe line of external commands and
handles in and output depending on the first character of the argument.

    First arugment  input      output

    !               none       Name+Errors
    |               Selection  Selection
    <               none       Selection
    >               Selection  Name+Errors

If the Selection is empty on input the complete file is used.

Pipe expects a single argument which will be splitted following the quoting
rules. If the arguments contains a pipe, it connects multiple programs.

### Example

    pipe "!grep -n alpha file | wc -l"
