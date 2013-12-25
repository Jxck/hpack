package hpack

import (
	"testing"
)

func TestReferenceSetAdd(t *testing.T) {
	rs := NewReferenceSet()
	hf := NewHeaderField("key", "value")
	rs.Add(hf)
	if rs.Len() != 1 {
		t.Errorf("\ngot  %v\nwant %v", rs.Len(), 1)
	}
}

func TestReferenceSetEmpty(t *testing.T) {
	rs := NewReferenceSet()
	hf := NewHeaderField("key", "value")
	rs.Add(hf)
	rs.Empty()
	if rs.Len() != 0 {
		t.Errorf("\ngot  %v\nwant %v", rs.Len(), 0)
	}
}

func TestReferenceSetHas(t *testing.T) {
	rs := NewReferenceSet()
	ref1 := NewHeaderField("key1", "value1")
	ref2 := NewHeaderField("key2", "value2")
	ref3 := NewHeaderField("key3", "value3")
	ref4 := NewHeaderField("key4", "value4")

	rs.Add(ref1)
	rs.Add(ref2)
	rs.Add(ref3)

	if !(rs.Has(ref1) && rs.Has(ref2) && rs.Has(ref3)) {
		t.Errorf("RefSet.Has() faild at %v", rs)
	}

	if rs.Has(ref4) {
		t.Errorf("RefSet shouldn't has %v", ref4)
	}
}

func TestReferenceSetRemove(t *testing.T) {
	rs := NewReferenceSet()
	ref1 := NewHeaderField("key1", "value1")
	ref2 := NewHeaderField("key2", "value2")
	ref3 := NewHeaderField("key3", "value3")
	ref4 := NewHeaderField("key4", "value4")

	rs.Add(ref1)
	rs.Add(ref2)
	rs.Add(ref3)

	ok := rs.Remove(ref2)
	if !ok {
		t.Errorf("remove faild")
	}

	if rs.Has(ref2) {
		t.Errorf("\ngot  %v\nwant %v", rs.Has(ref2), false)
	}

	ok = rs.Remove(ref4)
	if ok {
		t.Errorf("unexpected remove")
	}

	// remove all
	rs.Remove(ref1)
	rs.Remove(ref3)

	if rs.Len() != 0 {
		t.Errorf("\ngot  %v\nwant %v", rs.Len(), 0)
	}
}
