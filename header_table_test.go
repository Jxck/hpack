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

	// should evict 2 entry before add
	hf := NewHeaderField("hello", "world")
	ht.Push(hf) // 42byte

	// should reduce size 2 * 40byte
	size := ht.Size()
	if size != 200-40-40+42 {
		t.Errorf("got %v\nwant %v", size, 200-40-40+42)
	}

	// should reduce length 2
	length := ht.Len()
	if length != 4 {
		t.Errorf("got %v\nwant %v", length, 4)
	}

	// first field should added hf
	h := ht.HeaderFields[0]
	if h.Name != "hello" || h.Value != "world" {
		t.Errorf("\ngot %v\nwant %v", h, hf)
	}
}

func TestHeaderTableAddBigEntry(t *testing.T) {
	ht := HeaderTable{
		40,
		[]*HeaderField{
			{"1234", "1234"}, // 40
		},
	}

	hf := NewHeaderField("12345", "12345") // over 40
	ht.Push(hf)                            // 42byte

	size := ht.Size()
	length := ht.Len()

	if size != 0 {
		t.Errorf("got %v\nwant %v", size, 0)
	}

	if length != 0 {
		t.Errorf("got %v\nwant %v", length, 0)
	}
}

func TestHeaderTableRemove(t *testing.T) {
	ht := HeaderTable{
		DEFAULT_HEADER_TABLE_SIZE,
		[]*HeaderField{ // 200byte
			{"1111", "1111"},
			{"2222", "2222"},
			{"3333", "3333"},
			{"4444", "4444"},
			{"5555", "5555"}, // 40byte
		},
	}

	ht.Remove(3)

	hf := ht.HeaderFields[3]
	if hf.Name != "5555" || hf.Value != "5554" {
		t.Errorf("ht[%v] of %v shoud %v", 3, ht.Dump(), HeaderField{"5555", "5555"})
	}
}

/*
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
