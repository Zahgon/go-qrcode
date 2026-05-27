package qrcode

type EncodeOption interface {
	apply(option *encodingOption)
}

// DefaultEncodingOption with EncMode = EncModeAuto, EcLevel = ErrorCorrectionQuart
func DefaultEncodingOption() *encodingOption { _ = "STUB: not implemented"; return nil }

type encodingOption struct {
	// Version of target QR code.
	Version int

	// MinimumVersion specifies the minimum version of target QR code.
	// If the automatically analyzed version is lower than this, this minimum will be used.
	MinimumVersion int

	// EncMode specifies which encMode to use
	EncMode encMode

	// EcLevel specifies which ecLevel to use
	EcLevel ecLevel

	// PS: The version (which implicitly defines the byte capacity of the qrcode) is dynamically selected at runtime
}

type fnEncodingOption struct {
	fn func(*encodingOption)
}

func (f fnEncodingOption) apply(option *encodingOption) { _ = "STUB: not implemented"; return }

func newFnEncodingOption(fn func(*encodingOption)) fnEncodingOption {
	_ = "STUB: not implemented"
	return *new(fnEncodingOption)
}

// WithEncodingMode sets the encoding mode.
func WithEncodingMode(mode encMode) EncodeOption {
	_ = "STUB: not implemented"
	return *new(EncodeOption)
}

// WithErrorCorrectionLevel sets the error correction level.
func WithErrorCorrectionLevel(ecLevel ecLevel) EncodeOption {
	_ = "STUB: not implemented"
	return *new(EncodeOption)
}

// WithVersion sets the version of target QR code.
func WithVersion(version int) EncodeOption { _ = "STUB: not implemented"; return *new(EncodeOption) }

// WithMinimumVersion sets the minimum version of target QR code.
// If the automatically analyzed version is lower than this minimum,
// the minimum version will be used instead.
func WithMinimumVersion(version int) EncodeOption {
	_ = "STUB: not implemented"
	return *new(EncodeOption)
}
