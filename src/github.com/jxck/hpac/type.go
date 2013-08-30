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
	Index uint8
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
	Index       uint32
	NameLength  uint32
	NameString  string
	ValueLength uint32
	ValueString string
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
	Index       uint32
	ValueLength uint32
	ValueString string
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
	Index       uint32
	ValueLength uint32
	ValueString string
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
	NameLength  uint32
	NameString  string
	ValueLength uint32
	ValueString string
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
	Index            uint32
	SubstitutedIndex uint32
	ValueLength      uint32
	ValueString      string
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
	NameLength       uint32
	NameString       string
	SubstitutedIndex uint32
	ValueLength      uint32
	ValueString      string
}
