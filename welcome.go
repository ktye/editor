package main

import (
	"fmt"
	"net/http"
)

// welcome serves the front-end and sets the "root" javascript variable.
func welcome(w http.ResponseWriter, r *http.Request) {

	// This is basically a static page, only the "root" variable must be set for
	// javascript. It is the path in the URL.
	// A template would be nicer, but might be overkill.

	w.Write([]byte(`
<!DOCTYPE html>
<html>
<head>
`))
	fmt.Fprintf(w, "<title>%s</title>\n", r.URL.Path)
	w.Write([]byte(`
<meta charset="UTF-8">


<link rel="stylesheet" href="/static/lib/blackboard.css">
<link rel="stylesheet" href="/static/lib/codemirror.css">
<link rel="stylesheet" href="/static/lib/dialog.css">
<link rel="stylesheet" href="/static/lib/fullscreen.css">
<link rel="stylesheet" href="/static/lib/simplescrollbars.css">
<link rel="stylesheet" href="/static/lib/matchesonscrollbar.css">
<script src="/static/lib/codemirror.js"></script>
<script src="/static/lib/dialog.js"></script>
<script src="/static/lib/search.js"></script>
<script src="/static/lib/searchcursor.js"></script>
<script src="/static/lib/matchbrackets.js"></script>
<script src="/static/lib/simplescrollbars.js"></script>
<!--script src="/static/lib/annotatescrollbar.js"></script-->
<!--script src="/static/lib/matchesonscrollbar.js"></script-->
<!--script src="/static/lib/fullscreen.js"></script-->
<script src="/static/lib/vim.js"></script>
<script src="/static/lib/mode/c.js"></script>
<script src="/static/lib/mode/go.js"></script>
<script src="/static/lib/mode/lua.js"></script>
<script src="/static/lib/mode/shell.js"></script>
<script src="/static/lib/mode/javascript.js"></script>
<script src="/static/lib/mode/html.js"></script>
<script src="/static/lib/mode/octave.js"></script>

<link rel="stylesheet" href="/static/editor.css">
`))
	fmt.Fprintf(w, "<script>var root = '%s'</script>\n", r.URL.Path)
	w.Write([]byte(`
<script src="/static/editor.js"></script>
<script src="/static/menu.js"></script>
<script src="/static/execute.js"></script>
<script src="/static/window.js"></script>
<script src="/static/column.js"></script>
<script src="/static/layout.js"></script>
<script src="/static/command.js"></script>
<script src="/static/main.js"></script>
</head>
<body onload="init()" oncontextmenu="return false" onresize="layout.Resize()">
<div class="maincontainer" id="maincontainer">
</div>
</body>
</html>
`))
}
