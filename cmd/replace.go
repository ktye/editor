package cmd

import "bytes"

// ReplaceSelections replaces all selections with repl.
// It selects the replaced strings and sets their selection.
func (c *Cmd) ReplaceSelections(repl string) {
	if c.Text == "" {
		c.Text = repl
		c.Selections = RuneRanges{{0, countRunes(repl)}}
	}

	var buf bytes.Buffer
	last := 0
	oldSelections := c.Selections
	c.Selections = nil
	off := 0
	for _, rr := range oldSelections {
		br := rr.ByteRange(c.Text)
		buf.Write([]byte(c.Text[last:br[0]]))
		buf.Write([]byte(repl))
		last = br[1]
		replCount := countRunes(repl)
		c.Selections = append(c.Selections, RuneRange{rr[0] + off, rr[0] + off + replCount})
		off += replCount - rr[1] + rr[0]
	}
	buf.Write([]byte(c.Text[last:len(c.Text)]))
	c.Text = string(buf.Bytes())
}

func countRunes(s string) int {
	n := 0
	for range s {
		n++
	}
	return n
}
