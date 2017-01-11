function Menu(win) {
	this.win = win
	this.div = document.createElement("div") // menu div element
	this.div.className = "windowmenu"
	this.win.div.appendChild(this.div)
	this.button = {}

	this.AddEntry = function(title, callback) {
		var but = document.createElement('button')
		but.className = "menubutton"
		but.textContent = title
		var menu = this
		but.onclick = function() {
			if (callback != undefined)
				callback()
			menu.div.style.display = "none"
		}
		this.button[title] = but
		this.div.appendChild(but)
	}

	this.AddEntry("Cancel")
	this.AddEntry("Fullscreen", function() {win.ToggleFullscreen()})
	this.AddEntry("Toggle vi-mode", function(){win.editor.toggleVim()})
	this.AddEntry("Close", function(){win.Close()})
}
