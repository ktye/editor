package cmd

import (
	"fmt"
	"unicode"
)

// splitQuoted splits a line at whitespace, except if it is quoted.
// Whitespace at the beginning or end of a string is trimmed.
//
// Example:
//	`alpha  beta ` => {"alpha", "beta"}
//	`alpha " beta "` => {"alpha", " beta "}
//	`alpha "beta \" gamma"` => {"alpha", `beta " gamma`}
//      `alpha|beta` => {"alpha","|","beta"},
func SplitQuoted(line string) ([]string, error) {
	var v []string
	quoted := false
	escaped := false
	trailingEmpty := false
	start := true
	s := ""
	for _, c := range line {
		if start {
			if unicode.IsSpace(c) {
				continue
			} else {
				start = false
			}
		}
		if (quoted == false && escaped == false) && unicode.IsSpace(c) {
			v = append(v, s)
			s = ""
			start = true
			trailingEmpty = false
			continue
		}
		if c == '\\' && quoted {
			if escaped {
				s += `\`
				escaped = false
				continue
			} else {
				escaped = true
				continue
			}
		}
		if c == '"' {
			if escaped {
				s += `"`
				escaped = false
				continue
			} else if quoted {
				quoted = false
				if s == "" {
					trailingEmpty = true
				}
				continue
			} else {
				quoted = true
				continue
			}
		}
		if escaped {
			return nil, fmt.Errorf("cannot escape '%c'", c)
		}
		s += string(c)
	}
	if escaped == true {
		return nil, fmt.Errorf("escape character at end of line")
	}
	if quoted == true {
		return nil, fmt.Errorf("unmatched quotation character")
	}
	if s != "" || trailingEmpty == true {
		v = append(v, s)
	}
	return v, nil
}

// SplitQuotedPipe splits the input line at '|' characters,
// if they are not quoted.
func SplitQuotedPipe(line string) []string {
	var v []string
	quoted := false
	escaped := false
	s := ""
	for _, c := range line {
		if quoted == false && c == '|' {
			v = append(v, s)
			s = ""
			continue
		}
		if c == '\\' {
			escaped = !escaped
		}
		if c == '"' && escaped == false {
			quoted = !quoted
		}
		s += string(c)
	}
	if s != "" {
		v = append(v, s)
	}
	return v
}
