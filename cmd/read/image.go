package main

// Image forwards the request to the image command.
func (p *program) image(file, addr string) error {
	return p.Forward("image", nil)
}
