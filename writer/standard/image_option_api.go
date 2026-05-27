package standard

import (
	"image"
	"image/color"
)

// funcOption wraps a function that modifies outputImageOptions into an
// implementation of the ImageOption interface.
type funcOption struct {
	f func(oo *outputImageOptions)
}

func (fo *funcOption) apply(oo *outputImageOptions) { _ = "STUB: not implemented"; return }

func newFuncOption(f func(oo *outputImageOptions)) *funcOption {
	_ = "STUB: not implemented"
	return nil
}

// WithBgTransparent makes the background transparent.
func WithBgTransparent() ImageOption { _ = "STUB: not implemented"; return *new(ImageOption) }

// WithBgColor background color
func WithBgColor(c color.Color) ImageOption { _ = "STUB: not implemented"; return *new(ImageOption) }

// WithBgColorRGBHex background color
func WithBgColorRGBHex(hex string) ImageOption { _ = "STUB: not implemented"; return *new(ImageOption) }

// WithFgColor QR color
func WithFgColor(c color.Color) ImageOption { _ = "STUB: not implemented"; return *new(ImageOption) }

// WithFgColorRGBHex Hex string to set QR Color
func WithFgColorRGBHex(hex string) ImageOption { _ = "STUB: not implemented"; return *new(ImageOption) }

// WithFgGradient QR gradient
func WithFgGradient(g *LinearGradient) ImageOption {
	_ = "STUB: not implemented"
	return *new(ImageOption)
}

// WithLogoImage image should only has 1/5 width of QRCode at most
func WithLogoImage(img image.Image) ImageOption {
	_ = "STUB: not implemented"
	return *new(ImageOption)
}

// WithLogoImageFileJPEG load image from file, jpeg is required.
// image should only have 1/5 width of QRCode at most
func WithLogoImageFileJPEG(f string) ImageOption {
	_ = "STUB: not implemented"
	return *new(ImageOption)
}

// WithLogoImageFilePNG load image from file, PNG is required.
// image should only have 1/5 width of QRCode at most
func WithLogoImageFilePNG(f string) ImageOption {
	_ = "STUB: not implemented"
	return *new(ImageOption)
}

// WithQRWidth specify width of each qr block
func WithQRWidth(width uint8) ImageOption { _ = "STUB: not implemented"; return *new(ImageOption) }

// WithCircleShape use circle shape as rectangle(default)
func WithCircleShape() ImageOption { _ = "STUB: not implemented"; return *new(ImageOption) }

// WithCustomShape use custom shape as rectangle(default)
func WithCustomShape(shape IShape) ImageOption { _ = "STUB: not implemented"; return *new(ImageOption) }

// WithBuiltinImageEncoder option includes: JPEG_FORMAT as default, PNG_FORMAT.
// This works like WithBuiltinImageEncoder, the different between them is
// formatTyp is enumerated in (JPEG_FORMAT, PNG_FORMAT)
func WithBuiltinImageEncoder(format formatTyp) ImageOption {
	_ = "STUB: not implemented"
	return *new(ImageOption)
}

// WithCustomImageEncoder to use custom image encoder to encode image.Image into
// io.Writer
func WithCustomImageEncoder(encoder ImageEncoder) ImageOption {
	_ = "STUB: not implemented"
	return *new(ImageOption)
}

// WithBorderWidth specify the both 4 sides' border width. Notice that
// WithBorderWidth(a) means all border width use this variable `a`,
// WithBorderWidth(a, b) mean top/bottom equal to `a`, left/right equal to `b`.
// WithBorderWidth(a, b, c, d) mean top, right, bottom, left.
func WithBorderWidth(widths ...int) ImageOption {
	_ = "STUB: not implemented"
	return *new(ImageOption)
}

// 4+

// WithHalftone ...
func WithHalftone(path string) ImageOption { _ = "STUB: not implemented"; return *new(ImageOption) }

// WithHalftoneImage is identical to WithHalftone, but instead takes an image.Image as
// input. This is useful if the image will not be saved to disk.
func WithHalftoneImage(srcImg image.Image) ImageOption {
	_ = "STUB: not implemented"
	return *new(ImageOption)
}

// WithLogoSizeMultiplier used in Writer in validLogoImage method to validate logo size
func WithLogoSizeMultiplier(multiplier int) ImageOption {
	_ = "STUB: not implemented"
	return *new(ImageOption)
}

// WithLogoSafeZone enables the safe zone logic around the logo area.
func WithLogoSafeZone() ImageOption { _ = "STUB: not implemented"; return *new(ImageOption) }
