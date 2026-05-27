package standard

import (
	"image/color"

	"github.com/fogleman/gg"
)

var (
	_shapeRectangle IShape = rectangle{}
	_shapeCircle    IShape = circle{}
)

type IShape interface {
	// Draw the shape of QRCode block in IShape implemented way.
	Draw(ctx *DrawContext)

	// DrawFinder to fill the finder pattern of QRCode, what's finder? google it for more information.
	DrawFinder(ctx *DrawContext)
}

// DrawContext is a rectangle area
type DrawContext struct {
	*gg.Context

	x, y float64
	w, h int

	color      color.Color
	neighbours uint16
}

// UpperLeft returns the point which indicates the upper left position.
func (dc *DrawContext) UpperLeft() (dx, dy float64) {
	_ = "STUB: not implemented"

	// Edge returns width and height of each shape could take at most.
	return 0, 0
}

func (dc *DrawContext) Edge() (width, height int) {
	_ = "STUB: not implemented"

	// Bit flags for the 8 surrounding cells in a 3x3 grid around the center (x, y).
	// Layout:
	// NTopLeft		NTop 	NTopRight
	// NLeft  		NSelf	NRight
	// NBotLeft 	NBot 	NBotRight
	return 0, 0
}

const (
	NTopLeft  uint16 = 1 << iota // top-left
	NTop                         // top
	NTopRight                    // top-right
	NLeft                        // left
	NSelf                        // center (self)
	NRight                       // right
	NBotLeft                     // bottom-left
	NBot                         // bottom
	NBotRight                    // bottom-right
)

// Neighbours returns a bitmask representing the neighboring blocks of the current block
func (dc *DrawContext) Neighbours() uint16 { _ = "STUB: not implemented"; return 0 }

// Color returns the color which should be fill into the shape. Note that if you're not
// using this color but your coded color.Color, some ImageOption functions those set foreground color
// would take no effect.
func (dc *DrawContext) Color() color.Color {
	_ = "STUB: not implemented"

	// rectangle IShape
	return *new(color.Color)
}

type rectangle struct{}

func (r rectangle) Draw(c *DrawContext) {
	_ = "STUB: not implemented"
	// FIXED(@yeqown): miss parameter of DrawRectangle
	return
}

func (r rectangle) DrawFinder(ctx *DrawContext) {
	_ = "STUB: not implemented"

	// circle IShape
	return
}

type circle struct{}

// Draw
// FIXED: Draw could not draw circle
func (r circle) Draw(c *DrawContext) {
	_ = "STUB: not implemented"
	// choose a proper radius values
	return
}

// get center point

func (r circle) DrawFinder(ctx *DrawContext) { _ = "STUB: not implemented"; return }
