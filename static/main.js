var layout
var commander
function init() {
	commander = new Commander()
	layout = new Layout()
	var col = layout.AddColumn()
	col.AddWindow(0)
	layout.columns[0].windows[0].SetTitle('/')
	layout.columns[0].windows[0].Execute("")
	console.log('page loaded')
}
