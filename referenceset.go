package hpack

import (
	"fmt"
)

const (
	EMITTED     bool = true
	NOT_EMITTED      = false
)

// Add Emitted Flag to HeaderField
type ReferencedField struct {
	*HeaderField
	Emitted bool
}

// an unordered set of references to entries of the header table.
type ReferenceSet []*ReferencedField

func NewReferenceSet() *ReferenceSet {
	return &ReferenceSet{}
}

func (rs *ReferenceSet) Len() int {
	return len(*rs)
}

// add header field with given emitted flag
func (rs *ReferenceSet) Add(hf *HeaderField, emitted bool) {
	rf := &ReferencedField{hf, emitted}
	*rs = append(*rs, rf)
}

// cleanup reference set
func (rs *ReferenceSet) Empty() {
	*rs = ReferenceSet{}
}

// search given reference is exists in refset or not
func (rs *ReferenceSet) Has(hf *HeaderField) bool {
	for _, rf := range *rs {
		if hf == rf.HeaderField {
			return true
		}
	}
	return false
}

// remove given reference from refset
func (rs *ReferenceSet) Remove(hf *HeaderField) bool {
	// search as same as Has() dose
	// not use Has() because it dosen't returns index
	// and I don't want to return (bool, int) as Has()
	// for 'if' statement
	for i, rf := range *rs {
		if hf == rf.HeaderField {
			tmp := *rs
			copy(tmp[i:], tmp[i+1:])
			*rs = tmp[:len(tmp)-1]
			return true
		}
	}
	return false
}

// set all entry as "Not Emitted"
func (rs *ReferenceSet) Reset() {
	for _, rf := range *rs {
		rf.Emitted = NOT_EMITTED
	}
}

// Dump for Debug
func (rs *ReferenceSet) Dump() (str string) {
	str += "\n-------------- RS --------------\n"
	for i, v := range *rs {
		str += fmt.Sprintln(i, *v.HeaderField, v.Emitted)
	}
	str += "--------------------------------\n"
	return str
}
