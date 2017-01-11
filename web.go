package main

import (
	"bytes"
	"fmt"
)

// webRead embeds the web content
func webRead(w *bytes.Buffer, file string) (ok bool) {
	fmt.Fprintln(w, `<object class="embeddedpage" data="`+file+`"/>`)
	return true
}
