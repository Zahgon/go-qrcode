package shapes

import "github.com/yeqown/go-qrcode/writer/standard"

// ComposableShape implements the standard.IShape interface by delegating
// drawing behavior to externally supplied functions.
//
// This type enables flexible composition of shape logic, allowing clients
// to inject custom behavior for both the main shape (`Draw`) and its finder (`DrawFinder`).
type ComposableShape struct {
	onDrawFinder func(ctx *standard.DrawContext)
	onDraw       func(ctx *standard.DrawContext)
}

// Draw executes the injected draw function to render the shape body.
func (s *ComposableShape) Draw(ctx *standard.DrawContext) {
	_ = "STUB: not implemented"

	// DrawFinder executes the injected drawFinder function to render the shape's finder pattern.
	return
}

func (s *ComposableShape) DrawFinder(ctx *standard.DrawContext) { _ = "STUB: not implemented"; return }

// Assemble creates a new ComposableShape instance by assigning provided drawing
// functions. This allows dynamic, reusable construction of shape behaviors.
func Assemble(drawFinder, drawBlock func(ctx *standard.DrawContext)) standard.IShape {
	_ = "STUB: not implemented"
	return *new(standard.IShape)
}

// ----------- helpers -----------

func has(mask, bits uint16) bool { _ = "STUB: not implemented"; return false }
