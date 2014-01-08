package hpack

import (
	"testing"
)

func TestIndexedHeaderEncode(t *testing.T) {
	var index uint64
	var frame *IndexedHeader

	index = 2
	frame = NewIndexedHeader(index)
	actual := frame.Encode()
	expected := []byte{0x82}
	if !actual.Compare(expected) {
		t.Errorf("\ngot %v\nwant %v", actual, expected)
	}

	index = 180
	frame = NewIndexedHeader(index)
	actual = frame.Encode()
	expected = []byte{0xFF, 0x35}
	if !actual.Compare(expected) {
		t.Errorf("\ngot %v\nwant %v", actual, expected)
	}
}

func TestIndexedLiteralEncode(t *testing.T) {
	var indexing bool
	var index uint64
	var value string

	// No Indexing
	indexing = false
	index = 4
	value = "/sample/path"
	expected := []byte{
		0x44, 0x0c, 0x2f, 0x73,
		0x61, 0x6d, 0x70, 0x6c,
		0x65, 0x2f, 0x70, 0x61,
		0x74, 0x68,
	}

	frame := NewIndexedLiteral(indexing, index, value)
	actual := frame.Encode()
	if !actual.Compare(expected) {
		t.Errorf("\ngot %v\nwant %v", actual, expected)
	}

	// Indexing
	indexing = true
	index = 180
	value = "/sample/path"
	expected = []byte{
		0x3f, 0x75, 0x0c, 0x2f,
		0x73, 0x61, 0x6d, 0x70,
		0x6c, 0x65, 0x2f, 0x70,
		0x61, 0x74, 0x68,
	}

	frame = NewIndexedLiteral(indexing, index, value)
	actual = frame.Encode()
	if !actual.Compare(expected) {
		t.Errorf("\ngot %v\nwant %v", actual, expected)
	}
}

func TestIndexedLiteralEncodeHuffman(t *testing.T) {
	var indexing bool
	var index uint64
	var value string

	// Indexing
	indexing = true
	index = 4
	value = "www.example.com"
	expected := []byte{
		0x04, 0x8b, 0xdb, 0x6d,
		0x88, 0x3e, 0x68, 0xd1,
		0xcb, 0x12, 0x25, 0xba,
		0x7f,
	}

	frame := NewIndexedLiteral(indexing, index, value)
	actual := frame.EncodeHuffman(REQUEST)
	if !actual.Compare(expected) {
		t.Errorf("\ngot %v\nwant %v", actual, expected)
	}
}

func TestStringLiteralEncode(t *testing.T) {
	var indexing bool
	var name, value string = "custom-key", "custom-header"

	// No Indexing
	indexing = false
	expected := []byte{
		0x40, 0x0a, 0x63, 0x75,
		0x73, 0x74, 0x6f, 0x6d,
		0x2d, 0x6b, 0x65, 0x79,
		0x0d, 0x63, 0x75, 0x73,
		0x74, 0x6f, 0x6d, 0x2d,
		0x68, 0x65, 0x61, 0x64,
		0x65, 0x72,
	}

	frame := NewStringLiteral(indexing, name, value)
	actual := frame.Encode()
	if !actual.Compare(expected) {
		t.Errorf("\ngot %v\nwant %v", actual, expected)
	}

	// Indexing
	indexing = true
	expected = []byte{
		0x00, 0x0a, 0x63, 0x75,
		0x73, 0x74, 0x6f, 0x6d,
		0x2d, 0x6b, 0x65, 0x79,
		0x0d, 0x63, 0x75, 0x73,
		0x74, 0x6f, 0x6d, 0x2d,
		0x68, 0x65, 0x61, 0x64,
		0x65, 0x72,
	}

	frame = NewStringLiteral(indexing, name, value)
	actual = frame.Encode()
	if !actual.Compare(expected) {
		t.Errorf("\ngot %v\nwant %v", actual, expected)
	}
}

func TestStringLiteralEncodeHuffman(t *testing.T) {
	var indexing bool
	var name, value string = "custom-key", "custom-value"

	// Indexing
	indexing = true
	expected := []byte{
		0x00, 0x88, 0x4e, 0xb0,
		0x8b, 0x74, 0x97, 0x90,
		0xfa, 0x7f, 0x89, 0x4e,
		0xb0, 0x8b, 0x74, 0x97,
		0x9a, 0x17, 0xa8, 0xff,
	}

	frame := NewStringLiteral(indexing, name, value)
	actual := frame.EncodeHuffman(REQUEST)
	if !actual.Compare(expected) {
		t.Errorf("\ngot %v\nwant %v", actual, expected)
	}
}
