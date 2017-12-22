# git
--
Git is an editor command which wraps git.

Called with no argument, it executes "git status" and presents the output in a
new window. The window has it's default command set to "git -add".

Double-clicking any files in the status window will run git add on these files.
Files or blocks of files may also be selected and executing "-add" on the tag
bar will add these files. Git -add strips prefixes such as "modified:" and any
whitespace.

To commit changes, type a commit message in the window, select it and execute
"-commit" on the tag bar.
