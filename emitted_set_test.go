package hpack

import (
	"reflect"
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

func TestCheck(t *testing.T) {
	for _, name := range []string{"Scheme", ":scheme", "scheme"} {
		es := NewEmittedSet()
		es.Emit(name, "http")
		if !es.Check("Scheme", "http") {
			t.Errorf("got %v should have (%q, %q)", es, "Scheme", "http")
		}
	}
}

func TestRemovePrefix(t *testing.T) {
	fixture := []string{":foo", ":foo:", "::foo", ":foo:foo", "bar"}
	expected := []string{"foo", "foo:", "foo", "foo:foo", "bar"}
	actual := make([]string, len(fixture))
	for i, name := range fixture {
		actual[i] = RemovePrefix(name)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
