# Editor

A code editor with a go backend inspired by acme.

![screenshot](editor.png)

# Features / Usage
- The editor works even on windows.
- All files use *pseudo* absolute paths. The prefix is given in the URL, which could be the GOPATH or the root of a project.
- Start a new window by dragging the top one down (use the blue drag box)
- Move the window to another column by drgging it left by more than 1/2
- The drag box of the top window can be used to move the complete column left or right
- See util/editor.bat for a startup script with explanations on Windows

# Limitations
- It should be used for small files only (normal code editing). Each command takes a full loop to the server and back.
- Many more...

# Status
The editor is not good.
It is just an initial version suffering a complete overhaul.
Too much is happening in javascript which should be moved to the back end.
It has a lot of limitations and things should be done differently.
The problem is that the current version *works* for the daily go routine.
