package hpac

import (
	"testing"
)

func TestHeaderTableAdd(t *testing.T) {
	reqHT := NewRequestHeaderTable()
	reqHT.Add("Hello", "World")

	h := reqHT[len(reqHT)-1]
	if h.Name != "Hello" || h.Value != "World" {
		t.Errorf("got %v\nwant %v", h, Header{"Hello", "World"})
	}
}

func TestHeaderTableReplace(t *testing.T) {
	reqHT := NewRequestHeaderTable()
	reqHT.Replace("Hello", "World", 10)

	h := reqHT[10]
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
