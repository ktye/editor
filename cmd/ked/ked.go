package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/png"
	"io/ioutil"
	"math"
	"math/bits"
	"os"
	"runtime"
	"strconv"
	"strings"
	"unsafe"

	"github.com/ktye/editor/cmd"
)

type program struct {
	cmd.Cmd
}

var prog program

func main() {
	if err := prog.Run(); err != nil {
		prog.Fatal(err)
	}
	prog.Exit()
}

var O *bytes.Buffer
var D *bytes.Buffer
var M *image.RGBA

func (p *program) Run() error {
	if e := prog.Parse(); e != nil {
		return e
	}
	if e := os.Chdir(prog.Directory()); e != nil {
		return e
	}
	if e := p.Forward("Write", nil); e != nil {
		return e
	}

	O = bytes.NewBuffer(nil)
	fmt.Fprintf(O, "")
	kinit()
	s := replace(trim(prog.Text))
	if fp, _ := prog.Path(); strings.HasSuffix(fp, ".j") {
		if e := kj(s, fp); e != nil {
			return nil
		}
	}
	run(s)

	prog.Name = p.TargetPath(prog.Directory()) + "/+Out"
	prog.Default = "read"
	args := prog.Args()
	if len(args) == 1 && args[0] == "-draw" {
		if M != nil {
			prog.Name = p.TargetPath(prog.Directory()) + "/+Img"
			var buf bytes.Buffer
			b := base64.NewEncoder(base64.StdEncoding, &buf)
			if e := png.Encode(b, M); e != nil {
				return e
			}
			b.Close()
			prog.Type = "html"
			prog.Text = fmt.Sprintf("<img width=%d height=%d src=\"data:image/png;base64, %s \"/>", M.Bounds().Dx(), M.Bounds().Dy(), template.HTMLEscapeString(string(buf.Bytes())))
			return nil
		} else if D != nil {
			prog.Name = p.TargetPath(prog.Directory()) + "/+Img"
			fmt.Fprintf(D, "window.editordiv.appendChild(canvas)\nwindow.editor={}}\n")
			prog.Type = "javascript"
			prog.Text = string(D.Bytes())
			return nil
		}
	}
	prog.Text = string(O.Bytes())
	return nil
}
func fatal(e error) {
	if e != nil {
		prog.Fatal(e)
	}
}

func kj(s, j string) error {
	k := strings.TrimSuffix(j, ".j") + ".k"
	return ioutil.WriteFile(k, []byte(s), 0644)
}
func trim(s string) string {
	if n := strings.Index(s, "\n\\\n"); n != -1 {
		return s[:n+1]
	}
	return s
}
func replace(s string) string {
	v := strings.Split(s, "\n")
	for i, s := range v {
		if len(s) > 1 && s[0] == '\\' && strings.HasSuffix(s, ".k") {
			file := s[1:]
			f, e := ioutil.ReadFile(file)
			fatal(e)
			v[i] = "/" + file + "\n" + trim(string(f))
		}
	}
	return strings.Join(v, "\n")
}
func run(t string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(O, stack())
			fmt.Fprintln(O, r)
		}
	}()
	if x := val(kstring(t)); x > 0 {
		x = kst(x)
		O.Write(MC[x+8 : x+nn(x)+8])
		O.Write([]byte{10})
		dx(x)
	}
}
func kinit() {
	m0 := 16
	MJ = make([]J, (1<<m0)>>3)
	msl()
	mt_init()
	ini(16)
}
func kstring(s string) I { x := mk(1, I(len(s))); copy(MC[x+8:], s); return x }
func stringk(x I) string { dx(x); return string(MC[8+x : 8+x+nn(x)]) }
func init() {
	NAN = math.Float64frombits(18444492273895866368) // 0/0 (not math.NaN)
	INFINITY = math.Inf(1)
}
func stack() string {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return string(buf[:n])
		}
		buf = make([]byte, 2*len(buf))
	}
	return string(buf)
}
func jk(x I) string {
	n := nn(x)
	if tp(x) != 1 && nn(x) != 1 {
		x = lx(x)
	}
	defer dx(x)
	switch tp(x) {
	case 1:
		return `"` + string(MC[8+x:8+x+n]) + `"`
	case 2:
		x >>= 2
		return strconv.Itoa(int(MI[2+x]))
	case 3:
		x >>= 3
		return strconv.FormatFloat(MF[1+x], 'g', -1, 64)
	case 6:
		x >>= 2
		r := make([]string, n)
		for i := range r {
			xi := MI[2+uint32(i)+x]
			rx(xi)
			r[i] = jk(xi)
		}
		return "(" + strings.Join(r, ",") + ")"
	}
	return ""
}

