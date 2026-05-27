package qrcode

// evaluation calculate a score after masking matrix.
//
// reference:
// - https://www.thonky.com/qr-code-tutorial/data-masking#Determining-the-Best-Mask
func evaluation(mat *Matrix) (score int) { _ = "STUB: not implemented"; return 0 }

// check each row one-by-one. If there are five consecutive modules of the same color,
// add 3 to the penalty. If there are more modules of the same color after the first five,
// add 1 for each additional module of the same color. Afterward, check each column one-by-one,
// checking for the same condition. Add the horizontal and vertical total to obtain penalty score
func rule1(mat *Matrix) (score int) {
	_ = "STUB: not implemented"
	// prerequisites:
	// mat.Width() == mat.Height()
	return 0
}

// rule2
// look for areas of the same color that are at least 2x2 modules or larger.
// The QR code specification says that for a solid-color block of size m × n,
// the penalty score is 3 × (m - 1) × (n - 1).
func rule2(mat *Matrix) int { _ = "STUB: not implemented"; return 0 }

// rule3 calculate punishment score in rule3, find pattern in QR Code matrix.
// Looks for patterns of dark-light-dark-dark-dark-light-dark that have four
// light modules on either side. In other words, it looks for any of the
// following two patterns: 1011101 0000 or 0000 1011101.
//
// Each time this pattern is found, add 40 to the penalty score.
func rule3(mat *Matrix) (score int) { _ = "STUB: not implemented"; return 0 }

// prerequisites:
//
// mat.Width() == mat.Height()

// DONE(@yeqown): statePattern1 and statePattern2 are fixed, so maybe kmpGetNext
// could cache result to speed up.

// rule4 is based on the ratio of light modules to dark modules:
//
// 1. Count the total number of modules in the matrix.
// 2. Count how many dark modules there are in the matrix.
// 3. Calculate the percent of modules in the matrix that are dark: (darkmodules / totalmodules) * 100
// 4. Determine the previous and next multiple of five of this percent.
// 5. Subtract 50 from each of these multiples of five and take the absolute qrbool of the result.
// 6. Divide each of these by five. For example, 10/5 = 2 and 5/5 = 1.
// 7. Finally, take the smallest of the two numbers and multiply it by 10.
func rule4(mat *Matrix) int {
	_ = "STUB: not implemented"
	// prerequisites:
	//
	// mat.Width() == mat.Height()
	return 0
}

// count dark modules

// in range [0, 100]
