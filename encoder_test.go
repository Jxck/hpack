package hpack

import (
	"bytes"
	"testing"
)

func TestIndexedHeaderEncode(t *testing.T) {
	frame := new(IndexedHeader)
	frame.Index = 38

	actual := frame.Encode().Bytes()
	expected := []byte{0xA6}
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestNewNameWithoutIndexingEncode(t *testing.T) {
	frame := new(NewNameWithoutIndexing)
	frame.NameLength = 11
	frame.NameString = "mynewheader"
	frame.ValueLength = 5
	frame.ValueString = "first"

	actual := frame.Encode().Bytes()
	buf := bytes.NewBuffer([]byte{0x60, 0x0B})
	buf.WriteString("mynewheader")
	buf.WriteByte(0x05)
	buf.WriteString("first")
	expected := buf.Bytes()
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestIndexedNameWithoutIndexingEncode(t *testing.T) {
	frame := new(IndexedNameWithoutIndexing)
	frame.Index = 3
	frame.ValueLength = 5
	frame.ValueString = "first"

	actual := frame.Encode().Bytes()
	buf := bytes.NewBuffer([]byte{0x64, 0x05})
	buf.WriteString("first")
	expected := buf.Bytes()
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestIndexedNameWithIncrementalIndexingEncode(t *testing.T) {
	frame := new(IndexedNameWithIncrementalIndexing)
	frame.Index = 3
	frame.ValueLength = 22
	frame.ValueString = "/my-example/index.html"

	actual := frame.Encode().Bytes()
	buf := bytes.NewBuffer([]byte{0x44, 0x16})
	buf.WriteString("/my-example/index.html")
	expected := buf.Bytes()
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestNewNameWithIncrementalIndexingEncode(t *testing.T) {
	frame := new(NewNameWithIncrementalIndexing)
	frame.NameLength = 11
	frame.NameString = "mynewheader"
	frame.ValueLength = 5
	frame.ValueString = "first"

	actual := frame.Encode().Bytes()
	buf := bytes.NewBuffer([]byte{0x40, 0x0B})
	buf.WriteString("mynewheader")
	buf.WriteByte(0x05)
	buf.WriteString("first")
	expected := buf.Bytes()
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
