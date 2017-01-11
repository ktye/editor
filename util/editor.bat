@echo off
REM
REM Shortcut to start the editor server.
REM Create a link on the taskbar with this batch file as a target.
REM It also starts the godoc server.
REM
REM This will not work as is. Configure as needed.
REM

REM Set ROOT to the project the editor should work on.
REM All files in the editor will be relative to this directory.
set ROOT=c:/local/msys64/home/elmar/src/github.com/ktye

REM Set the GOPATH if you are working on a go project.
set GOPATH=c:/local/msys64/home/elmar

REM This should point to bin/ directory of the go installation to find go.
REM It should also contain the folder for goimports (golang.org/x/tools/cmd/goimports which should be installed).
set PATH=c:/local/go/bin;%GOPATH%/bin

REM Set EDITORDIR to the editor source.
REM It must point to the parent of the directory static/.
set EDITORDIR=c:/local/msys64/home/elmar/src/github.com/ktye/editor

REM Set EDITORBIN to the editor.exe path.
set EDITORBIN=c:/local/msys64/home/elmar/bin/editor.exe

REM Set BUSYBOX to to sh.exe (e.g: from https://frippery.org/busybox/)
set BUSYBOX=c:/local/busybox/sh.exe

REM This starts godoc.
CALL "cmd /c start /min c:\local\go\bin\godoc.exe -http=:6060"

REM Open the frontend in a webserver or new tab
start http://127.0.0.1:1978/%ROOT%

REM Start the editor.
%EDITORBIN% -install=%EDITORDIR% -shell=%BUSYBOX%
