package hpack

import (
	assert "github.com/jxck/assertion"
	"github.com/jxck/swrap"
	"testing"
)

func TestIndexedHeaderDecode(t *testing.T) {
	buf := swrap.New([]byte{0x82})

	// expected
	var index uint64 = 2

	decoded := DecodeHeader(&buf)
	frame, ok := decoded.(*IndexedHeader)
	if !ok {
		t.Errorf("Decoded to incorrect frame type: %T", frame)
	}
	assert.Equal(t, frame.Index, index)
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

	decoded := DecodeHeader(&buf)
	frame, ok := decoded.(*IndexedLiteral)
	if !ok {
		t.Errorf("Decoded to incorrect frame type: %T", frame)
	}
	assert.Equal(t, frame, NewIndexedLiteral(indexing, index, value))
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

	decoded := DecodeHeader(&buf)
	frame, ok := decoded.(*StringLiteral)
	if !ok {
		t.Errorf("Decoded to incorrect frame type: %T", frame)
	}
	assert.Equal(t, frame, NewStringLiteral(indexing, name, value))
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

	decoded := DecodeHeader(&buf)
	frame, ok := decoded.(*IndexedLiteral)
	if !ok {
		t.Errorf("Decoded to incorrect frame type: %T", frame)
	}
	assert.Equal(t, frame, NewIndexedLiteral(indexing, index, value))
}
