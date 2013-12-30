package hpack

import (
	"github.com/jxck/swrap"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestIndexedHeaderDecode(t *testing.T) {
	buf := swrap.New([]byte{0x82})

	// expected
	var index uint64 = 2

	decoded := DecodeHeader(&buf, REQUEST)
	frame, ok := decoded.(*IndexedHeader)
	if !ok {
		t.Errorf("Decoded to incorrect frame type: %T", frame)
	}
	if frame.Index != index {
		t.Errorf("got %v\nwant %v", frame.Index, index)
	}
}

func TestIndexedLiteralDecode_NoIndexing_NoHuffman(t *testing.T) {
	buf := swrap.New([]byte{
		0x44, 0x0c, 0x2f, 0x73,
		0x61, 0x6d, 0x70, 0x6c,
		0x65, 0x2f, 0x70, 0x61,
		0x74, 0x68,
	})

	// expected
	var indexing bool = false
	var index uint64 = 4
	var value string = "/sample/path"

	decoded := DecodeHeader(&buf, REQUEST)
	frame, ok := decoded.(*IndexedLiteral)
	if !ok {
		t.Errorf("Decoded to incorrect frame type: %T", frame)
	}
	if frame.Indexing != indexing ||
		frame.Index != index ||
		frame.ValueLength != uint64(len(value)) ||
		frame.ValueString != value {
		t.Errorf(`
frame      = %v
---should---
indexing   = %v
index      = %v
len(value) = %v
value      = %v
`, frame, indexing, index, len(value), value)
	}
}

func TestStringLiteralDecode_Indexing_NoHuffman(t *testing.T) {
	buf := swrap.New([]byte{
		0x00, 0x0a, 0x63, 0x75,
		0x73, 0x74, 0x6f, 0x6d,
		0x2d, 0x6b, 0x65, 0x79,
		0x0d, 0x63, 0x75, 0x73,
		0x74, 0x6f, 0x6d, 0x2d,
		0x68, 0x65, 0x61, 0x64,
		0x65, 0x72,
	})

	// expected
	var indexing bool = true
	var name, value string = "custom-key", "custom-header"

	decoded := DecodeHeader(&buf, REQUEST)
	frame, ok := decoded.(*StringLiteral)
	if !ok {
		t.Errorf("Decoded to incorrect frame type: %T", frame)
	}
	if frame.Indexing != indexing ||
		frame.Index != 0 ||
		frame.NameLength != uint64(len(name)) ||
		frame.NameString != name ||
		frame.ValueLength != uint64(len(value)) ||
		frame.ValueString != value {
		t.Errorf(`
frame      = %v
---should---
indexing   = %v
index      = %v
len(name)  = %v
name       = %v
len(value) = %v
value      = %v
`, frame, indexing, 0, len(name), name, len(value), value)
	}
}

func TestIndexedLiteralDecode_Indexing_Huffman(t *testing.T) {
	buf := swrap.New([]byte{
		0x04, 0x8b, 0xdb, 0x6d,
		0x88, 0x3e, 0x68, 0xd1,
		0xcb, 0x12, 0x25, 0xba,
		0x7f,
	})

	// expected
	var indexing bool = true
	var index uint64 = 4
	var value string = "www.example.com"

	decoded := DecodeHeader(&buf, REQUEST)
	frame, ok := decoded.(*IndexedLiteral)
	if !ok {
		t.Errorf("Decoded to incorrect frame type: %T", frame)
	}
	if frame.Indexing != indexing ||
		frame.Index != index ||
		frame.ValueLength != uint64(len(value)) ||
		frame.ValueString != value {
		t.Errorf(`
frame      = %v
---should---
indexing   = %v
index      = %v
len(value) = %v
value      = %v
`, frame, indexing, index, len(value), value)
	}
}
