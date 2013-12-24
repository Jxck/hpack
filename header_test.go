package hpack

import (
	"testing"
)

func TestHeaderFieldSize(t *testing.T) {
	h := NewHeaderField("hello", "world")
	if h.Size() != 42 {
		t.Errorf("got %v\nwant %v", h.Size(), 42)
	}
}
