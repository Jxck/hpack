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

/*
	es.Add("foo", "bar")
	es.Set("foo", "buz")
	es.Del("foo")
	for name, value := range es.Header {
		log.Println(name, value)
	}
*/
