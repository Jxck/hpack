package hpac

import (
	"bytes"
	"testing"
)

func TestIndexedHeaderDecode(t *testing.T) {
	// 0xa6       (indexed header, index = 38: removal from reference set)
	buf := bytes.NewBuffer([]byte{0xA6}) // 10100110

	frame := DecodeHeader(buf)
	f, ok := frame.(*IndexedHeader)
	if !ok {
		t.Fatal("Parsed incorrect frame type:", frame)
	}
	if f.flag1 != 1 {
		t.Errorf("got %v\nwant %v", f.flag1, 1)
	}
	if f.Index != 38 {
		t.Errorf("got %v\nwant %v", f.Index, 38)
	}
}

func TestNewNameWithoutIndexingDecode(t *testing.T) {
	// 0x60      (literal header without indexing, new name)
	// 0x0B      (header name string length = 11)
	// mynewheader
	// 0x05      (header value string length = 5)
	// first
	buf := bytes.NewBuffer([]byte{0x60, 0x0B}) // 0110 00000 0000 1011
	buf.WriteString("mynewheader")
	buf.WriteByte(0x05)
	buf.WriteString("first")

	frame := DecodeHeader(buf)
	f, ok := frame.(*NewNameWithoutIndexing)
	if !ok {
		t.Fatal("Parsed incorrect frame type:", frame)
	}
	if f.flag1 != 0 {
		t.Errorf("got %v\nwant %v", f.flag1, 0)
	}
	if f.flag2 != 1 {
		t.Errorf("got %v\nwant %v", f.flag2, 1)
	}
	if f.flag3 != 1 {
		t.Errorf("got %v\nwant %v", f.flag3, 1)
	}
	if f.Index != 0 {
		t.Errorf("got %v\nwant %v", f.Index, 0)
	}
	if f.NameLength != 11 {
		t.Errorf("got %v\nwant %v", f.NameLength, 11)
	}
	if f.NameString != "mynewheader" {
		t.Errorf("got %v\nwant %v", f.NameString, "mynewheader")
	}
	if f.ValueLength != 5 {
		t.Errorf("got %v\nwant %v", f.ValueLength, 5)
	}
	if f.ValueString != "first" {
		t.Errorf("got %v\nwant %v", f.ValueString, "first")
	}
}

func TestIndexedNameWithoutIndexingEncode(t *testing.T) {
	// 0x64      (literal header without indexing, name index = 3)
	// 0x05      (header value string length = 5)
	// first
	buf := bytes.NewBuffer([]byte{0x64, 0x05}) // 0110 00011 0000 0101
	buf.WriteString("first")

	frame := DecodeHeader(buf)
	f, ok := frame.(*IndexedNameWithoutIndexing)
	if !ok {
		t.Fatal("Parsed incorrect frame type:", frame)
	}
	if f.flag1 != 0 {
		t.Errorf("got %v\nwant %v", f.flag1, 0)
	}
	if f.flag2 != 1 {
		t.Errorf("got %v\nwant %v", f.flag2, 1)
	}
	if f.flag3 != 1 {
		t.Errorf("got %v\nwant %v", f.flag3, 1)
	}
	if f.Index != 3 {
		t.Errorf("got %v\nwant %v", f.Index, 3)
	}
	if f.ValueLength != 5 {
		t.Errorf("got %v\nwant %v", f.ValueLength, 5)
	}
	if f.ValueString != "first" {
		t.Errorf("got %v\nwant %v", f.ValueString, "first")
	}
}

func TestIndexedNameWithIncrementalIndexingDecode(t *testing.T) {
	// 0x44      (literal header with incremental indexing, name index = 3)
	// 0x16      (header value string length = 22)
	// /my-example/index.html
	buf := bytes.NewBuffer([]byte{0x44, 0x16}) // 0100 0100 0001 0110
	buf.WriteString("/my-example/index.html")

	frame := DecodeHeader(buf)
	f, ok := frame.(*IndexedNameWithIncrementalIndexing)
	if !ok {
		t.Fatal("Parsed incorrect frame type:", frame)
	}
	if f.flag1 != 0 {
		t.Errorf("got %v\nwant %v", f.flag1, 0)
	}
	if f.flag2 != 1 {
		t.Errorf("got %v\nwant %v", f.flag2, 1)
	}
	if f.flag3 != 0 {
		t.Errorf("got %v\nwant %v", f.flag3, 0)
	}
	if f.Index != 3 {
		t.Errorf("got %v\nwant %v", f.Index, 3)
	}
	if f.ValueLength != 22 {
		t.Errorf("got %v\nwant %v", f.ValueLength, 22)
	}
	if f.ValueString != "/my-example/index.html" {
		t.Errorf("got %v\nwant %v", f.ValueString, "/my-example/index.html")
	}
}

