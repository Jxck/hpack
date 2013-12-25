package hpack

// an unordered set of references to entries of the header table.
type ReferenceSet []*HeaderField

func NewReferenceSet() *ReferenceSet {
	return &ReferenceSet{}
}

func (rs *ReferenceSet) Len() int {
	return len(*rs)
}

func (rs *ReferenceSet) Add(ref *HeaderField) {
	*rs = append(*rs, ref)
}

func (rs *ReferenceSet) Empty() {
	*rs = ReferenceSet{}
}

func (rs *ReferenceSet) Has(ref *HeaderField) bool {
	for _, idx := range *rs {
		if idx == ref {
			return true
		}
	}
	return false
}

func (rs *ReferenceSet) Remove(ref *HeaderField) bool {
	for i, idx := range *rs {
		if idx == ref {
			tmp := *rs
			copy(tmp[i:], tmp[i+1:])
			*rs = tmp[:len(tmp)-1]
			return true
		}
	}
	return false
}
