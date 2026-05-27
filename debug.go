package qrcode

import (
	"io"
	"sync"
)

var (
	// _debug mode switch, true means enable debug mode, false means disable.
	_debug     = false
	_debugOnce sync.Once
)

func debugEnabled() bool {
	_ = "STUB: not implemented"
	// load debug switch from environment only once.
	return false
}

// SetDebugMode open debug switch, you can also enable debug by runtime
// environments variables: QRCODE_DEBUG=1 [1, true, TRUE, enabled, ENABLED] which is recommended.
func SetDebugMode() { _ = "STUB: not implemented"; return }

func debugLogf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func debugDraw(filename string, mat Matrix) error { _ = "STUB: not implemented"; return nil }

func debugDrawTo(w io.Writer, mat Matrix) error { _ = "STUB: not implemented"; return nil }

// width as image width, height as image height

// background

// choose color, false use black, others use black on white background

// save to writer
