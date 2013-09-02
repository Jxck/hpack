package hpac

type Frame interface {
}

// Indexed Header Representation
//
// 	0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 1 |        Index (7+)         |
// +---+---------------------------+
type IndexedHeader struct {
	Flag1 uint8
	Index uint64
}

func NewIndexedHeader() (frame *IndexedHeader) {
	frame = &IndexedHeader{}
	frame.Flag1 = 1
	return
}

// Literal Header without Indexing - New Name
//
// 0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 0 | 1 | 1 |         0         |
// +---+---+---+-------------------+
// |       Name Length (8+)        |
// +-------------------------------+
// |  Name String (Length octets)  |
// +-------------------------------+
// |       Value Length (8+)       |
// +-------------------------------+
// | Value String (Length octets)  |
// +-------------------------------+
type NewNameWithoutIndexing struct {
	Flag1       uint8
	Flag2       uint8
	Flag3       uint8
	Index       uint64
	NameLength  uint64
	NameString  string
	ValueLength uint64
	ValueString string
}

func NewNewNameWithoutIndexing() (frame *NewNameWithoutIndexing) {
	frame = &NewNameWithoutIndexing{}
	frame.Flag1 = 0
	frame.Flag2 = 1
	frame.Flag3 = 1
	frame.Index = 0
	return
}

// Literal Header without Indexing - Indexed Name
//
//   0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 0 | 1 | 1 |    Index (5+)     |
// +---+---+---+-------------------+
// |       Value Length (8+)       |
// +-------------------------------+
// | Value String (Length octets)  |
// +-------------------------------+
type IndexedNameWithoutIndexing struct {
	Flag1       uint8
	Flag2       uint8
	Flag3       uint8
	Index       uint64
	ValueLength uint64
	ValueString string
}

func NewIndexedNameWithoutIndexing() (frame *IndexedNameWithoutIndexing) {
	frame = &IndexedNameWithoutIndexing{}
	frame.Flag1 = 0
	frame.Flag2 = 1
	frame.Flag3 = 1
	return
}

// Literal Header with Incremental Indexing - Indexed Name
//
// 0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 0 | 1 | 0 |    Index (5+)     |
// +---+---+---+-------------------+
// |       Value Length (8+)       |
// +-------------------------------+
// | Value String (Length octets)  |
// +-------------------------------+
type IndexedNameWithIncrementalIndexing struct {
	Flag1       uint8
	Flag2       uint8
	Flag3       uint8
	Index       uint64
	ValueLength uint64
	ValueString string
}

func NewIndexedNameWithIncrementalIndexing() (frame *IndexedNameWithIncrementalIndexing) {
	frame = &IndexedNameWithIncrementalIndexing{}
	frame.Flag1 = 0
	frame.Flag2 = 1
	frame.Flag3 = 0
	return
}

// Literal Header with Incremental Indexing - New Name
//
// 0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 0 | 1 | 0 |         0         |
// +---+---+---+-------------------+
// |       Name Length (8+)        |
// +-------------------------------+
// |  Name String (Length octets)  |
// +-------------------------------+
// |       Value Length (8+)       |
// +-------------------------------+
// | Value String (Length octets)  |
// +-------------------------------+
type NewNameWithIncrementalIndexing struct {
	Flag1       uint8
	Flag2       uint8
	Flag3       uint8
	Index       uint8
	NameLength  uint64
	NameString  string
	ValueLength uint64
	ValueString string
}

func NewNewNameWithIncrementalIndexing() (frame *NewNameWithIncrementalIndexing) {
	frame = &NewNameWithIncrementalIndexing{}
	frame.Flag1 = 0
	frame.Flag2 = 1
	frame.Flag3 = 0
	frame.Index = 0
	return
}

// Literal Header with Substitution Indexing - Indexed Name
//
// 0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 0 | 0 |      Index (6+)       |
// +---+---+-----------------------+
// |    Substituted Index (8+)     |
// +-------------------------------+
// |       Value Length (8+)       |
// +-------------------------------+
// | Value String (Length octets)  |
// +-------------------------------+
type IndexedNameWithSubstitutionIndexing struct {
	Flag1            uint8
	Flag2            uint8
	Index            uint64
	SubstitutedIndex uint64
	ValueLength      uint64
	ValueString      string
}

func NewIndexedNameWithSubstitutionIndexing() (frame *IndexedNameWithSubstitutionIndexing) {
	frame = &IndexedNameWithSubstitutionIndexing{}
	frame.Flag1 = 0
	frame.Flag2 = 0
	return
}

// Literal Header with Substitution Indexing - New Name
//
// 0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 0 | 0 |           0           |
// +---+---+-----------------------+
// |       Name Length (8+)        |
// +-------------------------------+
// |  Name String (Length octets)  |
// +-------------------------------+
// |    Substituted Index (8+)     |
// +-------------------------------+
// |       Value Length (8+)       |
// +-------------------------------+
// | Value String (Length octets)  |
// +-------------------------------+
type NewNameWithSubstitutionIndexing struct {
	Flag1            uint8
	Flag2            uint8
	Index            uint8
	NameLength       uint64
	NameString       string
	SubstitutedIndex uint64
	ValueLength      uint64
	ValueString      string
}

func NewNewNameWithSubstitutionIndexing() (frame *NewNameWithSubstitutionIndexing) {
	frame = &NewNameWithSubstitutionIndexing{}
	frame.Flag1 = 0
	frame.Flag2 = 0
	frame.Index = 0
	return
}
