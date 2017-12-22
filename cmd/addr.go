package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// RuneRange defines the start and end address of a selection.
// The address is the rune index starting at 0.
// An empty selection defines a cursor position and has equal
// start and end address.
// The length of the selection is end-start, the same as
// slice indexing.
type RuneRange [2]int

type RuneRanges []RuneRange

func (rr RuneRanges) String() string {
	v := make([]string, len(rr))
	for i, r := range rr {
		v[i] = fmt.Sprintf("%d:%d", r[0], r[1])
	}
	return strings.Join(v, ",")
}

// Total returns the number of selected runs.
func (rr RuneRanges) Total() int {
	total := 0
	for _, r := range rr {
		total += r[1] - r[0]
	}
	return total
}

// ByteRange converts a RuneRange to a byte range.
func (rr RuneRange) ByteRange(text string) (b [2]int) {
	if rr[0] < 0 {
		rr[0] = 0
	}
	b[1] = -1
	runeIndex := 0
	for i := range text {
		if rr[0] == runeIndex {
			b[0] = i
		}
		if rr[1] == runeIndex {
			b[1] = i
		}
		runeIndex++
	}
	if b[1] < b[0] {
		b[1] = len(text)
	}
	return b
}

// Slice returns the text defined by the RuneRange.
func (rr RuneRange) Slice(text string) string {
	byteRange := rr.ByteRange(text)
	return text[byteRange[0]:byteRange[1]]
}

// CombinedSelectedText joins the selected text strings by newline.
func (c *Cmd) CombinedSelectedText() string {
	if c.Selections == nil {
		return ""
	}
	v := make([]string, len(c.Selections))
	for i, rr := range c.Selections {
		v[i] = rr.Slice(c.Text)
	}
	return strings.Join(v, "\n")
}

// FirstSelectedText return the text of the first selection.
func (c *Cmd) FirstSelectedText() string {
	if len(c.Selections) < 1 {
		return ""
	}
	return c.Selections[0].Slice(c.Text)
}

// SelectAddress sets c.Selection from the given address.
// The address has the form:
// N             select line N
// N:M           select line N starting at character M
// /reg/         select the first match of the regular expression
// /reg/g        select all matches of the regular expression
// If the last character is a colon, it is removed.
func (c *Cmd) SelectAddress(addr string) {
	// When selecting file:N:M: xxx, we remove the trailing colon.
	if len(addr) > 0 && addr[len(addr)-1] == ':' {
		addr = addr[:len(addr)-1]
	}
	if len(addr) == 0 {
		return
	}

	line := regexp.MustCompile(`^([0-9]+)$`)
	lineChar := regexp.MustCompile(`^([0-9]+):([0-9]+)$`)

	if v := line.FindStringSubmatch(addr); v != nil {
		n, _ := strconv.Atoi(v[1])
		c.selectLine(n-1, 0)
	} else if v := lineChar.FindStringSubmatch(addr); v != nil {
		n, _ := strconv.Atoi(v[1])
		m, _ := strconv.Atoi(v[2])
		c.selectLine(n-1, m-1)
	} else if addr[0] == '/' && addr[len(addr)-1] == '/' {
		s := addr[1 : len(addr)-1]
		if re, err := regexp.Compile(s); err != nil {
			return
		} else {
			if v := re.FindStringIndex(c.Text); v == nil {
				return
			} else {
				rr := c.ByteRangeToRuneRange(v[0], v[1])
				c.Selections = []RuneRange{rr}
			}
		}
	} else if addr[0] == '/' && strings.HasSuffix(addr, "/g") {
		s := addr[1 : len(addr)-2]
		if re, err := regexp.Compile(s); err != nil {
			return
		} else {
			if v := re.FindAllStringIndex(c.Text, -1); v == nil {
				return
			} else {
				var selections []RuneRange
				for _, idx := range v {
					rr := c.ByteRangeToRuneRange(idx[0], idx[1])
					selections = append(selections, rr)
				}
				c.Selections = selections
			}
		}
	}
}

// selectLine sets the Selection to the given line starting at the character position.
// Both line and char indexes start at 0.
// The newline is included in the selection.
func (cmd *Cmd) selectLine(line, char int) {
	l := 0
	c := 0
	var from, to int
	for i, r := range cmd.Text {
		if l > line {
			break
		}
		if l == line && c <= char {
			from = i
		}

		to = i
		if r == '\n' {
			l++
			c = 0
		} else {
			c++
		}
	}
	to++
	if to < from {
		to = from
	}
	cmd.Selections = []RuneRange{{from, to}}
}

// ByteRangeToRuneRange converts a byte range to a rune range.
func (cmd *Cmd) ByteRangeToRuneRange(from, to int) RuneRange {
	start := 0
	end := -1
	runeIdx := 0
	for i := range cmd.Text {
		if from == i {
			start = runeIdx
		}
		if to == i {
			end = runeIdx
			break
		}
		runeIdx++
	}
	if end < start {
		end = runeIdx
	}
	if from >= len(cmd.Text) {
		start = runeIdx
	}
	return RuneRange{start, end}
}
