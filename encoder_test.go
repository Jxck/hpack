package hpack

import (
	"bytes"
	"testing"
)

func TestIndexedHeaderEncode(t *testing.T) {
	frame := NewIndexedHeader()
	frame.Index = 38

	actual := frame.Encode().Bytes()
	expected := []byte{0xA6}
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestNewNameWithoutIndexingEncode(t *testing.T) {
	frame := NewNewNameWithoutIndexing()
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
	frame := NewIndexedNameWithoutIndexing()
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
	frame := NewIndexedNameWithIncrementalIndexing()
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
	frame := NewNewNameWithIncrementalIndexing()
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

func TestIndexedNameWithSubstitutionIndexingEncode(t *testing.T) {
	frame := NewIndexedNameWithSubstitutionIndexing()
	frame.Index = 3
	frame.SubstitutedIndex = 38
	frame.ValueLength = 31
	frame.ValueString = "/my-example/resources/script.js"

	actual := frame.Encode().Bytes()
	buf := bytes.NewBuffer([]byte{0x04, 0x26, 0x1f}) // 00000100 00100110 00011111
	buf.WriteString("/my-example/resources/script.js")
	expected := buf.Bytes()
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestNewNameWithSubstitutionIndexingEncode(t *testing.T) {
	frame := NewNewNameWithSubstitutionIndexing()
	frame.NameLength = 11
	frame.NameString = "mynewheader"
	frame.SubstitutedIndex = 38
	frame.ValueLength = 5
	frame.ValueString = "first"

	actual := frame.Encode().Bytes()
	buf := bytes.NewBuffer([]byte{0x0, 0x0B}) // 00000000 00001011
	buf.WriteString("mynewheader")
	buf.WriteByte(0x26)
	buf.WriteByte(0x05)
	buf.WriteString("first")
	expected := buf.Bytes()
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
