package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"strings"
)

// diffFile shows the differences between the given file on disk and the post body.
func diffFile(w *bytes.Buffer, file string, body io.ReadCloser, rootDir string) (ok bool) {

	// read disk file.
	diskFile, err := ioutil.ReadFile(rootDir + file)
	if err != nil {
		fmt.Fprintln(w, err)
		return false
	}
	diskLines := strings.Split(string(diskFile), "\n")

	// read editor file.
	var editorFile []byte
	editorFile, err = ioutil.ReadAll(body)
	if err != nil {
		fmt.Fprintln(w, err)
		return false
	}
	editorLines := strings.Split(string(editorFile), "\n")

	// compute the diff.
	dr := diff(diskLines, editorLines)

	// if both files are identical, return the original file from disk and mark ok.
	if isDifferent(dr) == false {
		ok, _ = fileRead(w, file, rootDir)
		return ok
	}

	// if there are differences, write them and mark as error (to be opened in a new window).
	linesRemoved := 0
	linesAdded := 0
	for i := 0; i < len(dr); i++ {
		if dr[i].Delta == leftOnly {
			linesRemoved++
		} else if dr[i].Delta == rightOnly {
			linesAdded++
		}
	}

	fmt.Fprintf(w, "- on disk file   (%d lines removed)\n", linesRemoved)
	fmt.Fprintf(w, "+ in memory file (%d lines added)\n\n", linesAdded)
	for i := 0; i < len(dr); i++ {
		fmt.Fprintln(w, dr[i])
	}
	return false
}

/* diff library is taken from: https://raw.githubusercontent.com/aryann/difflib/master/difflib.go */
// Copyright 2012 Aryan Naraghi (aryan.naraghi@gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

type deltaType int

const (
	common deltaType = iota
	leftOnly
	rightOnly
)

// String returns a string representation for deltaType.
func (t deltaType) String() string {
	switch t {
	case common:
		return " "
	case leftOnly:
		return "-"
	case rightOnly:
		return "+"
	}
	return "?"
}

type diffRecord struct {
	Payload string
	Delta   deltaType
}

func isDifferent(d []diffRecord) bool {
	for i := 0; i < len(d); i++ {
		if d[i].Delta != common {
			return true
		}
	}
	return false
}

// String returns a string representation of d. The string is a
// concatenation of the delta type and the payload.
func (d diffRecord) String() string {
	return fmt.Sprintf("%s %s", d.Delta, d.Payload)
}

// Diff returns the result of diffing the seq1 and seq2.
func diff(seq1, seq2 []string) (diff []diffRecord) {
	// Trims any common elements at the heads and tails of the
	// sequences before running the diff algorithm. This is an
	// optimization.
	start, end := numEqualStartAndEndElements(seq1, seq2)

	for _, content := range seq1[:start] {
		diff = append(diff, diffRecord{content, common})
	}

	diffRes := compute(seq1[start:len(seq1)-end], seq2[start:len(seq2)-end])
	diff = append(diff, diffRes...)

	for _, content := range seq1[len(seq1)-end:] {
		diff = append(diff, diffRecord{content, common})
	}
	return
}

// numEqualStartAndEndElements returns the number of elements a and b
// have in common from the beginning and from the end. If a and b are
// equal, start will equal len(a) == len(b) and end will be zero.
func numEqualStartAndEndElements(seq1, seq2 []string) (start, end int) {
	for start < len(seq1) && start < len(seq2) && seq1[start] == seq2[start] {
		start++
	}
	i, j := len(seq1)-1, len(seq2)-1
	for i > start && j > start && seq1[i] == seq2[j] {
		i--
		j--
		end++
	}
	return
}

// intMatrix returns a 2-dimensional slice of ints with the given
// number of rows and columns.
func intMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

// longestCommonSubsequenceMatrix returns the table that results from
// applying the dynamic programming approach for finding the longest
// common subsequence of seq1 and seq2.
func longestCommonSubsequenceMatrix(seq1, seq2 []string) [][]int {
	matrix := intMatrix(len(seq1)+1, len(seq2)+1)
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if seq1[len(seq1)-i] == seq2[len(seq2)-j] {
				matrix[i][j] = matrix[i-1][j-1] + 1
			} else {
				matrix[i][j] = int(math.Max(float64(matrix[i-1][j]),
					float64(matrix[i][j-1])))
			}
		}
	}
	return matrix
}

// compute is the unexported helper for Diff that returns the results of
// diffing left and right.
func compute(seq1, seq2 []string) (diff []diffRecord) {
	matrix := longestCommonSubsequenceMatrix(seq1, seq2)
	i, j := len(seq1), len(seq2)
	for i > 0 || j > 0 {
		if i > 0 && matrix[i][j] == matrix[i-1][j] {
			diff = append(diff, diffRecord{seq1[len(seq1)-i], leftOnly})
			i--
		} else if j > 0 && matrix[i][j] == matrix[i][j-1] {
			diff = append(diff, diffRecord{seq2[len(seq2)-j], rightOnly})
			j--
		} else if i > 0 && j > 0 {
			diff = append(diff, diffRecord{seq1[len(seq1)-i], common})
			i--
			j--
		}
	}
	return
}