func TestNewNameWithIncrementalIndexingDecode(t *testing.T) {
	// 0x40      (literal header with incremental indexing, new name)
	// 0x0B      (header name string length = 11)
	// mynewheader
	// 0x05      (header value string length = 5)
	// first
	buf := bytes.NewBuffer([]byte{0x40, 0x0B}) // 01000000 00001011
	buf.WriteString("mynewheader")
	buf.WriteByte(0x05)
	buf.WriteString("first")

	frame := DecodeHeader(buf)
	f, ok := frame.(*NewNameWithIncrementalIndexing)
	if !ok {
		t.Fatal("Parsed incorrect frame type:", frame)
	}
	if f.flag1 != 0 {
		t.Errorf("got %v\nwant %v", f.flag1, 0)
	}
	if f.flag2 != 1 {
		t.Errorf("got %v\nwant %v", f.flag2, 1)
	}
	if f.flag3 != 0 {
		t.Errorf("got %v\nwant %v", f.flag3, 0)
	}
	if f.Index != 0 {
		t.Errorf("got %v\nwant %v", f.Index, 0)
	}
	if f.NameLength != 11 {
		t.Errorf("got %v\nwant %v", f.NameLength, 11)
	}
	if f.NameString != "mynewheader" {
		t.Errorf("got %v\nwant %v", f.NameString, "mynewheader")
	}
	if f.ValueLength != 5 {
		t.Errorf("got %v\nwant %v", f.ValueLength, 5)
	}
	if f.ValueString != "first" {
		t.Errorf("got %v\nwant %v", f.ValueString, "first")
	}
}

func TestIndexedNameWithSubstitutionIndexingDecode(t *testing.T) {
	// 0x04       (literal header, substitution indexing, name index = 3)
	// 0x26       (replaced entry index = 38)
	// 0x1f       (header value string length = 31)
	// /my-example/resources/script.js
	buf := bytes.NewBuffer([]byte{0x04, 0x26, 0x1f}) // 00000100 00100110 00011111
	buf.WriteString("/my-example/resources/script.js")

	frame := DecodeHeader(buf)
	f, ok := frame.(*IndexedNameWithSubstitutionIndexing)
	if !ok {
		t.Fatal("Parsed incorrect frame type:", frame)
	}
	if f.flag1 != 0 {
		t.Errorf("got %v\nwant %v", f.flag1, 0)
	}
	if f.flag2 != 0 {
		t.Errorf("got %v\nwant %v", f.flag2, 0)
	}
	if f.Index != 3 {
		t.Errorf("got %v\nwant %v", f.Index, 3)
	}
	if f.SubstitutedIndex != 38 {
		t.Errorf("got %v\nwant %v", f.SubstitutedIndex, 38)
	}
	if f.ValueLength != 31 {
		t.Errorf("got %v\nwant %v", f.ValueLength, 31)
	}
	if f.ValueString != "/my-example/resources/script.js" {
		t.Errorf("got %v\nwant %v", f.ValueString, "/my-example/resources/script.js")
	}
}

/*
func TestIndexedNameWithIncrementalIndexing3(t *testing.T) {

	// 0x5f 0101 1111 (literal header, incremental indexing, name index = 40) 40n5=[31 9]
	// 0x0a 0000 1010
	// 0x06 0000 0110 (header value string length = 6)
	// second
	buf := bytes.NewBuffer([]byte{0x5f, 0x0a, 0x06})
	buf.WriteString("second")

	frame := DecodeHeader(buf)

	f, ok := frame.(*IndexedNameWithIncrementalIndexing)
	if !ok {
		t.Fatal("Parsed incorrect frame type:", frame)
	}
	if f.flag1 != 0 {
		t.Errorf("got %v\nwant %v", f.flag1, 0)
	}
	if f.flag2 != 1 {
		t.Errorf("got %v\nwant %v", f.flag2, 1)
	}
	if f.flag3 != 0 {
		t.Errorf("got %v\nwant %v", f.flag3, 0)
	}
	if f.Index != 40 {
		t.Errorf("got %v\nwant %v", f.Index, 40)
	}
	if f.ValueLength != 6 {
		t.Errorf("got %v\nwant %v", f.ValueLength, 6)
	}
	if f.ValueString != "second" {
		t.Errorf("got %v\nwant %v", f.ValueString, "second")
	}
}
*/

func TestNewNameWithSubstitutionIndexing(t *testing.T) {
	// 0x0       (literal header with substitution indexing, new name)
	// 0x0B      (header name string length = 11)
	// mynewheader
	// 0x26      (replaced entry index = 38)
	// 0x05      (header value string length = 5)
	// first
	buf := bytes.NewBuffer([]byte{0x0, 0x0B}) // 00000000 00001011
	buf.WriteString("mynewheader")
	buf.WriteByte(0x26)
	buf.WriteByte(0x05)
	buf.WriteString("first")

	frame := DecodeHeader(buf)

	f, ok := frame.(*NewNameWithSubstitutionIndexing)
	if !ok {
		t.Fatal("Parsed incorrect frame type:", frame)
	}
	if f.flag1 != 0 {
		t.Errorf("got %v\nwant %v", f.flag1, 0)
	}
	if f.flag2 != 0 {
		t.Errorf("got %v\nwant %v", f.flag2, 0)
	}
	if f.Index != 0 {
		t.Errorf("got %v\nwant %v", f.Index, 0)
	}
	if f.NameLength != 11 {
		t.Errorf("got %v\nwant %v", f.NameLength, 11)
	}
	if f.NameString != "mynewheader" {
		t.Errorf("got %v\nwant %v", f.NameString, "mynewheader")
	}
	if f.SubstitutedIndex != 38 {
		t.Errorf("got %v\nwant %v", f.SubstitutedIndex, 38)
	}
	if f.ValueLength != 5 {
		t.Errorf("got %v\nwant %v", f.ValueLength, 5)
	}
	if f.ValueString != "first" {
		t.Errorf("got %v\nwant %v", f.ValueString, "first")
	}
}
