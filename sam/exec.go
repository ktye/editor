package sam

import (
	"errors"
)

// Exec executes the command in samfile for the current dot.
// It returns the address to which the new content has been inserted
// and all new dot addresses.
func (f *Samfile)exec() (insertAddress Address, newDots []Address, err error) {
	f.pos = 0
	if len(f.tokens) == 0 {
		return f.dot, []Address{f.dot}, nil
	}
	err = f.parseAddressRange()
	if err != nil {
		return insertAddress, newDots, err
	}
	t := f.currentToken()
	if t.id == cEND {
		// No command follows, only new address selection is returned.
		return f.dot, []Address{f.dot}, nil
	}
	// Execute command part.
	return f.execCommand()
}

// ExecCommand executes the command part of the sam command after the address parts has been executed.
func (f *Samfile)execCommand() (insertAddress Address, newDots []Address, err error) {
	t := f.currentToken()

	// Commands may be preceded by blanks.
	if t.id == cBLANK {
		f.pos++
		t = f.currentToken()
	}
	if t.id != cCMD {
		return insertAddress, newDots, errors.New("command expected: "+t.s)
	}

	switch t.s {
	case "a", "i", "c":
		c := t.s
		f.pos++
		t = f.currentToken()
		if t.id != cTEXT {
			return insertAddress, newDots, errors.New("command "+c+" must be followed by append text: "+t.s)
		}
		newContent := []byte(t.s)
		var newdot Address
		if c == "a" { // append (after dot)
			insertAddress.from = f.dot.to
			insertAddress.to = f.dot.to
			newdot = Address{f.dot.to, f.dot.to+len(newContent)}
		} else if c == "i" { // insert (before dot)
			insertAddress.from = f.dot.from
			insertAddress.to = f.dot.from
			newdot = Address{f.dot.from, f.dot.from+len(newContent)}
		} else if c == "c" { // change (replace dot)
			insertAddress = f.dot
			newdot = Address{f.dot.from, f.dot.from+len(newContent)}
		}
		newDots = append(newDots, newdot)
		f.dot = newdot
		f.insertContent(insertAddress, newContent)
		return insertAddress, newDots, nil
	case "d": // delete (dot)
		insertAddress = f.dot
		f.dot.from = f.dot.from
		f.dot.to = f.dot.from
		newDots = append(newDots, f.dot)
		f.insertContent(insertAddress, nil)
		return insertAddress, newDots, nil
	case "s": // substitute
		f.pos++
		t = f.currentToken()
		if t.id != cRE {
			return insertAddress, newDots, errors.New("command s must be followed by a regexp: "+t.s)
		}
		re := t.re
		f.pos++
		t = f.currentToken()
		s := t.s
		if t.id != cTEXT {
			return insertAddress, newDots, errors.New("command s misses the replacement text: "+t.s)
		}
		b := f.b[f.dot.from:f.dot.to]
		idx := re.FindSubmatchIndex(b)
		if idx == nil {
			// No match found: dot remains unchanged
			insertAddress = f.dot
			newDots = append(newDots, f.dot)
			return insertAddress, newDots, nil
		}
		newContent := re.Expand(nil, []byte(s), b, idx)
		insertAddress.from = f.dot.from+idx[0]
		insertAddress.to = f.dot.from+idx[1]
		newDots = append(newDots, Address{f.dot.from+idx[0], f.dot.from+idx[0]+len(newContent)})
		f.insertContent(insertAddress, newContent)
		return insertAddress, newDots, nil
	case "g", "v": // conditional if
		c := t.s
		f.pos++
		t = f.currentToken()
		if t.id != cRE {
			return insertAddress, newDots, errors.New("command "+c+" must be followed by a regexp: "+t.s)
		}
		m := t.re.Match(f.b[f.dot.from:f.dot.to])
		if (m == true && c == "g") || (m == false && c == "v") {
			f.pos++
			return f.execCommand()
		}
		// If the conditional is not executed, return the unchanged dot.
		newDots = append(newDots, f.dot)
		return f.dot, newDots, nil
	case "x": // loop. The original sam's x command may contain an address after the regexp. This is not implemented.
		f.pos++
		t = f.currentToken()

		// Special case: cBLANK is treated as /^.*$/ is already replaced by the tokenizer.
		if t.id != cRE {
			return insertAddress, newDots, errors.New("command x must be followed by a regexp: "+t.s)
		}

		idxs := t.re.FindAllIndex(f.b[f.dot.from:f.dot.to], -1)

		// No match, let dot unchanged.
		if idxs == nil {
			insertAddress = f.dot
			newDots = append(newDots, insertAddress)
			return insertAddress, newDots, nil
		}

		// Add offset to idxs, because the match was local to the current dot.
		for i:=0; i<len(idxs); i++ {
			idxs[i][0] += f.dot.from
			idxs[i][1] += f.dot.from
		}

		f.pos++	// Position of next command.
		commandPos := f.pos; // Save this position and reset it before every execution in the loop.
		endAddr := 0 // save the end of the last edit.
		oldLen := len(f.b)
		for i:=0; i<len(idxs); i++ {
			// Set the current dot to the i'th match.
			f.dot.from = idxs[i][0]
			f.dot.to = idxs[i][1]

			// The current dot of the step must be in the allowed address range
			if f.dot.from < endAddr {
				return insertAddress, newDots, errors.New("intersecting edits of x command")
			}

			// Execute an edit step.
			f.pos = commandPos
			var iA Address
			var nD []Address
			iA, nD, err = f.execCommand()
			if err != nil {
				return insertAddress, newDots, err
			}
			shift := len(f.b) - oldLen
			oldLen = len(f.b)

			// Check if the result of the edit overlaps.
			if iA.from < endAddr {
				return insertAddress, newDots, errors.New("intersecting edits of x command")
			}

			// Include the change to the total change address.
			if i==0 {
				insertAddress.from = iA.from
			}
			insertAddress.to = iA.to

			// Shift all remaining addresses.
			for k := i+1; k < len(idxs); k++ {
				idxs[k][0] += shift
				idxs[k][1] += shift
			}

			// Calculate new line addresses.
			f.lineAddresses = GetLineAddresses(f.b)

			// The new dots are the new dots from the last step.
			newDots = nD
		}
		return insertAddress, newDots, nil
	case "X": // loop and mark. This is used to return multiple newDots.
		f.pos++
		t = f.currentToken()
		if t.id != cRE {
			return insertAddress, newDots, errors.New("command X must be followed by a regexp: "+t.s)
		}
		idxs := t.re.FindAllIndex(f.b[f.dot.from:f.dot.to], -1)
		if idxs == nil {
			insertAddress = f.dot
			newDots = append(newDots, f.dot)
		} else {
			for i:=0; i<len(idxs); i++ {
				if i == 0 {
					insertAddress.from = f.dot.from + idxs[i][0]
				}
				insertAddress.to = f.dot.from + idxs[i][1]
				newDots = append(newDots, Address{from:idxs[i][0]+f.dot.from, to:idxs[i][1]+f.dot.from})
			}
		}
		return insertAddress, newDots, nil
	default:
		return insertAddress, newDots, errors.New("unknown command: "+t.s)
	}
	return insertAddress, newDots, errors.New("unreachable")
}

// CurrentToken returns the current token.
func (f *Samfile) currentToken() (t token) {
	if f.pos == len(f.tokens) {
		t.id = cEND
		return t
	}
	return f.tokens[f.pos]
}

// Insert new content to sam buffer.
func (f *Samfile) insertContent(insertAddress Address, newContent []byte) {
	shift := insertAddress.to - insertAddress.from + len(newContent)
	newbuffer := make([]byte, 0, len(f.b)+shift)
	newbuffer = append(newbuffer, f.b[0:insertAddress.from]...)
	newbuffer = append(newbuffer, newContent...)
	newbuffer = append(newbuffer, f.b[insertAddress.to:]...)
	f.b = newbuffer
}
