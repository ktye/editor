// Find is an editor command which finds and replaces text.
//
// If called with a single argument, it sets the selection to all
// occurences of the regular expression passed as the argument.
//
// If called without an argument, it uses the FirstSelectedText
// as the regular expression.
//
// If called with 2 arguments it uses the second argument as the
// replace text.
// Within the replace text, $ signs are interpreted as in
// regexp.ReplaceAllString.
//
// Find (and replace) is only applied to the selected text, if there
// is any selection and find has not been called without arguments.
package main

//go:generate godocdown -output README.md

import (
	"fmt"
	"regexp"

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

	findString := ""
	args := p.Args()
	if len(args) == 0 {
		findString = p.FirstSelectedText()
		if len(findString) == 0 {
			return fmt.Errorf("find is called with no arguments and no selections")
		}
		p.Selections = nil
	} else if len(args) == 1 {
		findString = args[0]
	} else if len(args) == 2 {
		return p.findReplace(args[0], args[1])
	} else {
		return fmt.Errorf("too many argument for find")
	}

	return p.find(findString)
}

func (p *program) find(findString string) error {
	if p.Selections.Total() == 0 {
		p.Selections = cmd.RuneRanges{{0, countRunes(p.Text)}}
	}

	if re, err := regexp.Compile(findString); err != nil {
		return err
	} else {
		var byteRangeMatches [][2]int
		for _, rr := range p.Selections {
			br := rr.ByteRange(p.Text)
			idx := re.FindAllStringIndex(p.Text[br[0]:br[1]], -1)
			for i := range idx {
				byteRangeMatches = append(byteRangeMatches, [2]int{br[0] + idx[i][0], br[0] + idx[i][1]})
			}
		}
		p.Selections = make(cmd.RuneRanges, len(byteRangeMatches))
		for i, br := range byteRangeMatches {
			p.Selections[i] = p.ByteRangeToRuneRange(br[0], br[1])
		}
		return nil
	}
}

func (p *program) findReplace(find, repl string) error {
	if find == "" || repl == "" {
		p.Selections = nil
		return nil
	}

	re, err := regexp.Compile(find)
	if err != nil {
		return err
	}

	idx := re.FindAllStringSubmatchIndex(p.Text, -1)
	if idx == nil {
		p.Selections = nil
		return nil
	}

	if p.Selections.Total() == 0 {
		p.Selections = cmd.RuneRanges{{0, countRunes(p.Text)}}
	}

	idx = p.restrictToSelections(idx)

	p.Selections = nil
	head := p.Text[0:idx[0][0]]
	dst := []byte(head)
	start := countRunes(head)
	for i, match := range idx {
		b := re.ExpandString(nil, repl, p.Text, match)
		n := countRunes(string(b))
		p.Selections = append(p.Selections, cmd.RuneRange{start, start + n})
		end := len(p.Text)
		if i < len(idx)-1 {
			end = idx[i+1][0]
		}
		t := p.Text[match[1]:end]
		start += n + countRunes(t)
		dst = append(dst, b...)
		dst = append(dst, []byte(t)...)
	}
	p.Text = string(dst)
	return nil
}

// restrictToSelections removes byte range indexes, which are outside any initial selection.
func (p *program) restrictToSelections(idx [][]int) [][]int {
	byteRanges := make([][2]int, len(p.Selections))
	for i, rr := range p.Selections {
		byteRanges[i] = rr.ByteRange(p.Text)
	}

	var out [][]int
	for _, r := range idx {
		for _, b := range byteRanges {
			if r[0] >= b[0] && r[1] <= b[1] {
				out = append(out, r)
				continue
			}
		}
	}
	return out
}

func countRunes(s string) int {
	n := 0
	for range s {
		n++
	}
	return n
}
