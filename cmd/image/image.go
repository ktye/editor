// Image is an editor command to display images.
//
// The images are drawn by javascript on a canvas.
// The image may not show up directly, click into it. // TODO: how to fix?
package main

//go:generate godocdown -output README.md

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/ktye/editor/cmd"
)

type program struct {
	cmd.Cmd
}

func main() {
	var p program
	if err := p.Run(); err != nil {
		p.Fatal(err)
	}
	p.Exit()
}

func (p *program) Run() error {
	if err := p.Parse(); err != nil {
		return err
	}

	args := p.Args()
	if len(args) == 0 {
		return p.read()
	} else if args[0] == "-crop" {
		return p.crop()
	} else if args[0] == "-resize" {
		return p.resize(args[1:])
	}
	return fmt.Errorf("image: unknown arguments: %v", args)
}

// Decode tries to decode the image from p.Text.
func (p *program) decode() (image.Image, error) {
	prefix := "data:image/png;base64,"
	if strings.HasPrefix(p.Text, prefix) == false {
		return nil, fmt.Errorf("image has wrong data prefix")
	}
	if b, err := base64.StdEncoding.DecodeString(p.Text[len(prefix):]); err != nil {
		return nil, err
	} else {
		return png.Decode(bytes.NewReader(b))
	}
}

func (p *program) read() error {
	path, _ := p.Path()
	ext := filepath.Ext(path)
	if dec := decoders[ext]; dec == nil {
		return fmt.Errorf("no decoder for image file extension '%s'", ext)
	} else {
		if r, err := os.Open(path); err != nil {
			return err
		} else {
			defer r.Close()
			if im, err := dec.Decode(r); err != nil {
				return err
			} else {
				return p.send(im)
			}
		}
	}
}

func (p *program) send(im image.Image) error {
	var d struct {
		Data   string
		Width  int
		Height int
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, im); err != nil {
		return err
	}
	d.Data = base64.StdEncoding.EncodeToString(buf.Bytes())
	d.Width = im.Bounds().Dx()
	d.Height = im.Bounds().Dy()

	t, err := template.New("image").Parse(tpl)
	if err != nil {
		return err
	}

	var text bytes.Buffer
	if err := t.Execute(&text, d); err != nil {
		return err
	}

	p.Type = "javascript"
	p.Text = string(text.Bytes())
	p.Tags = "-crop -resize 50%"
	p.Default = "image"
	return nil
}

const tpl = `

function load(window) {
	var image = new Image({{.Width}}, {{.Height}})
	var data = "data:image/png;base64,{{.Data}}"
	image.src = data
	
	var canvas = document.createElement('canvas')
	canvas.width = {{.Width}}
	canvas.height = {{.Height}}
	
	var ctx = canvas.getContext('2d')
	var rect = {}
	var drag = false
	var draw = function() {
		console.log('draw rect=', rect)
		ctx.drawImage(image, 0, 0)
		ctx.beginPath()
		ctx.rect(rect.startX, rect.startY, rect.width, rect.height)
		ctx.lineWidth = 1
		ctx.strokeStyle = 'red'
		ctx.stroke()
	}
	var mouseDown = function(e) {
		var r = canvas.getBoundingClientRect()
		rect.startX = e.clientX - r.left
		rect.startY = e.clientY - r.top
		drag = true
	}
	var mouseUp = function(e) {
		drag = false
	}
	var mouseMove = function(e) {
		if (drag) {
			var r = canvas.getBoundingClientRect()
			rect.width = (e.clientX - r.left) - rect.startX
			rect.height = (e.clientY - r.top) - rect.startY
			ctx.clearRect(0, 0, canvas.width, canvas.height)
			draw()
		}
	}
	var getSelections = function() {
		if (rect.startX != undefined) {
			var x = Math.round(rect.startX)
			var y = Math.round(rect.startY)
			var w = Math.round(rect.width)
			var h = Math.round(rect.height)
			return [[x, y], [w, h]]
		} else {
			return []
		}
	}
	
	canvas.addEventListener('mousedown', mouseDown, false);
	canvas.addEventListener('mouseup', mouseUp, false);
	canvas.addEventListener('mousemove', mouseMove, false);
	ctx.drawImage(image, 0, 0)
	
	window.editordiv.appendChild(canvas)
	window.editor = {}
	window.editor.GetSelections = getSelections
	window.editor.GetText = function() { return data }
}
`
