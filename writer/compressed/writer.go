package compressed

import (
	"image/color"
	"io"

	"github.com/yeqown/go-qrcode/v2"
)

type Option struct {
	Padding   int
	BlockSize int
}

// compressedWriter implements issue#69, generating compressed images
// in some special situations, such as, network transferring.
// https://github.com/yeqown/go-qrcode/issues/69
type compressedWriter struct {
	fd io.WriteCloser

	option *Option
}

var (
	backgroundColor = color.Gray{Y: 0xff}
	foregroundColor = color.Gray{Y: 0x00}
)

func New(filename string, opt *Option) (qrcode.Writer, error) {
	_ = "STUB: not implemented"
	return *new(qrcode.Writer), nil
}

func NewWithWriter(writeCloser io.WriteCloser, opt *Option) qrcode.Writer {
	_ = "STUB: not implemented"
	return *new(qrcode.Writer)
}

func (w compressedWriter) Write(mat qrcode.Matrix) error { _ = "STUB: not implemented"; return nil }

// background

//switch v.IsSet() {
//case false:
//	gray = backgroundColor
//default:
//	gray = foregroundColor
//}

func (w compressedWriter) Close() error { _ = "STUB: not implemented"; return nil }
