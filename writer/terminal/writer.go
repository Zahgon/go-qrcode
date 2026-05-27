package terminal

import (
	"github.com/yeqown/go-qrcode/v2"

	termbox "github.com/nsf/termbox-go"
)

var _ qrcode.Writer = (*Writer)(nil)

// Writer implements qrcode.Writer based on termbox to print QRCode into
// terminal / console.
type Writer struct{}

func New() *Writer { _ = "STUB: not implemented"; return nil }

func (w Writer) init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	termbox.SetInputMode(termbox.InputEsc)
	termbox.SetOutputMode(termbox.Output256)
}

func (w Writer) preDraw(width, height, padding int, bg termbox.Attribute) {
	_ = "STUB: not implemented"
	return
}

// drawBlock draws a block at (x, y) with fg and bg colors.
// each block takes 2 times width of one character terminal, it looks like: ██
func (w Writer) drawBlock(x, y, padding int, fg termbox.Attribute, bg termbox.Attribute) {
	_ = "STUB: not implemented"
	return
}

func (w Writer) Write(mat qrcode.Matrix) error {
	_ = "STUB: not implemented"
	// width, height, whratio := terminalSize()
	// _ = width
	// _ = height
	// _ = whratio
	return nil
}

func printTip(y int) { _ = "STUB: not implemented"; return }

func hold() error { _ = "STUB: not implemented"; return nil }

func (w Writer) Close() error { _ = "STUB: not implemented"; return nil }
