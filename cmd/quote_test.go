package cmd

import "testing"

func TestSplitQuotedPipe(t *testing.T) {
	testCases := []struct {
		s        string
		expected []string
	}{
		{`alpha|beta`, []string{"alpha", "beta"}},
		{`alpha|"beta"`, []string{"alpha", `"beta"`}},
		{`alpha|beta|`, []string{"alpha", "beta"}},
		{`"alpha|beta"`, []string{`"alpha|beta"`}},
		{`"alpha\"|beta"`, []string{`"alpha\"|beta"`}},
	}
	for i, tc := range testCases {
		if v := SplitQuotedPipe(tc.s); equalStrings(v, tc.expected) == false {
			t.Fatalf("%d: %q: got %q\n", i, tc.s, v)
		}
	}
}

func TestSplitQuoted(t *testing.T) {
	testCases := []struct {
		s        string
		expected []string
	}{
		{`alpha beta`, []string{"alpha", "beta"}},
		{`alpha  beta`, []string{"alpha", "beta"}},
		{` alpha beta `, []string{"alpha", "beta"}},
		{`  alpha   beta   `, []string{"alpha", "beta"}},
		{` alpha beta gamma`, []string{"alpha", "beta", "gamma"}},
		{`"alpha beta " gamma`, []string{"alpha beta ", "gamma"}},
		{`"alpha\"beta"`, []string{`alpha"beta`}},
		{`\"`, nil},
		{`alpha\beta`, []string{`alpha\beta`}},
		{`"\""`, []string{`"`}},
		{`"\\"`, []string{`\`}},
		{`\abc`, []string{`\abc`}},
		{`"alpha`, nil},
		{`a ""b`, []string{"a", "b"}},
		{`a b""`, []string{"a", "b"}},
		{`a "" b`, []string{"a", "", "b"}},
		{`a "" b `, []string{"a", "", "b"}},
		{`a ""`, []string{"a", ""}},
	}
	for i, tc := range testCases {
		v, err := SplitQuoted(tc.s)
		if err != nil && v != nil {
			t.Fatalf("%d: %q: %s\n", i, tc.s, err)
		} else if err == nil && v == nil {
			t.Fatalf("%d: %q: should have failed\n", i, tc.s)
		} else if equalStrings(v, tc.expected) == false {
			t.Fatalf("%d: %q: got %q\n", i, tc.s, v)
		}
	}
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
