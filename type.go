package hpack

import (
	"bytes"
)

type Frame interface {
	Encode() *bytes.Buffer
}

// Indexed Header Field
//
// 	0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 1 |        Index (7+)         |
// +---+---------------------------+
type IndexedHeader struct {
	Index uint64
}

func NewIndexedHeader() (frame *IndexedHeader) {
	frame = &IndexedHeader{}
	return
}

func CreateIndexedHeader(index uint64) (frame *IndexedHeader) {
	frame = NewIndexedHeader()
	frame.Index = index
	return
}

// Literal Header Field without Indexing - Indexed Name
//
//  0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 0 | 1 |      Index (6+)       |
// +---+---+---+-------------------+
// |       Value Length (8+)       |
// +-------------------------------+
// | Value String (Length octets)  |
// +-------------------------------+
type IndexedNameWithoutIndexing struct {
	Index       uint64
	ValueLength uint64
	ValueString string
}

func NewIndexedNameWithoutIndexing() (frame *IndexedNameWithoutIndexing) {
	frame = &IndexedNameWithoutIndexing{}
	return
}

func CreateIndexedNameWithoutIndexing(index uint64, value string) (frame *IndexedNameWithoutIndexing) {
	frame = NewIndexedNameWithoutIndexing()
	frame.Index = index
	frame.ValueLength = uint64(len(value))
	frame.ValueString = value
	return
}

// Literal Header Field without Indexing - New Name
//   0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 0 | 1 |           0           |
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
	Index       uint64
	NameLength  uint64
	NameString  string
	ValueLength uint64
	ValueString string
}

func NewNewNameWithoutIndexing() (frame *NewNameWithoutIndexing) {
	frame = &NewNameWithoutIndexing{}
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

// Literal Header Field with Incremental Indexing - Indexed Name
//
//  0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 0 | 0 |      Index (6+)       |
// +---+---+---+-------------------+
// |       Value Length (8+)       |
// +-------------------------------+
// | Value String (Length octets)  |
// +-------------------------------+
type IndexedNameWithIncrementalIndexing struct {
	Index       uint64
	ValueLength uint64
	ValueString string
}

func NewIndexedNameWithIncrementalIndexing() (frame *IndexedNameWithIncrementalIndexing) {
	frame = &IndexedNameWithIncrementalIndexing{}
	return
}

func CreateIndexedNameWithIncrementalIndexing(index uint64, value string) (frame *IndexedNameWithIncrementalIndexing) {
	frame = NewIndexedNameWithIncrementalIndexing()
	frame.Index = index
	frame.ValueLength = uint64(len(value))
	frame.ValueString = value
	return
}

// Literal Header Field with Incremental Indexing - New Name
//
//   0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 0 | 0 |           0           |
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
	Index       uint8
	NameLength  uint64
	NameString  string
	ValueLength uint64
	ValueString string
}

func NewNewNameWithIncrementalIndexing() (frame *NewNameWithIncrementalIndexing) {
	frame = &NewNameWithIncrementalIndexing{}
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
