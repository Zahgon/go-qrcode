package imgkit

import (
	"image"
)

// Read reads an image from a file. only support PNG and JPEG yet.
func Read(path string) (img image.Image, err error) {
	_ = "STUB: not implemented"
	return *new(image.Image), nil
}

// Save saves the image to the given path.
func Save(img image.Image, filename string) error { _ = "STUB: not implemented"; return nil }
