function Commander() {

	// each new window who's title matche one of these fields is given the associated initial header fields.
	this.header = {
		"/$": 		"New",		// default directory header
		"^/.*[^/]$":	"Read Write",	// default file header
		"^/.*\\.go$":	"Run Fmt Build Vet def doc Write",	// go file header
		"^/.*_test\\.go$": "Fmt Vet def doc Test Write",		// go test files
		"!":	"",		// all command output has now headers (unmatch previous pattern)
		"^log$": "Read",	// log window header
		"^/.*\\+Errors$": "",	// error window header
		"^http://": "",		// embedded web pages
	}

	// the editor mode is switched to the given mode, if the window title matches one of these lines
	this.languagemodes = {
		"\.c$":		"clike",
		"\.go$":	"go",
		"\.lua$":	"lua",
		"\.sh$":	"shell",
		"\.js$":	"javascript",
		"\.m$":		"octave",
		// "\.html$":	"htmlmixed", // this produces: TypeError: htmlMode.startState is not a function
	}

	// command matching table for predefined window names and commands.
	// everything which does not match any of these rows gets the default context.
	// the last match wins, so more specific entries must come later.
	this.table = [
		// filepattern: the current window title is matched against this filepattern, if it does match, try to go on.
		// commandpattern: the executed command is matched against this command pattern, if it does match, try to go on.
		// method: the http-request method used to contact the editor-server.
		// infilepattern: this infile pattern is expanded and send to the editor-server as the "file" parameter, if undefined, the current window title is used.
		// commandpatternvector: all elements in this array are expanded and concatenated with ";" and send to the editor-server as the "cmd" parameter
		// newfilepattern: this pattern is expanded and used as the window title for the new window or the one being replaced with.

		//filepattern,	commandpattern,	method,	infilepattern,	commandpatternvector, 	newfilepattern


		// godoc
		// [".",		"^godoc$",	"GET",	"http://localhost:6060/pkg",	["Read"], "godoc"],

		// cmp and diff
		[".",		"^cmp$",	"POST",	undefined,	["cmp"],		undefined],
		[".",		"^diff$",	"POST",	undefined,	["diff"],		undefined],

		// directory commands
		["/$",		"^(.*)$",	"GET",	undefined,	["{{FILE}}{{CMD}}"],	"{{FILE}}{{CMD}}" ], // this is for line clicks in a directory listing, command will be /clickedfile
		["/$",		"^New$",	"POST",	undefined,	["New"],		undefined],

		// help
		[".",		"^Help$",	"GET",	undefined,	["Help"], "Help"],

		// default file commands
		["^/.*[^/]$",	"^Write$",	"POST",	undefined,	["Write","Read"], 	undefined], // the default primitive "Write" on the server produces no output, that's why the file must be "Read" afterwards
		["^/.*[^/]$",	"^Read$",	"GET",	undefined,	["Read"],		undefined],

		// go file commands
		["^/.*\\.go$",	"^Run$",	"POST",	undefined,	["Write", "!go run {{LOCALFILE}}" ], "{{FILE}}!run"], // open result in a new window
		["^/.*\\.go$",	"^Fmt$",	"POST",	undefined,	["Write", "!goimports -w {{LOCALFILE}} 1>&-", "Read"], undefined], // open result in the same window, suppress go fmt output on success
		["^/.*\\.go$",	"^Build$",	"POST",	undefined,	["Write", "!go build", "Read" ], undefined], // on success a build produces no output and the final "Read" puts the window back to normal state  on error {{FILE}}+Errors are shown as usual
		["^/.*\\.go$",	"^Vet$",	"POST",	undefined,	["Write", "!go tool vet {{LOCALFILE}}", "Read" ], undefined], // run go vet on the single file
		["^/.*\\.go$",	"^Test$",	"POST",	undefined,	["Write", "!go test"], "{{FILE}}!test"], // run go test in the directory
		["^/.*\\.go$",	"^def$", 	"GET",	undefined,	["!godef -A -f {{LOCALFILE}} {{TEXT}}"], "{{FILE}}!godef {{TEXT}}"], // run godef doc with the argument of the selected text, similar to "doc"
		["^/.*\\.go$",	"^doc$", 	"GET",	undefined,	["!go doc {{TEXT}}"], "!go doc {{TEXT}}"], // run go doc with the argument of the selected text (e.g. mark string.Split in the file and click the "doc" command to look it up)

		// error windows
		["^/.*\\+Errors$", "(.+)",				"GET",	"{{DIR}}{{\\1}}", ["Read"], "{{DIR}}{{\\1}}"], // command is file name with optional address ranges
		
		// godef output
		["!godef ", 	"(.+)",					"GET",	"{{DIR}}{{\\1}}", ["Read"], "{{DIR}}{{\\1}}"], // command is file name with optional address ranges

		// shell commands
		[".",		"^!",		"GET",	"{{DIR}}",	["{{CMD}}"],	"{{DIR}}{{CMD}}"], // shell commands (commands starting with "!") can be executed on directories or files, and open a new window with the directory name + ! + shell command
		[".",		"^\\|",		"POST",	undefined,	["{{CMD}}"],	undefined], // shell filter commands (commands starting with "|") are executed on files, they filter the hole file, or a selection. Multiple selections are not allowed. 

		// search and edit commands (sam)
		[".",		"^:",		"POST",	undefined,	["{{CMD}}"],		undefined], // any command starting with a : is an editing command for this window
//		["^:(.*)",	".*",		"GET", "{{F\\1}}:{{CMD}}",	["Read"],		"{{F\\1}}"], // an editing window starts with :/file and contains editing commands. They are sent to the server and apply for the same file (which should be open elsewhere)
		["^:(.*)",	".*",		"POST", "{{F\\1}}",	[":{{CMD}}"],		"{{F\\1}}"], // an editing window starts with :/file and contains editing commands. They are sent to the server and apply for the same file (which should be open elsewhere)

		// embedded web page
		["^http://", "", "GET", undefined, ["Read"], undefined],
	]
	/////////////////////////////// no customizations should be done below /////////////////////////////
	this.fieldnames = ["filepattern", "commandpattern", "method", "infilepattern", "commandpatternvector", "newfilepattern"]

	// rows of Table can be index with fieldnames instead of rownames: e.g. this.Table[3].filepattern
	this.Table = []
	for (var i=0; i<this.table.length; i++) {
		this.Table[i] = {}
		for (var k=0; k<this.fieldnames.length; k++) {
			var fieldname = this.fieldnames[k]
			this.Table[i][fieldname] = this.table[i][k]
		}
	}

	// GetHeader returns the window headers for a new window with a given file name.
	this.GetHeader = function(file) {
		var headers = ""
		for (var pattern in this.header) {
			if (file.match(pattern) != null) {
				headers = this.header[pattern]
			}
		}
		return headers
	}

	// Lookup requests to match file and cmd to the Commander table.
	// It returns the name of the new window (newfile), the server request method and file parameter, and and the expanded server command list.
	this.Lookup = function(file, cmd, selectedText) {
	// return {
	// 	method: string: "GET" | "POST": server request method
	// 	cmds: string: primitive server command list, separated by ";"
	// 	file: string: the file that the server request get's as the parameter
	// 	newfile: string: new file name for the resulting window
	//	rulenumber: int: for debug: which rule number did the match
	// }
		if (cmd == undefined) {
			console.log("command is undefined: this is an error")
			return
		}
		
		var ret = {
			method: undefined,
			cmds: undefined,
			file: undefined,
			newfile: undefined,
			rulenumber: undefined,
		}
		for (var i=0; i<this.Table.length; i++) {
			var R = this.Table[i]
			if (file.match(R.filepattern) != null) {
				if (cmd.match(R.commandpattern) != null) {
					ret.method = R.method
					ret.cmds = this.expandCommands(file, cmd, R, selectedText)
					ret.file = this.expandFile(file, cmd, R, selectedText) // expand infilepattern in the current context
					ret.newfile = this.expandNewFile(file, cmd, R, selectedText) // expand newfilepattenr in the current context
					ret.rulenumber = i
				}
			}
		}
		
		// default fields
		if (ret.method == undefined)
			ret.method = "GET"
		if (ret.cmds == undefined)
			ret.cmds = []
		if (ret.file == undefined) {
			ret.file = file
		}
		if (ret.newfile == undefined)
			ret.newfile = file
			
		// strip address part from newfile
		var idx = ret.newfile.indexOf(":")
		if (idx != -1)
			ret.newfile = ret.newfile.substr(0,idx)

		return ret
	}

	// expandCommand expands the commandpatternvector of the Table element R
	this.expandCommands = function(file, cmd, R, selectedText) {
		var ret = []
		for (var i=0; i<R.commandpatternvector.length; i++) {
			ret.push(this.expand(R.commandpatternvector[i], file, cmd, R, selectedText))
		}
		return ret		
	}

	// expandFile expands the infilepattern of the Table element R
	this.expandFile = function(file, cmd, R, selectedText) {
		if (R.infilepattern == undefined)
			return undefined
		return this.expand(R.infilepattern, file, cmd, R, selectedText)	
	}

	// Expand command expands the newfilepattern of the Table element R
	this.expandNewFile = function(file, cmd, R, selectedText) {
		if (R.newfilepattern == undefined)
			return undefined
		return this.expand(R.newfilepattern, file, cmd, R, selectedText)
	}

	// expand expands the special variables:
	//	{{FILE}}	=> file name (window title)
	//	{{LOCALFILE}}	=> file name without path
	//	{{DIR}}		=> directory of the file 
	//	{{CMD}}		=> executed command
	//	{{LINE}}	=> line content for clicks on line numbers
	//	{{\1}}		=> 1st match of cmd in command regular expression
	//	{{\2}}		=> 2nd match
	//	{{\3}}		=> 3rd match
	//	{{F\\1}}	=> 1st match of file in the file regular expression
	//	{{TEXT}}	=> selected text
	// in the input string s
	this.expand = function(s, file, cmd, R, selectedText) {
		var matches = cmd.match(R.commandpattern)
		var filematches = file.match(R.filepattern)
		var pats = ["{{FILE}}", "{{LOCALFILE}}", "{{DIR}}", "{{CMD}}", "{{\\1}}", "{{\\2}}", "{{\\3}}", "{{F\\1}}", "{{TEXT}}"]
		for (var p=0; p<pats.length; p++) {
			var pattern = pats[p]
			while (true) {	// try as often until all specials are expanded
				var idx = s.indexOf(pattern)
				if (idx == -1)
					break
				var start = s.substr(0, idx)
				var tail = s.substr(idx+pattern.length)
				var expand = ""
				if (pattern == "{{FILE}}") {
					expand = file
				} else if (pattern == "{{LOCALFILE}}") {
					expand = file.substr(file.lastIndexOf("/")+1) // strip directory, works also if there is no "/"
				} else if (pattern == "{{CMD}}") {
					expand = cmd
				} else if (pattern == "{{DIR}}") {
					var li = file.lastIndexOf("/")
					if (li == -1)
						expand = file
					else
						expand = file.substr(0, li+1)
//				} else if (pattern == "{{LINE}}") {
//					expand = cmd.substr(5) // cmd = 'line:...'
				} else if ((pattern == "{{\\1}}") && (matches.length > 1)) {
					expand = matches[1]
				} else if ((pattern == "{{\\2}}") && (matches.length > 2)) {
					expand = matches[2]
				} else if ((pattern == "{{\\3}}") && (matches.length > 3)) {
					expand = matches[3]
				} else if ((pattern == "{{F\\1}}") && (filematches.length > 1)) {
					expand = filematches[1]
				} else if (pattern == "{{TEXT}}") {
					expand = selectedText
				}
				s = start + expand + tail
			}
		}
		return s
	}
}

