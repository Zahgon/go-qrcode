//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	js.Global().Set("generateQRCode", js.FuncOf(genqrcode))
	fmt.Println("com.github.yeqown.goqrcode.wasm loaded")

	select {}
}

// genqrcode generates a qrcode image and returns the base64 encoded string.
// args should be a string array with length of 2 at most. the first one is the
// content string which will be encoded to qrcode, the second one is the encoding
// option, which is optional.
//
//	let result = generateQRCode("content", {
//	  qrWidth: 200,
//	  qrMargin: 10,
//	  qrColor: "#000000",
//	  qrBackColor: "#ffffff",
//	  encLevel: "H",
//	  encVersion: 7,
//	})
//
// more options refer to the `genOption` struct
func genqrcode(_ js.Value, args []js.Value) (v interface{}) { _ = "STUB: not implemented"; return nil }

//if err := opt.validate(); err != nil {
//	r.setError(errors.Wrap(err, "invalid option"))
//	return
//}

// apply image to result
