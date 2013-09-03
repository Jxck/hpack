package hpac

import (
	"testing"
)

func TestCreateIndexedHeader(t *testing.T) {
	var index uint64 = 10
	var frame *IndexedHeader
	frame = CreateIndexedHeader(index)

	actual := frame.Index
	expected := index
	if actual != expected {
		t.Errorf("actual = %v\nexpected = %v", actual, expected)
	}
}
