package main

import (
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"github.com/yeqown/go-qrcode/writer/standard/shapes"
)

func HChainBlock(ctx *standard.DrawContext) { _ = "STUB: not implemented"; return }

// todo:

func main() {
	// assemble qr injecting build in function or you own for drawing

	shape := shapes.Assemble(shapes.RoundedFinder(), shapes.LiquidBlock())
	//shape := shapes.Assemble(shapes.RoundedFinder(), HChainBlock)

	qrc, err := qrcode.New(`https://github.com/yeqown/go-qrcode`)
	if err != nil {
		panic(err)
	}

	w, err := standard.New("./smaller.png",
		standard.WithCustomShape(shape),
	)
	if err != nil {
		panic(err)
	}

	err = qrc.Save(w)
	if err != nil {
		panic(err)
	}
}
