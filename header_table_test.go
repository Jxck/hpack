package hpack

import (
	"testing"
)

func TestHeaderTableSizeLen(t *testing.T) {
	ht := HeaderTable{
		DEFAULT_HEADER_TABLE_SIZE,
		[]*HeaderField{
			NewHeaderField("1234", "1234"),
			NewHeaderField("1234", "1234"),
			NewHeaderField("1234", "1234"),
			NewHeaderField("1234", "1234"),
			NewHeaderField("1234", "1234"),
		},
	}

	size := ht.Size()
	length := ht.Len()

	if size != 200 {
		t.Errorf("got %v\nwant %v", size, 200)
	}

	if length != 5 {
		t.Errorf("got %v\nwant %v", length, 5)
	}
}

func TestHeaderTablePush(t *testing.T) {
	ht := HeaderTable{
		200,
		[]*HeaderField{ // 200byte
			{"1234", "1234"},
			{"1234", "1234"},
			{"1234", "1234"},
			{"1234", "1234"},
			{"1234", "1234"}, // 40byte
		},
	}

	// should remove 2 entry before add
	hf := NewHeaderField("hello", "world")
	ht.Push(hf) // 42byte
	size := ht.Size()
	expected := 200 - 40 - 40 + 42
	if size != expected {
		t.Errorf("got %v\nwant %v", size, expected)
	}

	length := ht.Len()
	expected = 4
	if length != expected {
		t.Errorf("got %v\nwant %v", length, expected)
	}

	h := ht.HeaderFields[0]
	if h.Name != "hello" || h.Value != "world" {
		t.Errorf("\ngot %v\nwant %v", h, hf)
	}
}

/*
func TestHeaderTableAddBigEntry(t *testing.T) {
	ht := HeaderTable{
		40,
		Headers{
			{"1234", "1234"}, // 40
		},
	}
	ht.Add("12345", "12345") // over 40
	size := ht.Size()
	if size != 0 {
		t.Errorf("got %v\nwant %v", size, 0)
	}
}

func TestHeaderTableReplace(t *testing.T) {
	ht := HeaderTable{
		200,
		Headers{ // 200byte
			{"0000", "0000"},
			{"1111", "1111"},
			{"2222", "2222"},
			{"3333", "3333"},
			{"4444", "4444"}, // 40byte
		},
	}
	ht.Replace("xxxx", "yyyy", 4)

	h := ht.Headers[4]
	if h.Name != "xxxx" || h.Value != "yyyy" {
		t.Errorf("got %v\nwant %v", h, Header{"xxxx", "yyyy"})
	}
}

func TestHeaderTableReplaceWithAlloc(t *testing.T) {
	ht := HeaderTable{
		200,
		Headers{ // 200byte
			{"0000", "0000"},
			{"1111", "1111"},
			{"2222", "2222"},
			{"3333", "3333"},
			{"4444", "4444"}, // 40byte
		},
	}
	ht.Replace("xxxxx", "yyyyy", 3)

	h := ht.Headers[2]
	if h.Name != "xxxxx" || h.Value != "yyyyy" {
		t.Errorf("got %v\nwant %v", h, Header{"xxxxx", "yyyyy"})
	}

	ht = HeaderTable{
		200,
		Headers{ // 200byte
			{"0000", "0000"},
			{"1111", "1111"},
			{"2222", "2222"},
			{"3333", "3333"},
			{"4444", "4444"}, // 40byte
		},
	}
	ht.Replace("xxxxx", "yyyyy", 0)

	h = ht.Headers[0]
	if h.Name != "xxxxx" || h.Value != "yyyyy" {
		t.Errorf("got %v\nwant %v", h, Header{"xxxxx", "yyyyy"})
	}
}

func TestHeaderTableReplaceBigEntry(t *testing.T) {
	ht := HeaderTable{
		40,
		Headers{
			{"1234", "1234"}, // 40
		},
	}
	ht.Replace("12345", "12345", 0) // over 40
	size := ht.Size()
	if size != 0 {
		t.Errorf("got %v\nwant %v", size, 0)
	}
}

func TestHeaderTableRemove(t *testing.T) {
	reqHT := NewRequestHeaderTable()
	reqHT.Remove(3)

	h := reqHT.Headers[3]
	if h.Name != ":method" || h.Value != "GET" {
		t.Errorf("got %v\nwant %v", h, Header{":method", "GET"})
	}

	reqHT = NewRequestHeaderTable()
	reqHT.Remove(0)

	h = reqHT.Headers[0]
	if h.Name != ":scheme" || h.Value != "https" {
		t.Errorf("got %v\nwant %v", h, Header{":scheme", "https"})
	}
}

func TestHeaderTableAllocSpace(t *testing.T) {
	ht := HeaderTable{
		200,
		Headers{ // 200byte
			{"1234", "1234"},
			{"1234", "1234"},
			{"1234", "1234"},
			{"1234", "1234"},
			{"1234", "1234"},
		},
	}
	removed := ht.AllocSpace(20) // remove 1 entry (40byte)
	if removed != 1 {
		t.Errorf("got %v\nwant %v", removed, 1)
	}

	size := ht.Size()
	if size != 160 {
		t.Errorf("got %v\nwant %v", size, 160)
	}
}

func TestHeaderTableDeleteAll(t *testing.T) {
	ht := HeaderTable{
		200,
		Headers{
			{"1234", "1234"},
			{"1234", "1234"},
			{"1234", "1234"},
			{"1234", "1234"},
			{"1234", "1234"},
		},
	}
	ht.DeleteAll()

	size := ht.Size()
	if size != 0 {
		t.Errorf("got %v\nwant %v", size, 0)
	}
}

func TestHeaderTableSearch(t *testing.T) {
	reqHT := NewRequestHeaderTable()
	i, h := reqHT.SearchHeader("Hello", "World")
	if i != -1 || h != nil {
		t.Errorf("got %v %v\nwant %v %v", i, h, -1, nil)
	}

	i, h = reqHT.SearchHeader(":path", "/index")
	if i != 3 || h != nil {
		t.Errorf("got %v %v\nwant %v %v", i, h, 3, nil)
	}

	i, h = reqHT.SearchHeader(":path", "/")
	expected := Header{":path", "/"}
	if i != 3 || *h != expected {
		t.Errorf("got %v %v\nwant %v %v", i, h, 3, expected)
	}
}
*/
