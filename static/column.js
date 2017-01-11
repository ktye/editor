function Column(layout) {
	this.layout = layout			// reference to parent
	this.index = layout.columns.length	// index of this in layout.columns
	this.div = document.createElement("div")
	this.div.className = "column"
	this.relWidth = 1
	this.left = 0
	document.getElementById('maincontainer').appendChild(this.div)
	this.windows = []

	// add a new window after at index
	this.AddWindow = function(index, title) {
		var win = new Window(this, index)
		this.windows.splice(index, 0, win)
		if (title != undefined)
			win.SetTitle(title)
		for (var i=0; i<this.windows.length; i++) {
			this.windows[i].index = i;
		}
		win.relHeight = 1 / this.windows.length
		for (var i=0; i<this.windows.length; i++) {
			if (i != index) {
				this.windows[i].relHeight *= (this.windows.length - 1)/this.windows.length
			}
		}
		this.layout.Resize()
		return win
	}

	// remove window from column
	this.RemoveWindow = function(index) {
		var win = this.windows[index]	// save window
		win.div.remove()		// remove window div from DOM
		this.windows.splice(index, 1)	// remove window reference from colum
		for (var i=0; i<this.windows.length; i++)	// update all window's indexes
			this.windows[i].index = i
		if (this.windows.length == 0)
			this.Close()
		return win
	}

	// insert window before nextwindow (can be undefined)
	this.InsertWindow = function(win, nextwin) {
		if (nextwin == undefined) {
			this.windows.push(win)
			this.div.appendChild(win.div)
		} else {
			this.windows.splice(nextwin.index, 0, win)
			this.div.insertBefore(win.div, nextwin.div)
		}
		win.column = this
		for (var i=0; i<this.windows.length; i++)
			this.windows[i].index = i
	}

	// normalize
	this.Normalize = function() {
		for (var i=0; i<this.windows.length; i++) {
			this.windows[i].weight = 1
			this.windows[i].div.style.flex = 1
		}
	}

	// check if clientX is inside this column
	this.IsInside = function(x) {
		if ((x >= this.div.offsetLeft) && (x <= this.div.offsetLeft+this.div.offsetWidth))
			return true
		return false
	}


	// close column
	this.Close = function() {
		this.layout.columns.splice(this.index, 1)
		// recalculate index of remaining columns
		for (var i=0; i<layout.columns; i++)
			layout.columns[i].index = i;
		this.div.remove()
	}

	// increase column width
	this.Inc = function() {
		this.relWidth *= 3/2
		this.layout.Resize()
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
}
