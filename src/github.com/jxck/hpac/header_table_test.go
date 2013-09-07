package hpac

import (
	"testing"
)

func TestHeaderSize(t *testing.T) {
	h := Header{"hello", "world"}
	if h.Size() != 42 {
		t.Errorf("got %v\nwant %v", h.Size(), 42)
	}
}

func TestHeaderTableSize(t *testing.T) {
	ht := HeaderTable{
		Headers{
			{"12345", "12345"},
			{"12345", "12345"},
			{"12345", "12345"},
			{"12345", "12345"},
			{"12345", "12345"},
		},
	}
	size := ht.Size()

	if size != 210 {
		t.Errorf("got %v\nwant %v", size, 210)
	}
}

func TestHeaderTableAdd(t *testing.T) {
	reqHT := NewRequestHeaderTable()
	reqHT.Add("Hello", "World")

	h := reqHT.Headers[len(reqHT.Headers)-1]
	if h.Name != "Hello" || h.Value != "World" {
		t.Errorf("got %v\nwant %v", h, Header{"Hello", "World"})
	}
}

func TestHeaderTableReplace(t *testing.T) {
	reqHT := NewRequestHeaderTable()
	reqHT.Replace("Hello", "World", 10)

	h := reqHT.Headers[10]
	if h.Name != "Hello" || h.Value != "World" {
		t.Errorf("got %v\nwant %v", h, Header{"Hello", "World"})
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
