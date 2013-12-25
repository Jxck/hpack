package hpack

import (
	"testing"
)

func TestNewHeaderField(t *testing.T) {
	var name, value string
	var hf *HeaderField
	// Normal Header
	name, value = "cookie", "xxxx"
	hf = NewHeaderField(name, value)

	if hf.Name != "cookie" && hf.Value != "xxx" {
		t.Errorf("got %v\nwant %v", hf.Name, hf.Value)
	}

	// Must Header
	name, value = "scheme", "http"
	hf = NewHeaderField(name, value)

	if hf.Name != ":scheme" && hf.Value != "http" {
		t.Errorf("got %v\nwant %v", hf.Name, hf.Value)
	}
}

func TestHeaderFieldSize(t *testing.T) {
	h := NewHeaderField("hello", "world")
	if h.Size() != 42 {
		t.Errorf("got %v\nwant %v", h.Size(), 42)
	}
}
