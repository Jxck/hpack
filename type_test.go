package hpack

import (
	"testing"
)

func TestNewIndexedHeader(t *testing.T) {
	var index uint64 = 10
	var frame *IndexedHeader
	frame = NewIndexedHeader(index)

	actual, expected := frame.Index, index
	if actual != expected {
		t.Errorf("actual = %v\nexpected = %v", actual, expected)
	}
}

func TestNewIndexedLiteral(t *testing.T) {
	var indexing bool = true
	var index uint64 = 10
	var value string = "var"
	var frame *IndexedLiteral = NewIndexedLiteral(indexing, index, value)

	if frame.Indexing != indexing ||
		frame.Index != index ||
		frame.ValueLength != uint64(len(value)) ||
		frame.ValueString != value {
		t.Errorf(`faild NewIndexedLiteral
							frame      = %v
							---should---
							indexing   = %v
							index      = %v
							value      = %v
							len(value) = %v
						`, frame, indexing, index, value, len(value))
	}
}

func TestNewStringLiteral(t *testing.T) {
	var indexing bool = true
	var name string = "foo"
	var value string = "var"
	var frame *StringLiteral = NewStringLiteral(indexing, name, value)

	if frame.Indexing != indexing ||
		frame.Index != 0 ||
		frame.NameLength != uint64(len(name)) ||
		frame.NameString != name ||
		frame.ValueLength != uint64(len(value)) ||
		frame.ValueString != value {
		t.Errorf(`faild NewStringLiteral
							frame      = %v
							---should---
							indexing   = %v
							index      = %v
							name       = %v
							len(name)  = %v
							value      = %v
							len(value) = %v
						`, frame, indexing, 0, name, len(name), value, len(value))
	}
}
