package hpack

import (
	"sort"
	"testing"
)

func TestEmit(t *testing.T) {
	es := NewHeaderSet()
	hf1 := NewHeaderField("key1", "value1")
	hf2 := NewHeaderField("key2", "value2")
	es.Emit(hf1)
	es.Emit(hf2)
	if es.Len() != 2 {
		t.Fatal("%v should length %v", es, 2)
	}
}

func TestEmitSort(t *testing.T) {
	es := &HeaderSet{
		HeaderField{":method", "GET"},
		HeaderField{":scheme", "http"},
		HeaderField{":path", "/"},
		HeaderField{":authority", "www.example.com"},
		HeaderField{"cache-control", "no-cache"},
	}

	sort.Sort(es)
	if (*es)[0].Name != ":authority" {
		t.Fatal("*es[0] should length %v but %v", ":authority", (*es)[0])
	}
}
