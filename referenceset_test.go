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

func TestReferenceSetRemove(t *testing.T) {
	rs := NewReferenceSet()
	rs.Add(1)
	rs.Add(2)
	rs.Add(4)

	ok := rs.Remove(2)
	if !ok {
		t.Errorf("remove faild")
	}

	if rs.Has(2) {
		t.Errorf("\ngot  %v\nwant %v", rs.Has(2), false)
	}

	ok = rs.Remove(100)
	if ok {
		t.Errorf("unexpected remove")
	}

	// remove all
	rs.Remove(1)
	rs.Remove(4)

	if rs.Len() != 0 {
		t.Errorf("\ngot  %v\nwant %v", rs.Len(), 0)
	}
}
