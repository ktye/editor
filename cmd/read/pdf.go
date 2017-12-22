package main

import (
	"encoding/base64"
	"io/ioutil"
)

// Pdf returns pdf files as html objects.
// The file content is encoded in the data part.
func (p *program) pdf(file, addr string) error {
	data := "data:application/pdf;base64,"
	if b, err := ioutil.ReadFile(file); err != nil {
		return err
	} else {
		data += base64.StdEncoding.EncodeToString(b)
	}

	p.Type = "html"
	p.Text = `<object data=` + data + ` type="application/pdf" width="100%" height="100%"></object>`
	return nil
}
