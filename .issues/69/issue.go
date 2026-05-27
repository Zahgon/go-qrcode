/*
 * Link:	https://github.com/yeqown/go-qrcode/issues/69
 * Title:	Feature(image-compression): PNG bit depth must be 1
 * Author:  stokito(https://github.com/stokito)
 */

package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"strings"
)

/*
See https://github.com/yeqown/go-qrcode/issues/69
Results: (unit: B)
Content length:     158   // source text length
qr-skip2-best.png:  641   // skip2 generated size
qr-yeqown-best.png: 958  // yeqown best compression level
*/
func main() {
	// some json
	content := "{\n  \"id\": \"beb6c04733d1a2da0f87c910a384\",\n  \"tmax\": 895,\n  \"cur\": [\n    \"USD\"\n  ],\n  \"imp\": [\n    {\n      \"id\": \"21292\",\n      \"instl\": 0,\n      \"secure\": 0}\n"
	// strip json to alpha num just for a test
	content = strings.Replace(content, "{", "$", -1)
	content = strings.Replace(content, "}", "%", -1)
	content = strings.Replace(content, "\n", " ", -1)
	content = strings.Replace(content, "\"", "*", -1)
	content = strings.Replace(content, ",", "/", -1)
	content = strings.Replace(content, "[", "/", -1)
	content = strings.Replace(content, "]", "/", -1)
	content = strings.ToUpper(content)
	fmt.Printf("Content lenght: %d\n", len(content))
	//imgEncoderNone := CustomPngEncoder{png.NoCompression}
	imgEncoderBest := CustomPngEncoder{png.BestCompression}

	encodeWithYeqown(content, "qr-yeqown.png")
	//encodeWithYeqownCompression(content, "qr-yeqown-none.png", imgEncoderNone)
	encodeWithYeqownCompression(content, "qr-yeqown-best.png", imgEncoderBest)
	encodeWithSkip2(content, "qr-skip2-best.png")
}

func encodeWithSkip2(content, name string) {
	_ = "STUB: not implemented"
	// err := skip2.WriteFile(content, skip2.Highest, 0, name)
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	return
}

type CustomPngEncoder struct {
	CompressionLevel png.CompressionLevel
}

func (j CustomPngEncoder) Encode(w io.Writer, img image.Image) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeWithYeqown(content, name string) { _ = "STUB: not implemented"; return }

func encodeWithYeqownCompression(content, name string, imageEncoder CustomPngEncoder) {
	_ = "STUB: not implemented"
	//cfg := yeqown.DefaultConfig()
	//cfg.EncMode = yeqown.EncModeAlphanumeric
	////cfg.EncMode = yeqown.EncModeByte
	//cfg.EcLevel = yeqown.ErrorCorrectionHighest
	//
	//imgOpts := yeqown.WithCustomImageEncoder(imageEncoder)
	//imgOpts2 := yeqown.WithQRWidth(1)
	//imgOpts3 := yeqown.WithBorderWidth(4)
	//qrc, err := yeqown.NewWithConfig(content, cfg, imgOpts, imgOpts2, imgOpts3)
	//if err != nil {
	//	panic(err)
	//}
	return
}
