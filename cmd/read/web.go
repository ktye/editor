package main

// Web returns an html object with the url in the Name field.
func (p *program) web() error {
	p.Clean = true
	p.Type = "html"
	p.Default = ""
	p.Tags = ""
	p.Text = `<object class="embeddedpage" data="` + p.Name + `" />`
	return nil
}