type C = byte
type I = uint32
type J = uint64
type F = float64
type SI = int32
type slice struct {
	p uintptr
	l int
	c int
}

var MC []C
var MI []I
var MJ []J
var MF []F
var NAN, INFINITY F

func sin(x F) F      { return math.Sin(x) }
func cos(x F) F      { return math.Cos(x) }
func exp(x F) F      { return math.Exp(x) }
func log(x F) F      { return math.Log(x) }
func atan2(x, y F) F { return math.Atan2(x, y) }
func hypot(x, y F) F { return math.Hypot(x, y) }
func draw(x, y, z I) {
	if tp(z) == 2 {
		M = image.NewRGBA(image.Rectangle{Max: image.Point{int(x), int(y)}})
		copy(M.Pix, MC[z+8:z+8+4*x*y])
		for i := 3; i < len(M.Pix); i += 4 {
			M.Pix[i] = 255
		}
	}
	if tp(z) == 7 {
		D = bytes.NewBuffer(nil)
		defer func() {
			if r := recover(); r != nil {
				fmt.Fprintln(D, stack())
				fmt.Fprintln(D, r)
			}
		}()
		k, v := MI[(z+8)>>2], MI[(z+12)>>2]
		rx(k)
		rx(v)
		n := nn(v)
		fmt.Fprintf(D, "function load(window){while(window.editordiv.firstChild)window.editordiv.removeChild(window.editordiv.lastChild)\n")
		fmt.Fprintf(D, "var canvas=document.createElement('canvas')\ncanvas.width=%d\ncanvas.height=%d\nvar ctx=canvas.getContext('2d');ctx.fillStyle=\"black\";ctx.fillRect(0,0,%d,%d)\n", x, y, x, y)
		for i := I(0); i < n; i++ {
			rx(k)
			rx(v)
			s := `"` + stringk(cs(atx(k, mki(i)))) + `"`
			a := jk(atx(v, mki(i)))
			ac := a
			if len(ac) > 0 && ac[0] != '(' {
				ac = "(" + a + ")"
			}
			if len(a) > 1 && a[0] == '(' {
				a = "undefined"
			}
			fmt.Fprintf(D, "if(typeof(ctx[%s])===\"function\")ctx[%s]%s;else ctx[%s]=%s\n", s, s, ac, s, a)
		}
	}
}
func msl() { // update slice headers after set/inc MJ
	cp := *(*slice)(unsafe.Pointer(&MC))
	ip := *(*slice)(unsafe.Pointer(&MI))
	jp := *(*slice)(unsafe.Pointer(&MJ))
	fp := *(*slice)(unsafe.Pointer(&MF))
	fp.l, fp.c, fp.p = jp.l, jp.c, jp.p
	ip.l, ip.c, ip.p = jp.l*2, jp.c*2, jp.p
	cp.l, cp.c, cp.p = ip.l*4, ip.c*4, ip.p
	MF = *(*[]F)(unsafe.Pointer(&fp))
	MI = *(*[]I)(unsafe.Pointer(&ip))
	MC = *(*[]byte)(unsafe.Pointer(&cp))
}
func grow(x I) I {
	if x > 31 {
		panic("oom")
	}
	c := make([]uint64, 1<<(x-3))
	copy(c, MJ)
	MJ = c
	msl()
	return x
}
func printc(x, y I) { fmt.Fprintf(O, "%s\n", string(MC[x:x+y])) }
func clz32(x I) I   { return I(bits.LeadingZeros32(x)) }
func clz64(x J) I   { return I(bits.LeadingZeros64(x)) }
func i32b(x bool) I {
	if x {
		return 1
	} else {
		return 0
	}
}
func n32(x I) I {
	if x == 0 {
		return 1
	} else {
		return 0
	}
}
