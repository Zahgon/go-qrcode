package standard

import (
	"image"
	"image/color"
)

// ColorStop represents a single color stop in a gradient.
// T defines the position along the gradient line, ranging from 0.0 to 1.0.
// Color defines the color at this stop.
type ColorStop struct {
	T     float64 // from 0.0 to 1.0.
	Color color.RGBA
}

// LinearGradient defines a linear gradient with angle and color stops.
// The gradient progresses in the direction of the given angle (in degrees).
// Angle is interpreted as: 0 - right, 90 - up, 180 - left, 270 - down.
type LinearGradient struct {
	Stops []ColorStop // Ordered list of color stops along the gradient
	Angle float64     // Gradient angle in degrees
}

// NewGradient creates a new LinearGradient with the specified angle (in degrees) and color stops.
// The stops are sorted in ascending order of T.
func NewGradient(angle float64, stops ...ColorStop) *LinearGradient {
	_ = "STUB: not implemented"
	return nil
}

// applyGradient applies the linear gradient to all pixels in the image that match the given foreground color.
func (g *LinearGradient) applyGradient(img image.Image, fgColor color.RGBA) *image.RGBA {
	_ = "STUB: not implemented"
	// Convert angle to radians and compute gradient direction vector
	return nil
}

// Get all 4 corners of the image

// Compute min and max projection of corners on the gradient axis

// Prepare output image

// Replace foreGround pixels with interpolated gradient colors

// Project pixel onto gradient axis

// Normalize to [0, 1]

// Set pixel color from gradient

// interpolateColor returns a color interpolated from the gradient stops based on position t.
func interpolateColor(stops []ColorStop, t float64) color.RGBA {
	_ = "STUB: not implemented"
	return *new(color.RGBA)
}

// Linear interpolation between two stops

// fallback (should not happen)

// blendColors returns the color interpolated between c1 and c2 using t in [0.0 - 1.0].
func blendColors(c1, c2 color.RGBA, t float64) color.RGBA {
	_ = "STUB: not implemented"
	return *new(color.RGBA)
}
