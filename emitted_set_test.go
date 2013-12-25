package hpack

import (
	"testing"
)

func TestEmit(t *testing.T) {
	es := NewEmittedSet()
	hf := NewHeaderField("a", "b")
	es.Emit(hf)

	if es.Get("A") != "b" {
		t.Errorf("got %v\nwant %v", es.Get("A"), "b")
	}
}
