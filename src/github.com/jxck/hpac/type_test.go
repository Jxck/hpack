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

func TestCreateIndexedNameWithoutIndexing(t *testing.T) {
	var index uint64 = 10
	var value string = "var"
	var frame *IndexedNameWithoutIndexing
	frame = CreateIndexedNameWithoutIndexing(index, value)

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

func TestCreateIndexedNameWithIncrementalIndexing(t *testing.T) {
	var index uint64 = 10
	var value string = "var"
	var frame *IndexedNameWithIncrementalIndexing
	frame = CreateIndexedNameWithIncrementalIndexing(index, value)

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

func TestCreateNewNameWithIncrementalIndexing(t *testing.T) {
	var name string = "foo"
	var value string = "var"
	var frame *NewNameWithIncrementalIndexing
	frame = CreateNewNameWithIncrementalIndexing(name, value)

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

func TestCreateIndexedNameWithSubstitutionIndexing(t *testing.T) {
	var index uint64 = 10
	var substitutedIndex uint64 = 20
	var value string = "var"
	var frame *IndexedNameWithSubstitutionIndexing
	frame = CreateIndexedNameWithSubstitutionIndexing(index, substitutedIndex, value)

	actual, expected := frame.Index, index
	if actual != expected {
		t.Errorf("actual = %v\nexpected = %v", actual, expected)
	}

	actual, expected = frame.SubstitutedIndex, substitutedIndex
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
