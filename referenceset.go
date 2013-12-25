package hpack

// an unordered set of references to entries of the header table.
type ReferenceSet []int

func NewReferenceSet() *ReferenceSet {
	return &ReferenceSet{}
}

func (rs *ReferenceSet) Len() int {
	return len(*rs)
}

func (rs *ReferenceSet) Add(index int) {
	*rs = append(*rs, index)
}

func (rs *ReferenceSet) Empty() {
	*rs = ReferenceSet{}
}

func (rs *ReferenceSet) Has(index int) bool {
	for _, idx := range *rs {
		if idx == index {
			return true
		}
	}
	return false
}
