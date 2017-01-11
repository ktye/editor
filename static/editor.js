// attach an editor to an element: new Editor(element)
function Editor(win) {
	this.win = win
	this.editor = CodeMirror(win.editordiv, {
//		vimMode: true,
		lineNumbers: true,
		matchBrackets: true,
//		showCursorWhenSelecing: true,
		tabSize: 8,
		indentUnit: 8,
		indentWithTabs: true,
		smartIndent: true,
		lineWrapping: false,
		scrollbarStyle: "simple",
	});
	this.selectedLines = []

	// postContent returns the editor content for POST requests.
	// It is either the complete body or the selected parts.
	this.postContent = function() {
		// TODO check for selections and return them instead
		return this.editor.getValue()
	}

	/*
	this.At = function(cmd) {
		if (position.match("^([0-9]+),([0-9]+)$") != null) {
			var pos = position.match("^([0-9]+),([0-9]+)$") 
			this.editor.setCursor({line:Number(pos[1])-1, ch:Number(pos[2])-1})
		} else if (position.match("^([0-9]+)$") != null) {
			this.editor.setCursor({line:Number(position)-1, ch:0})
		} else if ((position.length > 0)&&(position[0]=="/")) {
			this.Search(position.substr(1))
		} else {
			console.log("editor.At: unknown position", position)
		}

		this.editor.focus()
	}
	// search next occurance of regex
	this.Search = function(s) {
		var re = new RegExp(s)
		var sc = this.editor.getSearchCursor(re, this.editor.getCursor())
		if (sc.findNext()) {
			this.editor.setSelection(sc.from(), sc.to())
		} else {
			// unselect and goto start
			this.editor.setSelection({line:0,ch:0},{line:0,ch:0}, {scroll:false})
		}
	}
	*/

	// SetMode sets the syntax mode.
	// If the mode-string is undefined, it tries to auto detect it from the window.title
	// using the 'commander'.
	this.SetMode = function(mode) {
		if (mode != undefined) {
			this.editor.setOption('mode',mode)
			return
		}
		var title = this.win.title
		for (var key in commander.languagemodes) {
			if (title.match(key) != null) {
				this.editor.setOption('mode',commander.languagemodes[key])
				return
			}
		}
		this.editor.setOption('mode','null')
	}

	// set current selections 
	// as a list of selections SEL1;SEL2;...
	// with a selection as a position FROM,TO
	// with the position being LINE[:CHARPOS] as 1-based integers
	this.SetSelections = function(s) {
		var getPos = function(p) {
			var v = p.split(":")
			return {line:Number(v[0])-1, ch:Number(v[1])-1}
		}
		var selections = s.split(";")
		var sets = []
		for (var i=0; i<selections.length; i++) {
			var ranges = selections[0].split(",")
			var from = getPos(ranges[0])
			var to = getPos(ranges[1])
			sets.push({anchor:from, head:to})
		}
		this.editor.setSelections(sets)
		this.editor.focus()
	}

	// toggle vim mode
	this.toggleVim = function() {
		if (this.editor == undefined)
			return
		var vim = this.editor.getOption('vimMode')
		this.editor.setOption('vimMode', !vim)
	}

	// get current selected text as string
	this.GetSelectedText = function() {
		// multiple selections are concatenated and separated by newline
		return this.editor.getSelection("\n")
	}

	// get current selections encoded as a string, see SetSelection for the syntax
	this.GetSelections = function(s) {
		var getPos = function(p) {
			return String(p.line+1) + ":" + String(p.ch+1)
		}
		var getRange = function(r) {
			return getPos(r.anchor) + "," + getPos(r.head)
		}
		var sets = this.editor.listSelections()
		var selections = []
		for (var i=0; i<sets.length; i++) {
			var range = getRange(sets[i])
			selections.push(range)
		}
		return selections.join(";")
	}

	// MarkRectangle transforms a multi-line selection into a a set of rectangular selections.
	// Lines shorter than the selected range will be filled.
	this.MarkRectangle = function(s) {
		var sets = this.editor.listSelections()
		if (sets.length != 1)
			return // there should be only one selection
		sel = sets[0]
		if (sel.head.ch <= sel.anchor.ch)
			return
		var newsels = []
		for (var i=sel.anchor.line; i <= sel.head.line; i++) {
			this.FillLine(i, sel.head.ch)
			newsels.push({anchor:{line:i, ch:sel.anchor.ch}, head:{line:i, ch:sel.head.ch}})
		}
		this.editor.setSelections(newsels)
	}
	// MarkColumn transforms a single selection in to a set of rectangular selections over all lines.
	// Lines shorter than the selected range will be filled.
	this.MarkColumn = function(s) {
		var sets = this.editor.listSelections()
		if (sets.length != 1)
			return // there should be only one selection
		sel = sets[0]
		if (sel.head.ch <= sel.anchor.ch)
			return
		var newsels = []
		var lineCount = this.editor.lineCount()
		for (var i=0; i < lineCount; i++) {
			if ((i == lineCount-1) && (this.editor.getLine(i) == "")) {
				continue // this is the trailing newline, ignore this line
			}
			this.FillLine(i, sel.head.ch)
			newsels.push({anchor:{line:i, ch:sel.anchor.ch}, head:{line:i, ch:sel.head.ch}})
		}
		this.editor.setSelections(newsels)
	}

	// FillLine filles the numbered line with spaces upto position maxch.
	this.FillLine = function(linenum, maxch) {
		l = this.editor.getLine(linenum)
		if (l.length < maxch) {
			this.editor.replaceRange( " ".repeat(maxch-l.length), {line:linenum, ch:l.length})
		}
	}



	// double-click on a line: execute this line
	this.editor.on('dblclick', function(ev) {
		var c = ev.doc.getCursor()
		var s = ev.doc.getRange({line:c.line,ch:0},{line:c.line,ch:Infinity})
		win.Execute(s)
	})

	// set the border color to red if the editor is active
	var edt = this
	this.editor.on('focus', function() {
		edt.win.div.style.border = '1px solid red'
	})
	this.editor.on('blur', function() {
		edt.win.div.style.border = '1px solid black'
	})

	// mouse-click callback: right click: execute the selection
	this.editor.on('mousedown', function(cm, evt) {
		if (evt.button == 2) {
			var s = cm.getSelections()
			if ((s.length == 1) && s[0] != "") {
				s = s[0]
				win.Execute(s)
			}
		}
	})



// this does not work, it is triggerd for every key stroke and messes shift-somekey!
//
//	// If there is one selection and it is over multiple lines,
//	// mark all lines full.
//	// Select full columns (over all lines, by shift-click) on a different location on the same line as the cursor.
//	// Block-select by shift-click on a different line as the cursor.
//	this.editor.on('beforeSelectionChange', function(cm, a) {
//		if (edt.win.shiftdown) { // this is handled by mousedown
//			if (a.ranges.length == 1) {
//				// same line: select columns
//				if (a.ranges[0].anchor.line == a.ranges[0].head.line) {
//					var fromcol = a.ranges[0].anchor.ch
//					var tocol = a.ranges[0].head.ch
//					var fromline = 0
//					var toline = cm.lineCount()-1
//				} else { // block-select
//					var fromline = Math.min(a.ranges[0].anchor.line, a.ranges[0].head.line)
//					var toline = Math.max(a.ranges[0].anchor.line, a.ranges[0].head.line)
//					var fromcol = Math.min(a.ranges[0].anchor.ch , a.ranges[0].head.ch)
//					var tocol = Math.max(a.ranges[0].anchor.ch, a.ranges[0].head.ch)
//				}
//				a.ranges = []
//				for (var i=fromline; i<=toline; i++) {
//					a.ranges.push({anchor:{line:i,ch:fromcol},head:{line:i,ch:tocol}})
//				}
//				a.update(a.ranges)
//			}
//			return
//		}
//		if (a.ranges.length == 1) {
//			// selection is empty, reset selectedLines
//			if ((a.ranges[0].anchor.line == a.ranges[0].head.line) && (a.ranges[0].anchor.ch == a.ranges[0].head.ch) ) {
//				edt.selectedLines = []
//				return
//			}
//			var r = a.ranges[0]
//			if (a.ranges[0].anchor.line > a.ranges[0].head.line) {
//				a.ranges[0].anchor.ch = Infinity
//				a.ranges[0].head.ch = 0
//				edt.selectedLines = []
//				for (var i=a.ranges[0].head.line; i<=a.ranges[0].anchor.line; i++)
//					edt.selectedLines.push(i)
//			} else if (a.ranges[0].anchor.line < a.ranges[0].head.line) {
//				a.ranges[0].anchor.ch = 0
//				a.ranges[0].head.ch = Infinity
//				edt.selectedLines = []
//				for (var i=a.ranges[0].anchor.line; i<=a.ranges[0].head.line; i++)
//					edt.selectedLines.push(i)
//			} else {
//				return
//			}
//			a.update([a.ranges[0]])
//		}
//	})

	// select lines by clicking line numbers
	this.editor.on('gutterClick', function(cm, n, gutter, ev) {
		// If shift key is pressed, select all lines from the
		// first selectedLine (or from start if empty)
		// to the current line.
		//
		// Otherwise add clicked line to the selectedLines if not jet there,
		// or remove it.
		if (ev.shiftKey) {
			var start
			var end = n
			if (edt.selectedLines.length == 0)
				start = 0
			else
				start = edt.selectedLines[0]
			if (end < start) {
				end = start
				start = n
			}
			edt.selectedLines = []
			for (var i=start; i<=end; i++)
				edt.selectedLines.push(i)
		} else {
			var idx = edt.selectedLines.indexOf(n)
			if (idx == -1)
				edt.selectedLines.push(n)
			else
				edt.selectedLines.splice(idx,1)
		}
		edt.selectedLines.sort()

		var a = []
		for (var i=0; i<edt.selectedLines.length; i++) {
			var line = edt.selectedLines[i]
			a.push({anchor:{line:line,ch:0},head:{line:line,ch:Infinity}})
		}
		if (a.length > 0)
			cm.setSelections(a)
		else
			cm.setSelection(cm.getCursor(), cm.getCursor())

/*
		var s = cm.listSelections()
		for (var i=0; i<s.length; i++) {
			if (s[i].anchor.line == n) {
				s.splice(i,1)
				console.log('selection must be removed on line',n)
				console.log(s)
				cm.setSelections(s)
				return
			}
		}
		cm.addSelection({line:n,ch:0}, {line:n,ch:Infinity})
	*/
	})

	this.editor.on('dragstart', function(cm, ev) {
//		ev.preventDefault()
	})
	this.editor.on('drop', function(cm, ev) {
//		ev.preventDefault()
	})

	var win = this.win
	this.editor.on('change', function(cm, ev) {
		win.MarkModified(true)
	})

}
