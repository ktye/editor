// Execute a window command.
// This looks up the command from profile and sends a server request.
function Execute(win, cmd) {
	var file = win.wintitle.value
	if (ExecuteBuiltin(win, cmd))
		return

	var selectedText = ""
	if (win.editor != undefined)
		selectedText = win.editor.GetSelectedText()
	var C = commander.Lookup(file, cmd, selectedText)
	// console.log(C)
	var req = new XMLHttpRequest()

	// request's callback function
	req.onreadystatechange = function() {
		var DONE = this.DONE || 4
		if (this.readyState == DONE) {
			// look for open window with this filename
			// var file = req.getResponseHeader("file")
			var newfile = C.newfile
			// errors are indicated by the server with status 418
			if (this.status == 418)
				newfile += "+Errors" // create a new window with the +Errors suffix
			var newwin = win.column.layout.FindWindow(newfile)
			var responseText = undefined

			/* don't silently overwrite modified windows.
			// this does not work. It prevents normal "Writes" from working.
			if (newwin != undefined && newwin.IsModified()) {
				newfile += "+Errors"
				newwin = win.column.layout.FindWindow(newfile)
				responseText = "window is already open and marked as modified.\nWrite or Close it before executing the command again."
			}
			*/
			var selections = undefined
			if (newwin == undefined) {
				newwin = win.column.AddWindow(win.index+1)
				newwin.SetTitle(newfile)
				newwin.winbuttons.value = commander.GetHeader(newfile)
			} else if (newwin.editor != undefined)
				selections = newwin.editor.GetSelections()

			var ctype = req.getResponseHeader("Content-type")
			if (ctype.length > 0) {
				ctype = ctype.split(";")[0]
			}

			// selections can be overwritten by the response
			var at = req.getResponseHeader("Selections")
			if ((at != null) && (at.length > 0))
				selections = at

			if (ctype == "text/html") {
				newwin.editordiv.innerHTML = req.response
			} else if (ctype == "text/plain") {
				if (newwin.editor == undefined) {
					newwin.editor = new Editor(newwin)
				}
				if (responseText == undefined)
					responseText = req.response
				newwin.editor.editor.setValue(responseText)
				newwin.editor.SetMode(undefined)
				if (selections != undefined)
					newwin.editor.SetSelections(selections)
				newwin.editor.editor.focus()
				newwin.MarkModified(false)
			} else if (ctype == "image/png") {
				console.log("received image/png TODO")
			} else {
				console.log("received unknown content-type: ", ctype)
			}
		}
	}

	// special case: original requested file starts with ':'
	// this is the sam-command window for the requested file open in another window.
	// But we need to post the content and the selections of the other window, not the command window.
	if ((file[0] == ':')&&(file.length > 1)) {
		win = win.column.layout.FindWindow(file.substring(1))
	}

	// build and send the request
	// The request URL has the form fields: file, cmd1, cmd2, ...
	// If it is a post request, it contains also the editor content in the body
	var url = root + "?file=" + encodeURIComponent(C.file) 
	for (var i=0; i<C.cmds.length; i++) {
		url += "&cmd" + String(i+1) + "=" + encodeURIComponent(C.cmds[i])
	}
	if (win.editor != undefined)
		url += "&selections=" + encodeURIComponent((win.editor.GetSelections()))
	req.open(C.method, url, true)
	if (C.method == "POST") {
		req.send(win.editor.postContent())
	}
	else
		req.send(null)
}

// Execute built-in commands
function ExecuteBuiltin(win, cmd) {
	
	//var file = win.wintitle.value
	//if (file.length > 0 && file[0] == ":") {	
	//}
	
	if ((cmd != undefined) && (cmd.length > 1)) {
		if (cmd == "Rect") {
			win.editor.MarkRectangle()
			win.editor.editor.focus()
			return true
		} else if (cmd == "Col") {
			win.editor.MarkColumn()
			win.editor.editor.focus()
			return true
		}
	}
	return false
}
