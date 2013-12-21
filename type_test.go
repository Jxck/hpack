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

func TestNewNewNameWithoutIndexing(t *testing.T) {
	var name string = "foo"
	var value string = "var"
	var frame *NewNameWithoutIndexing
	frame = NewNewNameWithoutIndexing(name, value)

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

func TestNewIndexedNameWithoutIndexing(t *testing.T) {
	var index uint64 = 10
	var value string = "var"
	var frame *IndexedNameWithoutIndexing
	frame = NewIndexedNameWithoutIndexing(index, value)

	actual, expected := frame.Index, index
	if actual != expected {
		t.Errorf("actual = %v\nexpected = %v", actual, expected)
	}

	actual_str, expected_str := frame.ValueString, value
	if actual_str != expected_str {
		t.Errorf("actual = %v\nexpected = %v", actual_str, expected_str)
	}

	actual_len, expected_len := frame.ValueLength, uint64(len(value))
	if actual_len != expected_len {
		t.Errorf("actual_len = %v\nexpected = %v", actual_len, expected_len)
	}
}

func TestNewIndexedNameWithIncrementalIndexing(t *testing.T) {
	var index uint64 = 10
	var value string = "var"
	var frame *IndexedNameWithIncrementalIndexing
	frame = NewIndexedNameWithIncrementalIndexing(index, value)

	actual, expected := frame.Index, index
	if actual != expected {
		t.Errorf("actual = %v\nexpected = %v", actual, expected)
	}

	actual_str, expected_str := frame.ValueString, value
	if actual_str != expected_str {
		t.Errorf("actual = %v\nexpected = %v", actual_str, expected_str)
	}

	actual_len, expected_len := frame.ValueLength, uint64(len(value))
	if actual_len != expected_len {
		t.Errorf("actual_len = %v\nexpected = %v", actual_len, expected_len)
	}
}

func TestNewNewNameWithIncrementalIndexing(t *testing.T) {
	var name string = "foo"
	var value string = "var"
	var frame *NewNameWithIncrementalIndexing
	frame = NewNewNameWithIncrementalIndexing(name, value)

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
