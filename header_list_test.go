package hpack

import (
	assert "github.com/Jxck/assertion"
	"net/http"
	"sort"
	"testing"
)

func TestNewHeaderList(t *testing.T) {
	header := make(http.Header)
	header.Add("method", "get")
	header.Add("host", "example.com")
	header.Add(":authority", "example.com")
	header.Add("cookie", "a")
	header.Add("cookie", "b")

	expected := HeaderList{
		{"method", "get"},
		{"host", "example.com"},
		{":authority", "example.com"},
		{"cookie", "a"},
		{"cookie", "b"},
	}
	actual := ToHeaderList(header)
	assert.Equal(t, actual, expected)
}

func TestToHeader(t *testing.T) {
	header := make(http.Header)
	header.Add("method", "get")
	header.Add("host", "example.com")
	header.Add(":authority", "example.com")
	header.Add("cookie", "a")
	header.Add("cookie", "b")

	headerList := HeaderList{
		{"method", "get"},
		{"host", "example.com"},
		{":authority", "example.com"},
		{"cookie", "a"},
		{"cookie", "b"},
	}
	actual := headerList.ToHeader()
	assert.Equal(t, actual, header)
}

func TestEmit(t *testing.T) {
	hl := NewHeaderList()
	hf1 := NewHeaderField("key1", "value1")
	hf2 := NewHeaderField("key2", "value2")
	hl.Emit(hf1)
	hl.Emit(hf2)
	assert.Equal(t, hl.Len(), 2)
}

func TestEmitSort(t *testing.T) {
	hl := &HeaderList{
		HeaderField{":method", "GET"},
		HeaderField{":scheme", "http"},
		HeaderField{":path", "/"},
		HeaderField{":authority", "www.example.com"},
		HeaderField{"cache-control", "no-cache"},
	}
	sort.Sort(hl)

	expected := &HeaderList{
		HeaderField{":authority", "www.example.com"},
		HeaderField{":method", "GET"},
		HeaderField{":path", "/"},
		HeaderField{":scheme", "http"},
		HeaderField{"cache-control", "no-cache"},
	}

	assert.Equal(t, hl, expected)
}
