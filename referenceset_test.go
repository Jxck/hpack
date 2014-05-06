package hpack

import (
	assert "github.com/jxck/assertion"
	"testing"
)

func TestReferenceSetAdd(t *testing.T) {
	rs := NewReferenceSet()
	hf := NewHeaderField("key", "value")
	rs.Add(hf, true)
	assert.Equal(t, rs.Len(), 1)
}

func TestReferenceSetEmpty(t *testing.T) {
	rs := NewReferenceSet()
	hf := NewHeaderField("key", "value")
	rs.Add(hf, true)
	rs.Empty()
	assert.Equal(t, rs.Len(), 0)
}

func TestReferenceSetHas(t *testing.T) {
	rs := NewReferenceSet()
	ref1 := NewHeaderField("key1", "value1")
	ref2 := NewHeaderField("key2", "value2")
	ref3 := NewHeaderField("key3", "value3")
	ref4 := NewHeaderField("key4", "value4")

	rs.Add(ref1, true)
	rs.Add(ref2, true)
	rs.Add(ref3, true)

	assert.Equal(t, rs.Has(ref1) && rs.Has(ref2) && rs.Has(ref3), true)
	assert.Equal(t, rs.Has(ref4), false)
}

func TestReferenceSetRemove(t *testing.T) {
	rs := NewReferenceSet()
	ref1 := NewHeaderField("key1", "value1")
	ref2 := NewHeaderField("key2", "value2")
	ref3 := NewHeaderField("key3", "value3")
	ref4 := NewHeaderField("key4", "value4")

	rs.Add(ref1, true)
	rs.Add(ref2, true)
	rs.Add(ref3, true)

	ok := rs.Remove(ref2)
	if !ok {
		t.Errorf("remove faild")
	}

	assert.Equal(t, rs.Has(ref2), false)

	ok = rs.Remove(ref4)
	if ok {
		t.Errorf("unexpected remove")
	}

	// remove all
	rs.Remove(ref1)
	rs.Remove(ref3)

	assert.Equal(t, rs.Len(), 0)
}
