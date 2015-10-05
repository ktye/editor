package sam

import (
	"errors"
	"regexp"
	"strconv"
)

// ParseAddress parses an address range from tokens.
// It updates f.dot and f.pos
func (f *Samfile) parseAddressRange() (err error) {
	var newdot Address
	newdot.from = f.dot.from
	newdot.to = f.dot.to
	section := 1 // 1: from range, 2: to range, 3:error
	isfirst := true
	relative := 0 // -1:"-", 0:absolute, 1:"+", 2:":"
	for {
		t := f.currentToken()
		switch t.id {
		case cEND, cCMD:
			if relative == 1 || relative == -1 { // address ends in + or -: implicitly add 1
				newdot = f.parseRelativeLineNumber(newdot, section, relative)
			} else if f.pos > 0 && f.tokens[f.pos-1].id == cCOMMA { // address ends in ,: implicitly add $
				newdot.to = len(f.b)-1
			}
			// Addresses that stretch too far are clipped to the last position.
			if newdot.from >= len(f.b) {
				newdot.from = len(f.b)-1
			}
			if newdot.to >= len(f.b) {
				newdot.to = len(f.b)-1
			}
			f.dot.from = newdot.from
			f.dot.to = newdot.to
			return nil
		case cNUM:
			switch relative {
			case 0: // The number is an absolute line address.
				if t.n == 0 { // special case: "0"
					if section == 1 {
						newdot.from = 0
					}
					newdot.to = 0
				} else {
					if t.n-1 < 0 || t.n-1 >= len(f.lineAddresses) {
						return errors.New("line address out of range: " + t.s)
					}
					if section == 1 {
						newdot.from = f.lineAddresses[t.n-1].from
					}
					newdot.to = f.lineAddresses[t.n-1].to
				}
			case 2: // The number represents a character position.
				if section == 1 {
					newdot.from += t.n
				} else { // Go to beginning of line and advance number of char positions.
					line, _ := LineAddr(newdot.to, false, f.lineAddresses, f.b)
					newdot.to = f.lineAddresses[line-1].from + t.n
				}
				relative = 0
			case -1, 1: // The number is a relative line address.
				newdot = f.parseRelativeLineNumber(newdot, section, relative*t.n)
				relative = 0
			}
		case cSIGN:
			if relative == -1 || relative == 1 { // implicit 1
				newdot = f.parseRelativeLineNumber(newdot, section, relative)
			}
			relative = t.n
		case cDOT:
			if section == 1 {
				if !isfirst {
					return errors.New("unexpected . in line address")
				}
			} else {
				if relative != 0 {
					return errors.New("unexpected . in line address")
				}
				newdot.to = f.dot.to
			}
		case cCOMMA:
			if isfirst { // implicit 0
				newdot.from = 0
				newdot.to = 0
			}
			section++
			if section > 2 {
				return errors.New("2 commas in address")
			}
		case cCOLON:
			relative = 2
		case cDOLLAR:
			if relative != 0 {
				return errors.New("unexpected $ in relative address")
			}
			if section == 1 {
				newdot.from = len(f.b)-1
			}
			newdot.to = len(f.b)-1
		case cBLANK: // ignore
		case cRE:
			dir := 1
			if relative == -1 {
				dir = -1
			}
			var readdr Address
			readdr, err = f.parseRegexpAddress(t.re, dir, newdot)
			if err != nil {
				return err
			}
			if section == 1 {
				newdot.from = readdr.from
			}
			newdot.to = readdr.to
			relative = 0
		default:
			return errors.New("this should not happen: unknown token in parseAddressRange: " + strconv.Itoa(int(t.id)))
		}
		isfirst = false

		f.pos++
	}
}

// ParseRegexpAddress parses a regular expression address.
func (f *Samfile) parseRegexpAddress(re *regexp.Regexp, dir int, dot Address) (newdot Address, err error) {
	if dir == 1 {
		// In normal forward searching, find next occurance starting after dot.
		idx := re.FindIndex(f.b[dot.to:])
		if idx != nil {
			newdot.from = idx[0] + dot.to
			newdot.to = idx[1] + dot.to
			return newdot, nil
		}
		// If there is none, restart from the beginning of the file to the dot.
		idx = re.FindIndex(f.b[0:dot.to])
		if idx != nil {
			newdot.from = idx[0]
			newdot.to = idx[1]
			return newdot, nil
		}
		// Still no match: return dot unchanged.
		return dot, nil
	}
	// Backward searching must be implemented with FindAll.
	idxs := re.FindAllIndex(f.b, -1)
	if idxs == nil {
		// No matches, return dot unchanged.
		return dot, nil
	}
	// Look for last match before dot.
	for i := 0; i < len(idxs); i++ {
		if idxs[i][1] > dot.from {
			if i > 0 {
				newdot.from = idxs[i-1][0]
				newdot.to = idxs[i-1][1]
				return newdot, nil
			}
		}
	}
	// No match before the dot, restart from the end.
	for i := len(idxs) - 1; i >= 0; i-- {
		if idxs[i][0] < dot.to {
			if i != len(idxs)-1 {
				newdot.from = idxs[i+1][0]
				newdot.to = idxs[i+1][1]
				return newdot, nil
			}
		}
	}
	return dot, errors.New("implementation error: cannot reverse find")
}

// ParseRelativeLineNumber parses the next line number which is relative.
func (f *Samfile) parseRelativeLineNumber(dot Address, section, n int) (newdot Address) {
	newdot.from = dot.from
	newdot.to = dot.to

	// Relative line numbers add to the end of the current dot.
	line, _ := LineAddr(newdot.to, false, f.lineAddresses, f.b)
	if newdot.from >= len(f.b) { // The current line is after the last line. This is valid e.g: /alpha/+- with alpha in the last line.
		line++
	}
	line += n
	if line < 1 {
		if section == 1 {
			newdot.from = 0
		}
		newdot.to = 0
	} else if line > len(f.lineAddresses) { // Set current position after the last line. Line is 1-based.
		if section == 1 {
			newdot.from = len(f.b)
		}
		newdot.to = len(f.b)
	} else {
		if section == 1 {
			newdot.from = f.lineAddresses[line-1].from
		}
		newdot.to = f.lineAddresses[line-1].to
	}
	return newdot
}
