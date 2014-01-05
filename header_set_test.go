package hpack

import (
	"net/http"
	"testing"
)

func TestNewHeaderSet(t *testing.T) {
	header := make(http.Header)
	header.Add("method", "get")
	header.Add("host", "example.com")
	header.Add(":host", "example.com")
	header.Add("cookie", "a")
	header.Add("cookie", "b")

	expected := HeaderSet{
		{"method", "get"},
		{"host", "example.com"},
		{":host", "example.com"},
		{"cookie", "a"},
		{"cookie", "b"},
	}
	actual := NewHeaderSet(header)

	for i, hf := range expected {
		if !(*(actual[i]) == *hf) {
			t.Errorf("\ngot  %v\nwant %v", actual.Dump(), expected.Dump())
		}
	}
}
