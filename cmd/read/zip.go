package main

// Zip changes the Name to "/path/to/file.zip/" and forwards
// the request to the zip command.
func (p *program) zip(file, addr string) error {
	p.Name += "/"
	return p.Forward("zip", nil)
}
