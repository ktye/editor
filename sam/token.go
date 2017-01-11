package sam

import (
	"errors"
	"regexp"
	"strconv"
)

const (
	cEND    int = iota // end of string
	cNUM               // number
	cSIGN              // + or -, this leads a relative address (number or regex)
	cDOT               // dot character
	cCOMMA             // comma address range separator
	cCOLON             // colon line:character separator
	cDOLLAR            // $ character
	cBLANK             // blank, actually space
	cRE                // regular expression
	cTEXT              // text (insert, append, change, substitute)
	cCMD               // command
)

type token struct {
	id int            // token id
	s  string         // original token string content
	n  int            // number for id == cNUM
	re *regexp.Regexp // regular expression for cRE
}

var regNUM *regexp.Regexp = regexp.MustCompile("^([0-9]+)(.*)")
var regBLANK *regexp.Regexp = regexp.MustCompile("(^ +)(.*)")
var regCMD *regexp.Regexp = regexp.MustCompile("^[aicdsxgvX]")

// nextToken scans the string s for the next token.
// It returns the tokenId, the code string of the token and the remaining string.
// On an error the tokenId is 0 and the error is returned in err.
// The expectText argument is used to differentiate between cRE and cTEXT as command arguments.
func nextToken(s string, expectText bool) (t token, rem string, err error) {
	var v []string

	// cEND
	if s == "" {
		return token{id: cEND, s: ""}, "", nil
	}

	// cNUM
	v = regNUM.FindStringSubmatch(s)
	if v != nil {
		n, _ := strconv.Atoi(v[1])
		return token{id: cNUM, s: v[1], n: n}, v[2], nil
	}

	// cSIGN
	if s[0] == '+' || s[0] == '-' {
		if s[0] == '+' {
			return token{id: cSIGN, s: string(s[0]), n: 1}, s[1:], nil
		} else {
			return token{id: cSIGN, s: string(s[0]), n: -1}, s[1:], nil
		}
	}

	// cDOT
	if s[0] == '.' {
		return token{id: cDOT, s: "."}, s[1:], nil
	}

	// cCOMMA
	if s[0] == ',' {
		return token{id: cCOMMA, s: ","}, s[1:], nil
	}

	// cCOLON
	if s[0] == ':' {
		return token{id: cCOLON, s: ":"}, s[1:], nil
	}

	// cDOLLAR
	if s[0] == '$' {
		return token{id: cDOLLAR, s: "$"}, s[1:], nil
	}

	// cBLANK
	v = regBLANK.FindStringSubmatch(s)
	if v != nil {
		return token{id: cBLANK, s: " "}, v[2], nil
	}

	// cRE | cTEXT
	if s[0] == '/' {
		quote := false
		for i := 1; i < len(s); i++ {
			if s[i] == '\\' {
				quote = !quote
			} else if quote { // the current character is quoted and skipped
				quote = false
			} else if s[i] == '/' {
				if expectText == true {
					return token{id: cTEXT, s: unquote(s[1:i])}, s[i+1:], nil
				} else {
					txt := unquote(s[1:i])
					// default flags m: match ^ and $ for lines
					// default flags can be overwritten with prefixing the regexp
					// e.g: revert to non-line mode: /(?-m).../
					// e.g: let . match \n: /(?s).../
					defaultflags := "(?m)"
					re, rerr := regexp.Compile(defaultflags + txt)
					if rerr != nil {
						return t, s, rerr
					}
					return token{id: cRE, re: re, s: unquote(s[1:i])}, s[i+1:], nil
				}
			}
		}
		return t, s, errors.New("untermintated regular expression or string: " + s)
	}

	// cCMD
	if regCMD.MatchString(s) {
		return token{id: cCMD, s: string(s[0])}, s[1:], nil
	}

	return t, s, errors.New("cannot parse: " + s)
}

// tokenize splits the command string s to it's tokens.
func tokenize(s string) (tokens []token, err error) {
	var t token
	var rem string
	rem = s
	expectText := false
	for {
		t, rem, err = nextToken(rem, expectText)
		if err != nil {
			return nil, err
		}
		if t.id == cEND {
			break
		}

		// Special case: blank is replaced with /^.*$/ if it is following an x command.
		if t.id == cBLANK && len(tokens) > 0 && tokens[len(tokens)-1].id == cCMD && tokens[len(tokens)-1].s == "x" {
			t.id = cRE
			t.s = "(?m)^.*$"
			t.re, err = regexp.Compile(t.s)
			if err != nil {
				return tokens, err
			}
		}

		tokens = append(tokens, t)

		// Decide if the next /string/ is a text or a regexp.
		expectText = false
		if t.id == cCMD {
			switch t.s {
			case "a", "i", "c":
				expectText = true
			}
		}
		if len(tokens) > 1 && tokens[len(tokens)-2].id == cCMD && tokens[len(tokens)-2].s == "s" {
			expectText = true
			// The substitute command uses only 1 slash to separate both arguments.
			// It has been removed already, put it back that the replacement text
			// can be recognized.
			rem = "/" + rem
		}
	}
	return tokens, nil
}

// Unqoute removes quotes from slash.
func unquote(s string) (u string) {
	re := regexp.MustCompile("\\\\/")
	u = re.ReplaceAllString(s, "/")
	return u
}
