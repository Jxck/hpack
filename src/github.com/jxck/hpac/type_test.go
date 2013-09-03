package hpac

import (
	"testing"
)

func TestCreateIndexedHeader(t *testing.T) {
	var index uint64 = 10
	var frame *IndexedHeader
	frame = CreateIndexedHeader(index)

	actual, expected := frame.Index, index
	if actual != expected {
		t.Errorf("actual = %v\nexpected = %v", actual, expected)
	}
}

func TestCreateNewNameWithoutIndexing(t *testing.T) {
	var name string = "foo"
	var value string = "var"
	var frame *NewNameWithoutIndexing
	frame = CreateNewNameWithoutIndexing(name, value)

	actual_len, expected_len := frame.NameLength, uint64(len(name))
	if actual_len != expected_len {
		t.Errorf("actual_len = %v\nexpected = %v", actual_len, expected_len)
	}

	actual, expected := frame.NameString, name
	if actual != expected {
		t.Errorf("actual = %v\nexpected = %v", actual, expected)
	}

	actual_len, expected_len = frame.ValueLength, uint64(len(value))
	if actual_len != expected_len {
		t.Errorf("actual_len = %v\nexpected = %v", actual_len, expected_len)
	}

	actual, expected = frame.ValueString, value
	if actual != expected {
		t.Errorf("actual = %v\nexpected = %v", actual, expected)
	}
}
