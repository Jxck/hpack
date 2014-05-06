package hpack

import (
	assert "github.com/jxck/assertion"
	"net/http"
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
	assert.Equal(t, actual, expected)
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
	assert.Equal(t, actual, header)
}

func TestEmit(t *testing.T) {
	hs := NewHeaderSet()
	hf1 := NewHeaderField("key1", "value1")
	hf2 := NewHeaderField("key2", "value2")
	hs.Emit(hf1)
	hs.Emit(hf2)
	assert.Equal(t, hs.Len(), 2)
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

	expected := &HeaderSet{
		HeaderField{":authority", "www.example.com"},
		HeaderField{":method", "GET"},
		HeaderField{":path", "/"},
		HeaderField{":scheme", "http"},
		HeaderField{"cache-control", "no-cache"},
	}

	assert.Equal(t, hs, expected)
}
