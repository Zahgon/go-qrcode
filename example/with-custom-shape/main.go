package main

import (
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type smallerCircle struct {
	smallerPercent float64
}

func (sc *smallerCircle) DrawFinder(ctx *standard.DrawContext) { _ = "STUB: not implemented"; return }

func newShape(radiusPercent float64) standard.IShape {
	_ = "STUB: not implemented"
	return *new(standard.IShape)
}

func (sc *smallerCircle) Draw(ctx *standard.DrawContext) { _ = "STUB: not implemented"; return }

// choose a proper radius values

// 80 percent smaller

// get center point

func main() {
	shape := newShape(0.7)
	qrc, err := qrcode.New("with-custom-shape")
	// qrc, err := qrcode.New("with-custom-shape", qrcode.WithCircleShape())
	if err != nil {
		panic(err)
	}

	w, err := standard.New("./smaller.png", standard.WithCustomShape(shape))
	if err != nil {
		panic(err)
	}

	err = qrc.Save(w)
	if err != nil {
		panic(err)
	}
}
