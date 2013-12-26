package hpack

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

func (rs *ReferenceSet) Add(hf *HeaderField, emitted bool) {
	rf := &ReferencedField{hf, emitted}
	*rs = append(*rs, rf)
}

func (rs *ReferenceSet) Empty() {
	*rs = ReferenceSet{}
}

func (rs *ReferenceSet) Has(hf *HeaderField) bool {
	for _, rf := range *rs {
		if hf == rf.HeaderField {
			return true
		}
	}
	return false
}

func (rs *ReferenceSet) Remove(hf *HeaderField) bool {
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

func (rs *ReferenceSet) Reset() {
	for _, rf := range *rs {
		rf.Emitted = false
	}
}
