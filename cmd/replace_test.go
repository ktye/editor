package cmd

import "testing"

func TestReplaceSelections(t *testing.T) {
	testCases := []struct {
		text       string
		selections RuneRanges
		repl       string
		newText    string
		rr         RuneRanges
	}{
		{"alpha", RuneRanges{{0, 0}}, "beta", "betaalpha", RuneRanges{{0, 4}}},
		{"alpha", RuneRanges{{0, 1}, {4, 5}}, "be", "belphbe", RuneRanges{{0, 2}, {5, 7}}},
		{"", RuneRanges{{0, 0}}, "αα", "αα", RuneRanges{{0, 2}}},
		{"alphaαgammaα", RuneRanges{{5, 6}, {11, 12}}, "xxx", "alphaxxxgammaxxx", RuneRanges{{5, 8}, {13, 16}}},
	}
	for _, tc := range testCases {
		c := Cmd{
			Text:       tc.text,
			Selections: tc.selections,
		}
		c.ReplaceSelections(tc.repl)
		if c.Text != tc.newText {
			t.Fatalf("%+v: got: %s\n", tc, c.Text)
		}
		if c.Selections.String() != tc.rr.String() {
			t.Fatalf("%+v: got: %s\n", tc, c.Selections.String())
		}
	}
}
