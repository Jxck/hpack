package hpack

import (
	"testing"
)

func TestReferenceSetAdd(t *testing.T) {
	rs := NewReferenceSet()
	rs.Add(1)
	if rs.Len() != 1 {
		t.Errorf("\ngot  %v\nwant %v", rs.Len(), 1)
	}
}

func TestReferenceSetEmpty(t *testing.T) {
	rs := NewReferenceSet()
	rs.Add(1)
	rs.Empty()
	if rs.Len() != 0 {
		t.Errorf("\ngot  %v\nwant %v", rs.Len(), 0)
	}
}
