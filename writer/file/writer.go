package file

import (
	"os"

	"github.com/yeqown/go-qrcode/v2"
)

const (
	upRune     = 9600 // '▀'
	downRune   = 9604 // '▄'
	upDownRune = 9608 // '█'
	spaceRune  = 32   // ' '
)

var _ qrcode.Writer = (*Writer)(nil)

// Writer implements qrcode.Writer.
type Writer struct {
	out *os.File
}

// Close method to implement qrcode.Writer.
func (a *Writer) Close() error {
	_ = "STUB: not implemented"

	// Write method to implement qrcode.Writer.
	return nil
}

func (a *Writer) Write(mat qrcode.Matrix) error { _ = "STUB: not implemented"; return nil }

func New(f *os.File) *Writer { _ = "STUB: not implemented"; return nil }
