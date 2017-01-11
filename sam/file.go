package sam

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Address is the byte range in the byte array for a selection
// It is 0 based. The selection has the length to-from
// The first character is 0,1 and the empty selection
// before the first character is 0,0.
// The content of the address can be selected with slicing in the form b[from:to].
type Address struct {
	from int // start position in byte array
	to   int // end position in byte array
}

// Length returns the length of the Address.
func (a *Address) Length() int {
	return a.to - a.from
}

// GetRange returns the start and end position of the Address.
func (a *Address) GetRange() (from, to int) {
	return a.from, a.to
}

type Samfile struct {
	b             []byte    // file content
	tokens        []token   // tokenized commands
	lineAddresses []Address // line address ranges
	pos           int       // current parsing position in tokens slice
	dot           Address   // current dot
}

// InitSamFile returns the samfile struct.
// If the file is not terminated by '\n', the final newline is added for non-empty files.
func InitSamFile(b []byte, cmd string, initDot string) (f Samfile, initDots []Address, err error) {
	f.b = b

	// Make sure the file terminated by a newline.
	if len(f.b) > 0 && f.b[len(f.b)-1] != '\n' {
		f.b = append(f.b, '\n')
	}

	// Convert addresses given in initial dot string to byte addresses.
	f.lineAddresses = GetLineAddresses(b)
	initDots, err = f.parseInitialDot(initDot)
	if err != nil {
		return f, initDots, err
	}
	if len(initDots) == 1 && initDots[0].to == -1 {
		initDots[0].to = len(f.b)
	}
	for i := 0; i < len(initDots); i++ {
		addr := initDots[i]
		if addr.from < 0 || addr.to > len(f.b) {
			return f, initDots, errors.New("address range of initial dot exceeds content")
		}
	}

	// Tokenize the command
	f.tokens, err = tokenize(cmd)
	return f, initDots, err
}

// ParseInitialDot parses the initial dot string and returns its addresses representation.
func (f *Samfile) parseInitialDot(s string) (a []Address, err error) {
	if s == "" {
		a = append(a, Address{from: 0, to: -1}) // will be expanded to whole file
		return a, nil
	}

	rangex := regexp.MustCompile("^([0-9]+):([0-9]+),([0-9]+):([0-9]+)$")
	v := strings.Split(s, ";")
	for i := 0; i < len(v); i++ {
		w := rangex.FindStringSubmatch(v[i])
		if len(w) != 5 {
			return a, errors.New("initial dot address ranges must have the form L:C,L:C")
		}
		fromline, _ := strconv.Atoi(w[1])
		fromchar, _ := strconv.Atoi(w[2])
		toline, _ := strconv.Atoi(w[3])
		tochar, _ := strconv.Atoi(w[4])
		if fromline < 1 || fromline > len(f.lineAddresses) || toline < 1 || toline > len(f.lineAddresses) {
			return a, errors.New("initial line address is out of range: " + v[i])
		}
		addr := Address{from: f.lineAddresses[fromline-1].from + fromchar - 1, to: f.lineAddresses[toline-1].from + tochar - 1}
		a = append(a, addr)
	}
	return a, nil
}

// GetLineAddresses returns a slice of addresses which represent the lines in b.
// Example:
// 	abc|x||y|  (| indicates \n)
// 	012345678: [0:3, "abc"], [4:5, "x"], [6:6, ""], [7:8, "y"]
func GetLineAddresses(b []byte) (a []Address) {
	// empty input
	if len(b) == 0 {
		a = append(a, Address{0, 0})
		return a
	}

	start := 0
	for i := 0; i < len(b); i++ {
		if b[i] == '\n' {
			a = append(a, Address{start, i})
			start = i + 1
		}
	}
	return a
}

// LineAddr converts from byte position to a line and character position.
// Both line and ch are starting with 1.
// lineAddresses must have been updated before.
// b is the file content.
// There is an ambiguity when the position of a newline is passed:
// This can be the first character of the following line (from=true)
// or the last character of the previous line (from=false).
func LineAddr(p int, from bool, lineAddresses []Address, b []byte) (line, ch int) {
	if len(lineAddresses) == 0 {
		return 0, 0
	}

	// The address is the last character in the file.
	if p == len(b)-1 {
		i := len(lineAddresses) - 1
		return i + 1, lineAddresses[i].to - lineAddresses[i].from + 1
	}

	for i := 0; i < len(lineAddresses); i++ {
		a := lineAddresses[i]
		if p <= a.to {
			return i + 1, p - a.from + 1
		}
	}
	// Position is outside the address range.
	i := len(lineAddresses) - 1
	return len(lineAddresses), p - lineAddresses[i].from + 1
}
