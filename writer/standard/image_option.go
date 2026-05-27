package standard

import (
	"image"
	"image/color"

	"github.com/yeqown/go-qrcode/v2"
)

type ImageOption interface {
	apply(o *outputImageOptions)
}

// defaultOutputImageOption default output image background color and etc options
func defaultOutputImageOption() *outputImageOptions { _ = "STUB: not implemented"; return nil }

// white
// not transparent
// black
//

//
//

// outputImageOptions to output QR code image
type outputImageOptions struct {
	// bgColor is the background color of the QR code image.
	bgColor color.RGBA
	// bgTransparent only affects on PNG_FORMAT
	bgTransparent bool

	// qrColor is the foreground color of the QR code.
	qrColor color.RGBA

	// qrGradient is an optional linear gradient to apply to QR modules instead of a solid color.
	qrGradient *LinearGradient

	// logo this icon image would be put the center of QR Code image
	// NOTE: logo only should have 1 / logoSizeMultiplier size of QRCode image
	logo image.Image

	logoSizeMultiplier int

	// logoSafeZone indicates whether to reserve a clear area around the logo,
	// preventing QR code blocks from being drawn underneath it.
	logoSafeZone bool

	// qrWidth width of each qr block
	qrWidth int

	// shape means how to draw the shape of each cell.
	shape IShape

	// imageEncoder specify which file format would be encoded the QR image.
	imageEncoder ImageEncoder

	// borderWidths indicates the border width of the output image. the order is
	// top, right, bottom, left same as the WithBorder
	borderWidths [4]int

	// halftoneImg is the halftone image for the output image.
	halftoneImg image.Image
}

func (oo *outputImageOptions) backgroundColor() color.RGBA {
	_ = "STUB: not implemented"
	return *new(color.RGBA)
}

func (oo *outputImageOptions) logoImage() image.Image {
	_ = "STUB: not implemented"
	return *new(image.Image)
}

func (oo *outputImageOptions) qrBlockWidth() int { _ = "STUB: not implemented"; return 0 }

func (oo *outputImageOptions) getShape() IShape { _ = "STUB: not implemented"; return *new(IShape) }

// preCalculateAttribute this function must reference to draw function.
func (oo *outputImageOptions) preCalculateAttribute(dimension int) *Attribute {
	_ = "STUB: not implemented"
	return nil
}

var (
	color_WHITE = parseFromHex("#ffffff")
	color_BLACK = parseFromHex("#000000")
)

var (
	// _STATE_MAPPING mapping matrix.State to color.RGBA in debug mode.
	_STATE_MAPPING = map[qrcode.QRType]color.RGBA{
		qrcode.QRType_INIT:     parseFromHex("#ffffff"), // [bg]
		qrcode.QRType_DATA:     parseFromHex("#cdc9c3"), // [bg]
		qrcode.QRType_VERSION:  parseFromHex("#000000"), // [fg]
		qrcode.QRType_FORMAT:   parseFromHex("#444444"), // [fg]
		qrcode.QRType_FINDER:   parseFromHex("#555555"), // [fg]
		qrcode.QRType_DARK:     parseFromHex("#2BA859"), // [fg]
		qrcode.QRType_SPLITTER: parseFromHex("#2BA859"), // [fg]
		qrcode.QRType_TIMING:   parseFromHex("#000000"), // [fg]
	}
)

// translateToRGBA get color.RGBA by value State, if not found, return outputImageOptions.qrColor.
// NOTE: this function decides the state should use qrColor or bgColor.
func (oo *outputImageOptions) translateToRGBA(v qrcode.QRValue) (rgba color.RGBA) {
	_ = "STUB: not implemented"
	// TODO(@yeqown): use _STATE_MAPPING to replace this function while in debug mode
	// or some special flag.
	return *new(color.RGBA)
}

// color.RGBA is pre-multiplied by alpha, so set RGB to 0 when fully transparent.

// parseFromHex convert hex string into color.RGBA
func parseFromHex(s string) color.RGBA { _ = "STUB: not implemented"; return *new(color.RGBA) }

// Double the hex digits:

func parseFromColor(c color.Color) color.RGBA { _ = "STUB: not implemented"; return *new(color.RGBA) }
