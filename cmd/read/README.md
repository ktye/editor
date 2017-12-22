# read
--
Read is called for a new window when the Name in the title bar is changed. In
this case it has no arguments and behaves like that:

    Name
    /path/to/dir/       return directory listing
    /path/to/file       return file content

If the current window is a directory listing or a command output, and read is
executed with an argument, it opens the requested file or directory:

    Name            Argument
    /a/b/           c/ 	       return the directory listing /a/b/c/
    /a/b/+Errors    alpha.go    return file /a/b/alpha.go

Adresses A file may have an address appended to it's name. The address follows a
colon and has the following syntax:

    file:N             return file and select line N
    file:N:M           return file and select line N starting at character M

By default each directory or file will be shown in a new window, or an already
opened window for the target Name. This can be changed by passing the -r flag to
read, in which case the source window will be replaced. To enable this, call
"read -r /" manually from the Tag bar.

It somehow serves the purpose of plan9's plumber. There should be more
configurability. Currently each editor command which is bound to a file type is
linked in ./file.go.
