package hpack

import (
	"github.com/jxck/swrap"
)

type Indexing int

const (
	Version int      = 7
	WITH    Indexing = iota
	WITHOUT
	NEVER
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
type IndexedHeader struct {
	Index uint64
}

func NewIndexedHeader(index uint64) (frame *IndexedHeader) {
	frame = new(IndexedHeader)
	frame.Index = index
	return
}

// Literal Header Field with Incremental Indexing - Indexed Name
// | 0 | 1 |      Index (6+)       |
//
// Literal Header Field without Indexing - Indexed Name
// | 0 | 0 | 0 | 0 |  Index (4+)   |
//
// Literal Header Field never Indexed - Indexed Nmae
// | 0 | 0 | 0 | 1 |  Index (4+)   |
//
//  0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 0 |       Flag + Index        |
// +---+---+---+-------------------+
// | H |     Value Length (7+)     |
// +-------------------------------+
// | Value String (Length octets)  |
// +-------------------------------+
type IndexedLiteral struct {
	Indexing    Indexing
	Index       uint64
	ValueLength uint64
	ValueString string
}

func NewIndexedLiteral(indexing Indexing, index uint64, value string) (frame *IndexedLiteral) {
	frame = new(IndexedLiteral)
	frame.Indexing = indexing
	frame.Index = index
	frame.ValueLength = uint64(len(value))
	frame.ValueString = value
	return
}

// Literal Header Field with Incremental Indexing - New Name
// Flag = 64 | 0 | 1 | 0 | 0 | 0 | 0 | 0 | 0 |
//
// Literal Header Field without Indexing - New Name
// Flag =  0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
//
// Literal Header Field never Indexed - New Name
// Flag = 16 | 0 | 0 | 0 | 1 | 0 | 0 | 0 | 0 |
//
//   0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// |     Flag=( 0 | 16 | 64 )      |
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
	Indexing    Indexing
	Index       uint64
	NameLength  uint64
	NameString  string
	ValueLength uint64
	ValueString string
}

func NewStringLiteral(indexing Indexing, name, value string) (frame *StringLiteral) {
	frame = new(StringLiteral)
	frame.Indexing = indexing
	frame.NameLength = uint64(len(name))
	frame.NameString = name
	frame.ValueLength = uint64(len(value))
	frame.ValueString = value
	return
}

// 0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 0 | 0 | 1 | 1 |       0       |
// +---+---------------------------+
// Reference Set Emptying
type EmptyReferenceSet struct{}

func NewEmptyReferenceSet() (frame *EmptyReferenceSet) {
	frame = new(EmptyReferenceSet)
	return frame
}

//
// 0   1   2   3   4   5   6   7
// +---+---+---+---+---+---+---+---+
// | 0 | 0 | 1 | 0 | Max size (4+) |
// +---+---------------------------+
// Maximum Header Table Size Change
type ChangeHeaderTableSize struct {
	MaxSize uint64
}

func NewChangeHeaderTableSize(maxSize uint64) (frame *ChangeHeaderTableSize) {
	frame = new(ChangeHeaderTableSize)
	frame.MaxSize = maxSize
	return
}
