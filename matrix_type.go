package qrcode

type QRType = qrtype

// qrtype
type qrtype uint8

const (
	// QRType_INIT represents the initial block state of the matrix
	QRType_INIT qrtype = 1 << 1
	// QRType_DATA represents the data block state of the matrix
	QRType_DATA qrtype = 2 << 1
	// QRType_VERSION indicates the version block of matrix
	QRType_VERSION qrtype = 3 << 1
	// QRType_FORMAT indicates the format block of matrix
	QRType_FORMAT qrtype = 4 << 1
	// QRType_FINDER indicates the finder block of matrix
	QRType_FINDER qrtype = 5 << 1
	// QRType_DARK ...
	QRType_DARK     qrtype = 6 << 1
	QRType_SPLITTER qrtype = 7 << 1
	QRType_TIMING   qrtype = 8 << 1
)

func (s qrtype) String() string { _ = "STUB: not implemented"; return "" }

type QRValue = qrvalue

func (v QRValue) Type() qrtype { _ = "STUB: not implemented"; return *new(qrtype) }

func (v QRValue) IsSet() bool {
	_ = "STUB: not implemented"

	// qrvalue represents the value of the matrix, it is composed of the qrtype(7bits) and the value(1bits).
	// such as: 0b0000,0011 (QRValue_DATA_V1) represents the qrtype is QRType_DATA and the value is 1.
	return false
}

type qrvalue uint8

var (
	// QRValue_INIT_V0 represents the value 0 qrvalue(QRType_INIT | 0)
	QRValue_INIT_V0 = qrvalue(QRType_INIT)

	// QRValue_DATA_V0 represents the block has been set to false qrvalue(QRType_DATA | 0)
	QRValue_DATA_V0 = qrvalue(QRType_DATA)
	// QRValue_DATA_V1 represents the block has been set to TRUE
	QRValue_DATA_V1 = qrvalue(QRType_DATA | 1)

	// QRValue_VERSION_V0 represents the block has been set to false qrvalue(QRType_VERSION | 0)
	QRValue_VERSION_V0 = qrvalue(QRType_VERSION)
	// QRValue_VERSION_V1 represents the block has been set to TRUE
	QRValue_VERSION_V1 = qrvalue(QRType_VERSION | 1)

	// QRValue_FORMAT_V0 represents the block has been set to false qrvalue(QRType_FORMAT | 0)
	QRValue_FORMAT_V0 = qrvalue(QRType_FORMAT)
	// QRValue_FORMAT_V1 represents the block has been set to TRUE
	QRValue_FORMAT_V1 = qrvalue(QRType_FORMAT | 1)

	// QRValue_FINDER_V0 represents the block has been set to false qrvalue(QRType_FINDER | 0)
	QRValue_FINDER_V0 = qrvalue(QRType_FINDER)
	// QRValue_FINDER_V1 represents the block has been set to TRUE
	QRValue_FINDER_V1 = qrvalue(QRType_FINDER | 1)

	// QRValue_DARK_V0 represents the block has been set to false qrvalue(QRType_DARK | 0)
	QRValue_DARK_V0 = qrvalue(QRType_DARK)
	// QRValue_DARK_V1 represents the block has been set to TRUE
	QRValue_DARK_V1 = qrvalue(QRType_DARK | 1)

	// QRValue_SPLITTER_V0 represents the block has been set to false qrvalue(QRType_SPLITTER | 0)
	QRValue_SPLITTER_V0 = qrvalue(QRType_SPLITTER)
	// QRValue_SPLITTER_V1 represents the block has been set to TRUE
	QRValue_SPLITTER_V1 = qrvalue(QRType_SPLITTER | 1)

	// QRValue_TIMING_V0 represents the block has been set to false qrvalue(QRType_TIMING | 0)
	QRValue_TIMING_V0 = qrvalue(QRType_TIMING)
	// QRValue_TIMING_V1 represents the block has been set to TRUE
	QRValue_TIMING_V1 = qrvalue(QRType_TIMING | 1)
)

func (v qrvalue) qrtype() qrtype { _ = "STUB: not implemented"; return *new(qrtype) }

func (v qrvalue) qrbool() bool { _ = "STUB: not implemented"; return false }

func (v qrvalue) String() string { _ = "STUB: not implemented"; return "" }

func (v qrvalue) xor(v2 qrvalue) qrvalue { _ = "STUB: not implemented"; return *new(qrvalue) }
