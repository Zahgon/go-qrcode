package qrcode

import (
	"errors"
)

var (
	// ErrorOutRangeOfW x out of range of Width
	ErrorOutRangeOfW = errors.New("out of range of width")

	// ErrorOutRangeOfH y out of range of Height
	ErrorOutRangeOfH = errors.New("out of range of height")
)

// newMatrix generate a matrix with map[][]qrbool
func newMatrix(width, height int) *Matrix { _ = "STUB: not implemented"; return nil }

// Matrix is a matrix data type
// width:3 height: 4 for [3][4]int
type Matrix struct {
	mat    [][]qrvalue
	width  int
	height int
}

// do some init work
func (m *Matrix) init() {
	for w := 0; w < m.width; w++ {
		for h := 0; h < m.height; h++ {
			m.mat[w][h] = QRValue_INIT_V0
		}
	}
}

// print to stdout
func (m *Matrix) print() { _ = "STUB: not implemented"; return }

// Copy matrix into a new Matrix
func (m *Matrix) Copy() *Matrix { _ = "STUB: not implemented"; return nil }

// Width ... width
func (m *Matrix) Width() int {
	_ = "STUB: not implemented"

	// Height ... height
	return 0
}

func (m *Matrix) Height() int {
	_ = "STUB: not implemented"

	// set [w][h] as true
	return 0
}

func (m *Matrix) set(w, h int, c qrvalue) error { _ = "STUB: not implemented"; return nil }

// at state qrvalue from matrix with position {x, y}
func (m *Matrix) at(w, h int) (qrvalue, error) {
	_ = "STUB: not implemented"
	return *new(qrvalue), nil
}

// iterDirection scan matrix direction
type iterDirection uint8

const (
	// IterDirection_ROW for row first
	IterDirection_ROW iterDirection = iota + 1

	// IterDirection_COLUMN for column first
	IterDirection_COLUMN
)

// Iterate the Matrix with loop direction IterDirection_ROW major or IterDirection_COLUMN major.
// IterDirection_COLUMN is recommended.
func (m *Matrix) Iterate(direction iterDirection, fn func(x, y int, s QRValue)) {
	_ = "STUB: not implemented"
	return
}

func (m *Matrix) iter(dir iterDirection, visitFn func(x int, y int, v qrvalue)) {
	_ = "STUB: not implemented"
	// row direction first
	return
}

// column direction first

// Row return a row of matrix, cur should be y dimension.
func (m *Matrix) Row(cur int) []qrvalue { _ = "STUB: not implemented"; return nil }

// Col return a slice of column, cur should be x dimension.
func (m *Matrix) Col(cur int) []qrvalue { _ = "STUB: not implemented"; return nil }

// Bitmap outputs the QR Code as a matrix of pixels, each represented by a single bit.
func (m *Matrix) Bitmap() [][]bool { _ = "STUB: not implemented"; return nil }
