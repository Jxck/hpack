package hpack

import (
	"bytes"
	"testing"
)

func TestIndexedHeaderEncode(t *testing.T) {
	var index uint64 = 2
	var frame *IndexedHeader
	frame = NewIndexedHeader(index)

	actual := frame.Encode().Bytes()
	expected := []byte{0x82}
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestIndexedLiteral_NoIndexing(t *testing.T) {
	var indexing bool = false
	var index uint64 = 4
	var value string = "/sample/path"
	frame := NewIndexedLiteral(indexing, index, value)

	actual := frame.Encode().Bytes()
	expected := bytes.NewBuffer([]byte{
		0x44, 0x0c, 0x2f, 0x73,
		0x61, 0x6d, 0x70, 0x6c,
		0x65, 0x2f, 0x70, 0x61,
		0x74, 0x68,
	}).Bytes()
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestStringLiteral_Indexing(t *testing.T) {
	var indexing bool = true
	var name, value string = "custom-key", "custom-header"
	frame := NewStringLiteral(indexing, name, value)

	actual := frame.Encode().Bytes()
	expected := bytes.NewBuffer([]byte{
		0x00, 0x0a, 0x63, 0x75,
		0x73, 0x74, 0x6f, 0x6d,
		0x2d, 0x6b, 0x65, 0x79,
		0x0d, 0x63, 0x75, 0x73,
		0x74, 0x6f, 0x6d, 0x2d,
		0x68, 0x65, 0x61, 0x64,
		0x65, 0x72,
	}).Bytes()
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
