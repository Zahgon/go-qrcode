package qrcode

// maskPatternModulo ...
// mask Pattern ref to: https://www.thonky.com/qr-code-tutorial/mask-patterns
type maskPatternModulo uint32

const (
	// modulo0 (x+y) mod 2 == 0
	modulo0 maskPatternModulo = iota
	// modulo1 (x) mod 2 == 0
	modulo1
	// modulo2 (y) mod 3 == 0
	modulo2
	// modulo3 (x+y) mod 3 == 0
	modulo3
	// modulo4 (floor (x/ 2) + floor (y/ 3) mod 2 == 0
	modulo4
	// modulo5 (x * y) mod 2) + (x * y) mod 3) == 0
	modulo5
	// modulo6 (x * y) mod 2) + (x * y) mod 3) mod 2 == 0
	modulo6
	// modulo7 (x + y) mod 2) + (x * y) mod 3) mod 2 == 0
	modulo7
)

type mask struct {
	mat      *Matrix           // matrix
	mode     maskPatternModulo // mode
	moduloFn moduloFunc        // moduloFn masking function
}

// newMask ...
func newMask(mat *Matrix, mode maskPatternModulo) *mask { _ = "STUB: not implemented"; return nil }

// moduloFunc to define what's modulo func
type moduloFunc func(int, int) bool

func getModuloFunc(mode maskPatternModulo) (f moduloFunc) {
	_ = "STUB: not implemented"
	return *new(moduloFunc)
}

// init generate maks by mode
func (m *mask) masking() { _ = "STUB: not implemented"; return }

// skip the function modules

// modulo0Func for maskPattern function
// modulo0 (x+y) mod 2 == 0
func modulo0Func(x, y int) bool { _ = "STUB: not implemented"; return false }

// modulo1Func for maskPattern function
// modulo1 (y) mod 2 == 0
func modulo1Func(x, y int) bool {
	_ = "STUB: not implemented"

	// modulo2Func for maskPattern function
	// modulo2 (x) mod 3 == 0
	return false
}

func modulo2Func(x, y int) bool {
	_ = "STUB: not implemented"

	// modulo3Func for maskPattern function
	// modulo3 (x+y) mod 3 == 0
	return false
}

func modulo3Func(x, y int) bool { _ = "STUB: not implemented"; return false }

// modulo4Func for maskPattern function
// modulo4 (floor (x/ 2) + floor (y/ 3) mod 2 == 0
func modulo4Func(x, y int) bool { _ = "STUB: not implemented"; return false }

// modulo5Func for maskPattern function
// modulo5 (x * y) mod 2 + (x * y) mod 3 == 0
func modulo5Func(x, y int) bool { _ = "STUB: not implemented"; return false }

// modulo6Func for maskPattern function
// modulo6 (x * y) mod 2) + (x * y) mod 3) mod 2 == 0
func modulo6Func(x, y int) bool { _ = "STUB: not implemented"; return false }

// modulo7Func for maskPattern function
// modulo7 (x + y) mod 2) + (x * y) mod 3) mod 2 == 0
func modulo7Func(x, y int) bool { _ = "STUB: not implemented"; return false }
