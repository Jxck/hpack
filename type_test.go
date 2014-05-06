package hpack

import (
	assert "github.com/jxck/assertion"
	"testing"
)

func TestNewIndexedHeader(t *testing.T) {
	var index uint64 = 10
	var frame *IndexedHeader
	frame = NewIndexedHeader(index)

	actual, expected := frame.Index, index
	assert.Equal(t, actual, expected)
}

func TestNewIndexedLiteral(t *testing.T) {
	var indexing Indexing = WITH
	var index uint64 = 10
	var value string = "var"
	var frame *IndexedLiteral = NewIndexedLiteral(indexing, index, value)

	assert.Equal(t, frame.Indexing, indexing)
	assert.Equal(t, frame.Index, index)
	assert.Equal(t, frame.ValueLength, uint64(len(value)))
	assert.Equal(t, frame.ValueString, value)
}

func TestNewStringLiteral(t *testing.T) {
	var indexing Indexing = WITH
	var name string = "foo"
	var value string = "var"
	var frame *StringLiteral = NewStringLiteral(indexing, name, value)

	assert.Equal(t, frame.Indexing, indexing)
	assert.Equal(t, frame.NameLength, uint64(len(name)))
	assert.Equal(t, frame.NameString, name)
	assert.Equal(t, frame.ValueLength, uint64(len(value)))
	assert.Equal(t, frame.ValueString, value)
}
