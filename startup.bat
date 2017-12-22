@echo off
REM
REM Set the environment and start to the editor server on windows.
REM

REM Set the path to use the msys64 environment for running external programs.
REM Other environments should work as well.
set PATH=c:/local/msys64/usr/bin;c:/local/go/bin;c:/local/msys64/home/elmar/bin

REM Set the gopath explicitly. This is only needed, to work on go projects.
REM Environment variables can also be changed interactively. See the env built-in.
set GOPATH=c:/Users/elmar

REM Start the editor server minimized
REM This is the binary built from ./editor/editor.
REM "cmd /c start /min c:\path\to\editor.exe"
CALL "cmd /c start /min c:\local\msys64\home\elmar\bin\editor.exe"

REM Open a new tab in the browser with for this root directory.
REM start "" "http://127.0.0.1:2017/path/to/root"
start "" "http://127.0.0.1:2017/c:/local/msys64/home/elmar/src/github.com/ktye"

