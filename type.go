package hpack

import (
	"github.com/jxck/swrap"
)

const (
	Version int = 6
)

type Frame interface {
	Encode() *swrap.SWrap
}

// Indexed Header Field
//
// 	0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 1 |        Index (7+)         |
// +---+---------------------------+
//
// if Index == 0
//
// 0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 1 |             0             |
// +---+---------------------------+
// Reference Set Emptying
//
// 0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 0 |   New maximum size (7+)   |
// +---+---------------------------+
// Maximum Header Table Size Change
//
type IndexedHeader struct {
	Index uint64
	// Refset Emptify if Option == 128
	// Table Size Change if Option < 128
	Option uint8
}

func NewIndexedHeader(index uint64) (frame *IndexedHeader) {
	frame = new(IndexedHeader)
	frame.Index = index
	return
}

// Literal Header Field without Indexing - Indexed Name (F=1)
// Literal Header Field with Incremental Indexing - Indexed Name (F=0)
//
//  0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 0 | F |      Index (6+)       |
// +---+---+---+-------------------+
// | H |     Value Length (7+)     |
// +-------------------------------+
// | Value String (Length octets)  |
// +-------------------------------+
type IndexedLiteral struct {
	Indexing    bool
	Index       uint64
	ValueLength uint64
	ValueString string
}

func NewIndexedLiteral(indexing bool, index uint64, value string) (frame *IndexedLiteral) {
	frame = new(IndexedLiteral)
	frame.Indexing = indexing
	frame.Index = index
	frame.ValueLength = uint64(len(value))
	frame.ValueString = value
	return
}

// Literal Header Field without Indexing - New Name (F=1)
// Literal Header Field with Incremental Indexing - New Name (F=0)
//
//   0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 0 | F |           0           |
// +---+---+---+-------------------+
// | H |     Name Length (7+)      |
// +-------------------------------+
// |  Name String (Length octets)  |
// +-------------------------------+
// | H |     Value Length (7+)     |
// +-------------------------------+
// | Value String (Length octets)  |
// +-------------------------------+
type StringLiteral struct {
	Indexing    bool
	Index       uint64
	NameLength  uint64
	NameString  string
	ValueLength uint64
	ValueString string
}

func NewStringLiteral(indexing bool, name, value string) (frame *StringLiteral) {
	frame = new(StringLiteral)
	frame.Indexing = indexing
	frame.NameLength = uint64(len(name))
	frame.NameString = name
	frame.ValueLength = uint64(len(value))
	frame.ValueString = value
	return
}
