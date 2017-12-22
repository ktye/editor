# g
--
G is an editor command which runs a grep like command.

It prints files and lines of the matching regular expression (like grep -n) and
is recursive (like grep -r) starting on the window's directory.

The regular expression is passed as an argument or the FirstSelectedText is
used. If a second argument is given, it is interpreted as a regular expression
matching against the file name, e.g. "\.go$" to match only go files.
