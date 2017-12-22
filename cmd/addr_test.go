package cmd

import "testing"

func TestSelectAddress(t *testing.T) {
	testCases := []struct {
		text      string
		addr      string
		selection RuneRanges
	}{
		{"1\n2\n3", "1", RuneRanges{{0, 2}}},
		{"1\n2\n3", "1:", RuneRanges{{0, 2}}},
		{"1\n2\n3", "2", RuneRanges{{2, 4}}},
		{"a\nbc\n", "2:2", RuneRanges{{3, 5}}},
		{"alpha", "/ph/", RuneRanges{{2, 4}}},
		{"alpha\nbeta\nph ph", "/ph/g", RuneRanges{{2, 4}, {11, 13}, {14, 16}}},
		{"alphα\nbeta\nα", "/α/g", RuneRanges{{4, 5}, {11, 12}}},
	}
	for _, tc := range testCases {
		cmd := Cmd{
			Text: tc.text,
		}
		cmd.SelectAddress(tc.addr)
		if cmd.Selections.String() != tc.selection.String() {
			t.Fatalf("%v: got %s", tc, cmd.Selections)
		}
	}
}

func TestSelectLine(t *testing.T) {
	testCases := []struct {
		text       string
		line, char int
		rr         RuneRange
	}{
		{"1\n\n12\n123", 0, 0, RuneRange{0, 2}},
		{"1\n\n12\n123", 1, 0, RuneRange{2, 3}},
		{"1\n\n12\n123", 2, 0, RuneRange{3, 6}},
		{"1\n\n12\n123", 3, 1, RuneRange{7, 9}},
	}

	for _, tc := range testCases {
		cmd := Cmd{
			Text: tc.text,
		}
		cmd.selectLine(tc.line, tc.char)
		if len(cmd.Selections) != 1 || cmd.Selections[0] != tc.rr {
			t.Fatalf("%v: got %s", tc, cmd.Selections)
		}
	}
}

func TestByteRangeToSelection(t *testing.T) {
	testCases := []struct {
		text     string
		from, to int // byte range
		rr       RuneRange
	}{
		{"alpha", 0, 0, RuneRange{0, 0}},
		{"alpha", 0, 1, RuneRange{0, 1}},
		{"alpha", 2, 4, RuneRange{2, 4}},
		{"ab", 0, 2, RuneRange{0, 2}},
		{`α character`, 0, 2, RuneRange{0, 1}},
		{`α character`, 4, 6, RuneRange{3, 5}},
		{"xyz", 3, 3, RuneRange{3, 3}},
	}
	for _, tc := range testCases {
		cmd := Cmd{
			Text: tc.text,
		}
		rr := cmd.ByteRangeToRuneRange(tc.from, tc.to)
		if rr != tc.rr {
			t.Fatalf("%+v: got %v", tc, rr)
		}
	}
}

func TestByteRange(t *testing.T) {
	testCases := []struct {
		text      string
		rr        RuneRange
		byteRange [2]int
	}{
		{"", RuneRange{0, 0}, [2]int{0, 0}},
		{"alpha", RuneRange{0, 0}, [2]int{0, 0}},
		{"alpha", RuneRange{0, 1}, [2]int{0, 1}},
		{"a", RuneRange{0, 1}, [2]int{0, 1}},
		{"aαb", RuneRange{1, 2}, [2]int{1, 3}},
		{"aαb", RuneRange{2, 3}, [2]int{3, 4}},
	}
	for _, tc := range testCases {
		if br := tc.rr.ByteRange(tc.text); br != tc.byteRange {
			t.Fatalf("%v: got %v", tc, br)
		}
	}
}
