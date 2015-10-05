// Program sam is the stream editor version of sam.
// 
// It applies the arguments as sam commands to the
// standard input, one after another, while retaining
// the current dot.
// Initially the dot is set to the whole input.
package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"editor/sam"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	dot := ""
	for i:=1; i<len(os.Args); i++ {
		b, dot, err = sam.Edit(b, os.Args[i], dot)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	os.Stdout.Write(b)
}
