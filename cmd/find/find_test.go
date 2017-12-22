package main

import (
	"testing"

	"github.com/ktye/editor/cmd"
)

func TestFind(t *testing.T) {
	testCases := []struct {
		text string
		find string
		sel  string
	}{
		{"alphabeta", "a", "0:1,4:5,8:9"},
		{"xyz", "^", "0:0"},
		{"xyz", "$", "3:3"},
		{"xyz", "abc", ""},
		{"", "abc", ""},
	}
	for _, tc := range testCases {
		p := program{cmd.Cmd{Text: tc.text}}
		if err := p.find(tc.find); err != nil {
			t.Fatal(err)
		} else if sel := p.Selections.String(); sel != tc.sel {
			t.Fatalf("%+v: got %v\n", tc, sel)
		}
	}
}

func TestFindReplace(t *testing.T) {
	testCases := []struct {
		text string
		find string
		repl string
		res  string
		sel  string
	}{
		{"alphabeta", "a", "b", "blphbbetb", "0:1,4:5,8:9"},
		{"alphabeta", "a", "α", "αlphαbetα", "0:1,4:5,8:9"},
		{"xyz", "^", "α", "αxyz", "0:1"},
		{"xyz", "$", "www", "xyzwww", "3:6"},
		{"", "a", "b", "", ""},
		{"abc", "", "xxx", "abc", ""},
	}
	for _, tc := range testCases {
		p := program{cmd.Cmd{Text: tc.text}}
		if err := p.findReplace(tc.find, tc.repl); err != nil {
			t.Fatal(err)
		}
		if p.Text != tc.res {
			t.Fatalf("%+v, got text: %s\n", tc, p.Text)
		}
		if sel := p.Selections.String(); sel != tc.sel {
			t.Fatalf("%+v, got selections: %s\n", tc, sel)
		}
	}
}
