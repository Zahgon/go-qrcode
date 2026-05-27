package imgkit

import (
	"image"

	"golang.org/x/image/draw"
)

// Binaryzation process image with threshold value (0-255) and return new image.
func Binaryzation(src image.Image, threshold uint8) image.Image {
	_ = "STUB: not implemented"
	return *new(image.Image)
}

// var rgb int = int(gray[i][j][0]) + int(gray[i][j][1]) + int(gray[i][j][2])

func Gray(src image.Image) *image.Gray { _ = "STUB: not implemented"; return nil }

func Scale(src image.Image, rect image.Rectangle, scale draw.Scaler) image.Image {
	_ = "STUB: not implemented"
	return *new(image.Image)
}
