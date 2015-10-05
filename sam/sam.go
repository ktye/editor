package sam

// See doc.go for package documentation.

import (
	"errors"
	"strconv"
)

// Edit applies the sam command cmd to the byte array b
// and returns the modified b and the new selection addr
//
// The init dot is a list of absolute address range in the form
//	LINE:CHAR,LINE:CHAR;LINE:CHAR,LINE:CHAR...
// If it is empty, it contains the whole file in a single range
func Edit(b []byte, cmd string, initDot string) (out []byte, addr string, err error) {

	var f Samfile
	var initDots []Address
	var finalDots []Address

	// Initialize the sam file
	f, initDots, err = InitSamFile(b, cmd, initDot)
	if err != nil {
		return out, addr, err
	}

	rangeErr := errors.New("intersecting edits")
	var endAddr int = 0
	var oldLen int = len(f.b)
	for i:=0; i<len(initDots); i++ {

		// Set the next initial dot.
		f.dot = initDots[i]

		// The next dot must not start within the section of the previous change.
		if f.dot.from < endAddr {
			return out, addr, rangeErr
		}

		// Execute the command. The newDots are within the insertAddress.
		insertAddress, newDots, err := f.exec()
		if err != nil {
			return out, addr, err
		}
		shift := len(f.b) - oldLen
		oldLen = len(f.b)

		// Check if the new inserted text intersects with a previous edit.
		if insertAddress.from < endAddr {
			return out, addr, rangeErr
		}
		endAddr = insertAddress.to

		// Add new Dots to the final dots.
		finalDots = append(finalDots, newDots...)

		// Shift all following initial dots.
		for k:=i+1; k<len(initDots); k++ {
			initDots[k].from += shift
			initDots[k].to += shift
		}

		// Calculate new line addresses.
		f.lineAddresses = GetLineAddresses(f.b)
	}

	addr = ""
	for i := 0; i < len(finalDots); i++ {
		if i > 0 {
			addr += ";"
		}
		line, ch := LineAddr(finalDots[i].from, true, f.lineAddresses, f.b)
		addr += strconv.Itoa(line) + ":" + strconv.Itoa(ch)
		line, ch = LineAddr(finalDots[i].to, false, f.lineAddresses, f.b)
		addr += "," + strconv.Itoa(line) + ":" + strconv.Itoa(ch)
	}

	return f.b, addr, nil
}
