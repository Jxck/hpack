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

func TestIndexedNameWithoutIndexingEncode(t *testing.T) {
	var index uint64 = 3
	var value string = "first"
	frame := NewIndexedNameWithoutIndexing(index, value)

	actual := frame.Encode().Bytes()
	buf := bytes.NewBuffer([]byte{0x64, 0x05})
	buf.WriteString(value)
	expected := buf.Bytes()
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestNewNameWithoutIndexingEncode(t *testing.T) {
	var name, value string = "mynewheader", "first"
	frame := NewNewNameWithoutIndexing(name, value)

	actual := frame.Encode().Bytes()
	buf := bytes.NewBuffer([]byte{0x60, 0x0B})
	buf.WriteString(name)
	buf.WriteByte(0x05)
	buf.WriteString(value)
	expected := buf.Bytes()
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestIndexedNameWithIncrementalIndexingEncode(t *testing.T) {
	var index uint64 = 3
	var value string = "/my-example/index.html"
	frame := NewIndexedNameWithIncrementalIndexing(index, value)

	actual := frame.Encode().Bytes()
	buf := bytes.NewBuffer([]byte{0x44, 0x16})
	buf.WriteString(value)
	expected := buf.Bytes()
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestNewNameWithIncrementalIndexingEncode(t *testing.T) {
	var name, value string = "mynewheader", "first"
	frame := NewNewNameWithIncrementalIndexing(name, value)

	actual := frame.Encode().Bytes()
	buf := bytes.NewBuffer([]byte{0x40, 0x0B})
	buf.WriteString(name)
	buf.WriteByte(0x05)
	buf.WriteString(value)
	expected := buf.Bytes()
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
