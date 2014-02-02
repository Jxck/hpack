package hpack

import (
	"net/http"
	"reflect"
	"sort"
	"testing"
)

func TestNewHeaderSet(t *testing.T) {
	header := make(http.Header)
	header.Add("method", "get")
	header.Add("host", "example.com")
	header.Add(":authority", "example.com")
	header.Add("cookie", "a")
	header.Add("cookie", "b")

	expected := HeaderSet{
		{"method", "get"},
		{"host", "example.com"},
		{":authority", "example.com"},
		{"cookie", "a"},
		{"cookie", "b"},
	}
	actual := ToHeaderSet(header)

	for i, hf := range expected {
		if actual[i] != hf {
			t.Errorf("\ngot  %v\nwant %v", actual.Dump(), expected.Dump())
		}
	}
}

func TestToHeader(t *testing.T) {
	header := make(http.Header)
	header.Add("method", "get")
	header.Add("host", "example.com")
	header.Add(":authority", "example.com")
	header.Add("cookie", "a")
	header.Add("cookie", "b")

	headerSet := HeaderSet{
		{"method", "get"},
		{"host", "example.com"},
		{":authority", "example.com"},
		{"cookie", "a"},
		{"cookie", "b"},
	}
	actual := headerSet.ToHeader()

	if !reflect.DeepEqual(header, actual) {
		t.Errorf("\ngot  %v\nwant %v", actual, header)
	}
}

func TestEmit(t *testing.T) {
	hs := NewHeaderSet()
	hf1 := NewHeaderField("key1", "value1")
	hf2 := NewHeaderField("key2", "value2")
	hs.Emit(hf1)
	hs.Emit(hf2)
	if hs.Len() != 2 {
		t.Fatal("%v should length %v", hs, 2)
	}
}

func TestEmitSort(t *testing.T) {
	hs := &HeaderSet{
		HeaderField{":method", "GET"},
		HeaderField{":scheme", "http"},
		HeaderField{":path", "/"},
		HeaderField{":authority", "www.example.com"},
		HeaderField{"cache-control", "no-cache"},
	}

	sort.Sort(hs)
	if (*hs)[0].Name != ":authority" {
		t.Fatal("*hs[0] should length %v but %v", ":authority", (*hs)[0])
	}
}
