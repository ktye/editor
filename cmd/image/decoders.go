package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func init() {
	decoders = make(map[string]decoder)
	decoders[".jpg"] = new(jpegDecoder)
	decoders[".jpeg"] = new(jpegDecoder)
	decoders[".png"] = new(pngDecoder)
}

type decoder interface {
	Decode(io.Reader) (image.Image, error)
}

var decoders map[string]decoder

type jpegDecoder struct {
}

func (d jpegDecoder) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

type pngDecoder struct {
}

func (d pngDecoder) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}
