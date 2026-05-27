//go:build js && wasm

package main

import (
	"bytes"
	"io"
	"syscall/js"

	"github.com/yeqown/go-qrcode/v2"
	stdw "github.com/yeqown/go-qrcode/writer/standard"
)

// encodeOption refers to github.com/yeqown/go-qrcode/v2.encodingOption
type encodeOption struct {
	version int    // encodeVersion
	mode    uint8  // encodeMode specifies which encMode to use (0/1/2/3).
	ecLevel string // encodeECLevel specifies which ecLevel to use (L/M/Q/H)
}

// outputOption refers to github.com/yeqown/go-qrcode/writer/standard.outputImageOptions
type outputOption struct {
	bgColor       string // outputBgColor is the background color of the QR code image.
	bgTransparent bool   // outputBgTransparent indicates whether the background color is transparent.
	qrColor       string // outputQrColor is the foreground color of the QR code.
	qrWidth       uint8  // outputQrWidth is the width of the QR code.
	circleShape   bool   // outputCircleShape indicates whether to draw the qr block in circle shape.
	imageEncoder  string // outputImageEncoder specifies file format would be encoded the QR image. (jpg/jpeg/png)
	margin        int    // outputMargin is the border width of the output image.
}

// genOption is a type of option for generating code.
type genOption struct {
	encodeOption
	outputOption
}

// optionFromJSValue converts js.Value to genOption.
func optionFromJSValue(option js.Value) *genOption {
	_ = "STUB: not implemented"
	// fmt.Println(option.Type().String())
	return nil
}

// option must be an object, otherwise return default empty option.

// validate genOption.
func (o *genOption) validate() error { _ = "STUB: not implemented"; return nil }

func (o *genOption) encodeOptions() []qrcode.EncodeOption { _ = "STUB: not implemented"; return nil }

func (o *genOption) outputOptions() []stdw.ImageOption { _ = "STUB: not implemented"; return nil }

// genResult is result contains generated code image or error message.
type genResult struct {
	Success            bool   `json:"success"`
	Error              string `json:"error"`
	Base64EncodedImage string `json:"base64EncodedImage"`
}

func (r *genResult) setError(err error) { _ = "STUB: not implemented"; return }

func (r *genResult) setImage(buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

func (r *genResult) JSValue() js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error { _ = "STUB: not implemented"; return nil }
