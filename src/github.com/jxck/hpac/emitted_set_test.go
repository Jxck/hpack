package hpac

import (
	"testing"
)

func TestEmit(t *testing.T) {
	es := NewEmittedSet()
	es.Emit(":a", "b")

	if es.Get("A") != "b" {
		t.Errorf("got %v\nwant %v", es.Get("A"), "b")
	}
}

func TestInheriHttpHeader(t *testing.T) {
	es := NewEmittedSet()
	es.Add("foo", "bar")
	es.Set("foo", "buz")
	es.Del("foo")
	if len(es.Header) != 0 {
		t.Errorf("got %v\nwant %v", len(es.Header), 0)
	}
}
