package main

import (
	"fmt"
	"image"
)

// Crop decodes the image from the request's Text field,
// and returns a cropped image using the current Selections.
func (p *program) crop() error {
	if im, err := p.decode(); err != nil {
		return err
	} else {
		if len(p.Selections) != 2 {
			return fmt.Errorf("image has invalid selections")
		}
		xy := p.Selections[0]
		wh := p.Selections[1]

		if sub, ok := im.(subImager); ok {
			im = sub.SubImage(image.Rect(xy[0], xy[1], xy[0]+wh[0], xy[1]+wh[1]))
			return p.send(im)
		} else {
			return fmt.Errorf("image does not support SubImage")
		}
	}
}

type subImager interface {
	SubImage(r image.Rectangle) image.Image
}
