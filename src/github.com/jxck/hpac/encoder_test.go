package hpac

import (
	"bytes"
	"testing"
)

func TestIndexedHeaderDecode(t *testing.T) {
	frame := NewIndexedHeader()
	frame.Index = 10

	actual := EncodeHeader(frame).Bytes()
	expected := []byte{138}
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestNewNameWithoutIndexingDecode(t *testing.T) {
	frame := NewNewNameWithoutIndexing()
	frame.NameLength = 11
	frame.NameString = "mynewheader"
	frame.ValueLength = 5
	frame.ValueString = "first"

	actual := EncodeHeader(frame).Bytes()
	buf := bytes.NewBuffer([]byte{0x60, 0x0B})
	buf.WriteString("mynewheader")
	buf.WriteByte(0x05)
	buf.WriteString("first")
	expected := buf.Bytes()
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestIndexedNameWithoutIndexingDecode(t *testing.T) {
	frame := NewIndexedNameWithoutIndexing()
	frame.Index = 3
	frame.ValueLength = 5
	frame.ValueString = "first"

	actual := EncodeHeader(frame).Bytes()
	buf := bytes.NewBuffer([]byte{0x64, 0x05})
	buf.WriteString("first")
	expected := buf.Bytes()
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
