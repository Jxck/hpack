package hpack

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// The header table is a component used to associate stored header
// fields to index values.
// The data stored in this table is in first-in, first-out order.
type HeaderTable struct {
	HEADER_TABLE_SIZE uint32
	HeaderFields      []*HeaderField
}

func NewHeaderTable(SETTINGS_HEADER_TABLE_SIZE uint32) *HeaderTable {
	return &HeaderTable{
		SETTINGS_HEADER_TABLE_SIZE,
		make([]*HeaderField, 0),
	}
}

// get total size of Header Table
func (ht *HeaderTable) Size() uint32 {
	var size uint32
	for _, h := range ht.HeaderFields {
		size += h.Size()
	}
	return size
}

// get length of Header Table
func (ht *HeaderTable) Len() int {
	return len(ht.HeaderFields)
}

// push new Header Field to top of HeaderTable
// with eviction
func (ht *HeaderTable) Push(hf *HeaderField) {
	tmp := []*HeaderField{hf}
	ht.HeaderFields = append(tmp, ht.HeaderFields...)
}

// remove Header at index i
func (ht *HeaderTable) Remove(index int) *HeaderField {
	// https://code.google.com/p/go-wiki/wiki/SliceTricks
	copy(ht.HeaderFields[index:], ht.HeaderFields[index+1:])
	// avoid memory leak
	removed := ht.HeaderFields[len(ht.HeaderFields)-1]
	ht.HeaderFields[len(ht.HeaderFields)-1] = new(HeaderField) // GC
	ht.HeaderFields = ht.HeaderFields[:len(ht.HeaderFields)-1]
	return removed
}

// String for Debug
func (ht *HeaderTable) String() (str string) {
	str += fmt.Sprintf("\n--------- HT(%v/%v) ---------\n",
		ht.Size(), ht.HEADER_TABLE_SIZE)
	for i, v := range ht.HeaderFields {
		str += fmt.Sprintln(i, v)
	}
	str += "--------------------------------\n"
	return str
}
