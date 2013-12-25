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

func TestReferenceSetHas(t *testing.T) {
	rs := NewReferenceSet()
	rs.Add(1)
	rs.Add(2)
	rs.Add(4)

	expected := []bool{
		false, // 0
		true,  // 1
		true,  // 2
		false, // 3
		true,  // 4
		false, // 5
	}

	for i, b := range expected {
		if rs.Has(i) != b {
			t.Errorf("\ngot  %v\nwant %v", rs.Has(i), i)
		}
	}
}
