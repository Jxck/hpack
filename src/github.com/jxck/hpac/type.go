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
	flag1 uint8
	Index uint64
}

func NewIndexedHeader() (frame *IndexedHeader) {
	frame = &IndexedHeader{}
	frame.flag1 = 1
	return
}

func CreateIndexedHeader(index uint64) (frame *IndexedHeader) {
	frame = NewIndexedHeader()
	frame.Index = index
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
	flag1       uint8
	flag2       uint8
	flag3       uint8
	Index       uint64
	NameLength  uint64
	NameString  string
	ValueLength uint64
	ValueString string
}

func NewNewNameWithoutIndexing() (frame *NewNameWithoutIndexing) {
	frame = &NewNameWithoutIndexing{}
	frame.flag1 = 0
	frame.flag2 = 1
	frame.flag3 = 1
	frame.Index = 0
	return
}

func CreateNewNameWithoutIndexing(name, value string) (frame *NewNameWithoutIndexing) {
	frame = NewNewNameWithoutIndexing()
	frame.NameLength = uint64(len(name))
	frame.NameString = name
	frame.ValueLength = uint64(len(value))
	frame.ValueString = value
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
	flag1       uint8
	flag2       uint8
	flag3       uint8
	Index       uint64
	ValueLength uint64
	ValueString string
}

func NewIndexedNameWithoutIndexing() (frame *IndexedNameWithoutIndexing) {
	frame = &IndexedNameWithoutIndexing{}
	frame.flag1 = 0
	frame.flag2 = 1
	frame.flag3 = 1
	return
}

func CreateIndexedNameWithoutIndexing(index uint64, value string) (frame *IndexedNameWithoutIndexing) {
	frame = NewIndexedNameWithoutIndexing()
	frame.Index = index
	frame.ValueLength = uint64(len(value))
	frame.ValueString = value
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
	flag1       uint8
	flag2       uint8
	flag3       uint8
	Index       uint64
	ValueLength uint64
	ValueString string
}

func NewIndexedNameWithIncrementalIndexing() (frame *IndexedNameWithIncrementalIndexing) {
	frame = &IndexedNameWithIncrementalIndexing{}
	frame.flag1 = 0
	frame.flag2 = 1
	frame.flag3 = 0
	return
}

func CreateIndexedNameWithIncrementalIndexing(index uint64, value string) (frame *IndexedNameWithIncrementalIndexing) {
	frame = NewIndexedNameWithIncrementalIndexing()
	frame.Index = index
	frame.ValueLength = uint64(len(value))
	frame.ValueString = value
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
	flag1       uint8
	flag2       uint8
	flag3       uint8
	Index       uint8
	NameLength  uint64
	NameString  string
	ValueLength uint64
	ValueString string
}

func NewNewNameWithIncrementalIndexing() (frame *NewNameWithIncrementalIndexing) {
	frame = &NewNameWithIncrementalIndexing{}
	frame.flag1 = 0
	frame.flag2 = 1
	frame.flag3 = 0
	frame.Index = 0
	return
}

func CreateNewNameWithIncrementalIndexing(name, value string) (frame *NewNameWithIncrementalIndexing) {
	frame = NewNewNameWithIncrementalIndexing()
	frame.NameLength = uint64(len(name))
	frame.NameString = name
	frame.ValueLength = uint64(len(value))
	frame.ValueString = value
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
	flag1            uint8
	flag2            uint8
	Index            uint64
	SubstitutedIndex uint64
	ValueLength      uint64
	ValueString      string
}

func NewIndexedNameWithSubstitutionIndexing() (frame *IndexedNameWithSubstitutionIndexing) {
	frame = &IndexedNameWithSubstitutionIndexing{}
	frame.flag1 = 0
	frame.flag2 = 0
	return
}

func CreateIndexedNameWithSubstitutionIndexing(index, substitutedIndex uint64, value string) (frame *IndexedNameWithSubstitutionIndexing) {
	frame = NewIndexedNameWithSubstitutionIndexing()
	frame.Index = index
	frame.SubstitutedIndex = substitutedIndex
	frame.ValueLength = uint64(len(value))
	frame.ValueString = value
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
	flag1            uint8
	flag2            uint8
	Index            uint8
	NameLength       uint64
	NameString       string
	SubstitutedIndex uint64
	ValueLength      uint64
	ValueString      string
}

func NewNewNameWithSubstitutionIndexing() (frame *NewNameWithSubstitutionIndexing) {
	frame = &NewNameWithSubstitutionIndexing{}
	frame.flag1 = 0
	frame.flag2 = 0
	frame.Index = 0
	return
}

func CreateNewNameWithSubstitutionIndexing(name string, substitutedIndex uint64, value string) (frame *NewNameWithSubstitutionIndexing) {
	frame = NewNewNameWithSubstitutionIndexing()
	frame.NameLength = uint64(len(name))
	frame.NameString = name
	frame.SubstitutedIndex = substitutedIndex
	frame.ValueLength = uint64(len(value))
	frame.ValueString = value
	return
}
