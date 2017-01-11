function Window(column, index) {
	this.column = column 	// pointer to the parent column the window is contained in
	this.index = index 	// index in parent's windows array
	this.weight = 1

	// window main div
	this.div = document.createElement("div") // div element
	this.div.className = "win"

	// allow being draged over and drop a winbox
	var win = this
	this.div.ondrop = function(ev) { 
		if (ev.dataTransfer.getData("column") == "")
			return
		var c = Number(ev.dataTransfer.getData("column"))
		var i = Number(ev.dataTransfer.getData("index"))
		var x0 = Number(ev.dataTransfer.getData("startx"))
		var y0 = Number(ev.dataTransfer.getData("starty"))
		win.column.layout.DropWindow(c, i, ev.clientX, ev.clientY, x0, y0)
	}
	this.div.ondragover = function(ev) { ev.preventDefault() }
	this.shiftdown = false
	this.div.onkeydown = function(ev) {
		if (ev.key == "Shift")
			win.shiftdown = true
	}
	this.div.onkeyup = function(ev) {
		if (ev.key == "Shift")
			win.shiftdown = false
	}

	if (this.column.windows.length == index)
		this.column.div.appendChild(this.div)
	else 
		this.column.div.insertBefore(this.div, this.column.windows[index].div)
	
	// window header div
	this.headerdiv = document.createElement("div")
	this.headerdiv.className = "winheader"
	this.div.appendChild(this.headerdiv)
	
	// window editor div
	this.editordiv = document.createElement("div")
	this.editordiv.className = "editordiv"
	this.div.appendChild(this.editordiv)

	// window drag box
	var win = this
	this.winbox = document.createElement('div')
	this.winbox.className = "winbox"
	this.winbox.oncontextmenu = function() {
		win.menu.div.style.display = "flex"
//		win.menu.button.Close.focus()
	}
	this.winbox.onclick = function() { win.Inc() }
	var winbox = this.winbox
	this.winbox.ondragstart = function(ev) {
		ev.dataTransfer.setData("startx", ev.clientX)
		ev.dataTransfer.setData("starty", ev.clientY)
		ev.dataTransfer.setData("column", win.column.index)
		ev.dataTransfer.setData("index", win.index)
	}
	this.winbox.draggable = true
	this.headerdiv.appendChild(this.winbox)

	// window title
	this.wintitle = document.createElement('input')
	this.wintitle.type = "text"
	this.wintitle.className = "wintitle"
	this.wintitle.onkeydown = function(ev){
		if (ev.keyCode == 13) {
			win.SetTitle(win.wintitle.value);
			if (win.wintitle.value.substr(-1) == "/")
				win.Execute("") // a directory must not have a command
			else
				win.Execute("Read") // a file needs to have command "Read", otherwise it is empty
			win.winbuttons.value = commander.GetHeader(win.wintitle.value)
		}
	}
	this.headerdiv.appendChild(this.wintitle)

	this.winbuttons = document.createElement('input')
	this.winbuttons.type = "text"
	this.winbuttons.className = "winbuttons"
	this.winbuttons.value = "Help"
	this.winbuttons.ondblclick = function(e) { win.ButtonExecute() }
	var winbuttons = this.winbuttons
	// right-click on selected text in winbutton-area executes this command
	this.winbuttons.oncontextmenu = function(ev) {
		var start = winbuttons.selectionStart
		var end = winbuttons.selectionEnd
		var str = winbuttons.value.substr(start, end-start)
		if (start == end)
			console.log('TODO how to get cursor position from INPUT field?')
		else
			win.Execute(str)
	}
	this.headerdiv.appendChild(this.winbuttons)

	this.menu = new Menu(win)
	this.menu.div.style.display = "none"


	// return opened file name
	this.GetFile = function() {
		return this.title // this could be changed to split at "+" and return the basename
	}

	// check if clientY is inside this window
	this.IsInside = function(y) {
		if ((y >= this.div.offsetTop) && (y <= this.div.offsetTop+this.div.offsetHeight))
			return true
		return false
	}

	// Mark window as modified
	this.MarkModified = function(modified) {
		if (modified == true)
			this.winbox.className = "winboxchanged"
		else
			this.winbox.className = "winbox"
	}

	// Check if window is marked as modified
	this.IsModified = function() {
		if (this.winbox.className == "winboxchanged")
			return true
		return false
	}

	// set title
	this.SetTitle = function(title) {
		if (title == undefined)
			title = "/"
		this.wintitle.value = title
		this.title = title
	}

	// toggle fullscreen
	this.ToggleFullscreen = function() {
		if (this.div.className == "win") {
			this.div.className = "fullscreenwin"
			this.column.div.style.left = "0px"
			this.column.div.style.width = "100%"
			this.div.style.top = "0px"
			this.div.style.left = "0px"
			this.div.style.width = window.innerWidth + "px"
			this.div.style.height = window.innerHeight + "px"
			this.editordiv.style.width = window.innerWidth + "px"
			this.editordiv.style.height = window.innerHeight - 22 + "px"
		} else {
			this.div.className = "win"
			this.column.layout.Resize()
		}
	}

	// close window
	this.Close = function() {
		this.column.windows.splice(this.index, 1)
		// recalculate index of remaining columns
		for (var i=0; i<this.column.windows.length; i++)
			this.column.windows[i].index = i;
		this.div.remove()
		if (this.column.windows.length == 0)
			this.column.Close()
		this.column.layout.Resize()
	}

	// decrease column width
	this.Dec = function() {
		if (this.weight == Infinity)
			this.weight = 3
		else
			this.weight--
		if (this.weight < 0)
			this.weight = 0
		this.div.style.flex = this.weight
	}

	this.ButtonExecute = function() {
		var wb = this.winbuttons
		var selectedText = wb.value.substring(wb.selectionStart, wb.selectionEnd).trim()

		// all other commands are send to the server as requests
		this.Execute(selectedText)
	}

	this.Execute = function(cmd) {
		if (cmd == undefined)
			cmd = ""
		Execute(this, cmd) // execute.js
	}
}


// increase window height
// window's weight may be 0 (closed), 1 (normal), 2 (double), max (all others closed)
Window.prototype.Inc = function() {
	if (this.relHeight == 0) {
		inc = 1 / this.column.windows.length
		this.relHeight = inc
	} else {
		inc = 2/3
		this.relHeight /= inc
	}


	// set to full column: close all other windows
	if (this.relHeight > 0.8) {
		this.relHeight = 1
		for (var i=0; i<this.column.windows.length; i++) {
			if (i != this.index)
				this.column.windows[i].relHeight = 0
		}
		this.column.layout.Resize()
		return
	}

	// double height, decrease all other windows heights
	for (var i=0; i<this.column.windows.length; i++) {
		if (i != this.index) {
			this.column.windows[i].relHeight *= inc
		}
	}
	this.column.layout.Resize()
}


