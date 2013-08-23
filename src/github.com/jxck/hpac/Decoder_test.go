package hpac

import (
	"bytes"
	"testing"
)

func TestFirstHeaderSet1(t *testing.T) {

	// 0x44      (literal header with incremental indexing, name index = 3)
	// 0x16      (header value string length = 22)
	// /my-example/index.html
	buf := bytes.NewBuffer([]byte{0x44, 0x16}) // 0100 0100 0001 0110
	buf.WriteString("/my-example/index.html")

	frame := DecodeHeader(buf)

	f, ok := frame.(*IncrementalIndexingName)
	if !ok {
		t.Fatal("Parsed incorrect frame type:", frame)
	}
	if f.Flag1 != 0 {
		t.Errorf("got %v\nwant %v", f.Flag1, 0)
	}
	if f.Flag2 != 1 {
		t.Errorf("got %v\nwant %v", f.Flag2, 1)
	}
	if f.Flag3 != 0 {
		t.Errorf("got %v\nwant %v", f.Flag3, 0)
	}
	if f.Index != 3 {
		t.Errorf("got %v\nwant %v", f.Index, 3)
	}
	if f.ValueLength != 22 {
		t.Errorf("got %v\nwant %v", f.Index, 22)
	}
	if f.ValueString != "/my-example/index.html" {
		t.Errorf("got %v\nwant %v", f.ValueString, "/my-example/index.html")
	}
}

func TestFirstHeaderSet2(t *testing.T) {

	// 0x4D      (literal header with incremental indexing, name index = 12)
	// 0x0D      (header value string length = 13)
	// my-user-agent
	buf := bytes.NewBuffer([]byte{0x4D, 0x0D}) // 01001101 00001101
	buf.WriteString("my-user-agent")

	frame := DecodeHeader(buf)

	f, ok := frame.(*IncrementalIndexingName)
	if !ok {
		t.Fatal("Parsed incorrect frame type:", frame)
	}
	if f.Flag1 != 0 {
		t.Errorf("got %v\nwant %v", f.Flag1, 0)
	}
	if f.Flag2 != 1 {
		t.Errorf("got %v\nwant %v", f.Flag2, 1)
	}
	if f.Flag3 != 0 {
		t.Errorf("got %v\nwant %v", f.Flag3, 0)
	}
	if f.Index != 12 {
		t.Errorf("got %v\nwant %v", f.Index, 12)
	}
	if f.ValueLength != 13 {
		t.Errorf("got %v\nwant %v", f.Index, 13)
	}
	if f.ValueString != "my-user-agent" {
		t.Errorf("got %v\nwant %v", f.ValueString, "my-user-agent")
	}
}

func TestFirstHeaderSet3(t *testing.T) {

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

	f, ok := frame.(*IncrementalNewName)
	if !ok {
		t.Fatal("Parsed incorrect frame type:", frame)
	}
	if f.Flag1 != 0 {
		t.Errorf("got %v\nwant %v", f.Flag1, 0)
	}
	if f.Flag2 != 1 {
		t.Errorf("got %v\nwant %v", f.Flag2, 1)
	}
	if f.Flag3 != 0 {
		t.Errorf("got %v\nwant %v", f.Flag3, 0)
	}
	if f.Flag4 != 0 {
		t.Errorf("got %v\nwant %v", f.Flag4, 0)
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

func TestSecondHeaderSet1(t *testing.T) {

	// 0xa6       (indexed header, index = 38: removal from reference set)
	buf := bytes.NewBuffer([]byte{0xA6}) // 10100110

	frame := DecodeHeader(buf)

	f, ok := frame.(*IndexedHeader)
	if !ok {
		t.Fatal("Parsed incorrect frame type:", frame)
	}
	if f.Flag1 != 1 {
		t.Errorf("got %v\nwant %v", f.Flag1, 1)
	}
	if f.Index != 38 {
		t.Errorf("got %v\nwant %v", f.Index, 38)
	}
	t.Log(f.Index)
}

func TestSecondHeaderSet2(t *testing.T) {

	// 0xa8       (indexed header, index = 40: removal from reference set)
	buf := bytes.NewBuffer([]byte{0xA8}) // 10101000

	frame := DecodeHeader(buf)

	f, ok := frame.(*IndexedHeader)
	if !ok {
		t.Fatal("Parsed incorrect frame type:", frame)
	}
	if f.Flag1 != 1 {
		t.Errorf("got %v\nwant %v", f.Flag1, 1)
	}
	if f.Index != 40 {
		t.Errorf("got %v\nwant %v", f.Index, 40)
	}
}

func TestSecondHeaderSet3(t *testing.T) {

	// 0x04       (literal header, substitution indexing, name index = 3)
	// 0x26       (replaced entry index = 38)
	// 0x1f       (header value string length = 31)
	// /my-example/resources/script.js
	buf := bytes.NewBuffer([]byte{0x04, 0x26, 0x1f}) // 00000100 00100110 00011111
	buf.WriteString("/my-example/resources/script.js")

	frame := DecodeHeader(buf)

	f, ok := frame.(*SubstitutionIndexedName)
	if !ok {
		t.Fatal("Parsed incorrect frame type:", frame)
	}
	if f.Flag1 != 0 {
		t.Errorf("got %v\nwant %v", f.Flag1, 0)
	}
	if f.Flag2 != 0 {
		t.Errorf("got %v\nwant %v", f.Flag2, 0)
	}
	if f.Index != 3 {
		t.Errorf("got %v\nwant %v", f.Index, 3)
	}
	if f.SubstitutedIndex != 38 {
		t.Errorf("got %v\nwant %v", f.Index, 38)
	}
	if f.ValueLength != 31 {
		t.Errorf("got %v\nwant %v", f.Index, 31)
	}
	if f.ValueString != "/my-example/resources/script.js" {
		t.Errorf("got %v\nwant %v", f.ValueString, "/my-example/resources/script.js")
	}
}
