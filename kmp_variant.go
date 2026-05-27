package qrcode

// kmp is variant of kmp algorithm to count the pattern been in
// src slice.
// DONE(@yeqown): implement this in generic way.
func kmp[v comparable](src, pattern []v, next []int) (count int) {
	_ = "STUB: not implemented"
	return 0
}

// cursor of src
// cursor of pattern

// reset cursor to count duplicate pattern.
// such as: "aaaa" and "aa", we want 3 rather than 2.

func kmpGetNext[v comparable](pattern []v) []int { _ = "STUB: not implemented"; return nil }
