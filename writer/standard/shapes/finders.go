package shapes

import "github.com/yeqown/go-qrcode/writer/standard"

// RoundedFinder returns a function that renders the QR code's finder pattern
// with rounded transitions at the corners
func RoundedFinder() func(ctx *standard.DrawContext) { _ = "STUB: not implemented"; return nil }

// top right corners

// top left corners

// bot left corners

// bot right corners

// SquareFinder just square finder
func SquareFinder() func(ctx *standard.DrawContext) { _ = "STUB: not implemented"; return nil }
