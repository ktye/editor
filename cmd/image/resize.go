package main

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"golang.org/x/image/draw"
)

// Resize returns a resized image.
func (p *program) resize(args []string) error {
	var width, height int
	var err error
	if len(args) == 0 {
		args = []string{"50%"}
	} else if len(args) == 2 {
		width, err = strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("resize: wrong arguments: %v", args)
		}
		height, err = strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("resize: wrong arguments: %v", args)
		}
	} else if len(args) > 2 {
		return fmt.Errorf("resize: too many arguments")
	}

	var im image.Image
	im, err = p.decode()
	if err != nil {
		return err
	}

	if len(args) == 1 {
		// A single argument with a percent sign scales the image by this percentage.
		var fac float64
		if strings.HasSuffix(args[0], "%") {
			pct, err := strconv.Atoi(args[0][:len(args[0])-1])
			if err != nil {
				return fmt.Errorf("resize: wrong percentage argument")
			}
			fac = float64(pct) / 100.0
		} else {
			// A single integer argument is the new major size.
			size, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("resize: wrong argument: %s", args[0])
			}
			if im.Bounds().Dx() > im.Bounds().Dy() {
				fac = float64(size) / float64(im.Bounds().Dx())
			} else {
				fac = float64(size) / float64(im.Bounds().Dy())
			}
		}
		width = int(float64(im.Bounds().Dx()) * fac)
		height = int(float64(im.Bounds().Dy()) * fac)
	}

	scaled := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.ApproxBiLinear.Scale(scaled, scaled.Bounds(), im, im.Bounds(), draw.Src, nil)
	return p.send(scaled)
}
