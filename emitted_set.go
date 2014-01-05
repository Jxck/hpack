package hpack

func NewHeaderSet() *HeaderSet {
	return &HeaderSet{}
}

func (e *HeaderSet) Emit(hf *HeaderField) {
	*e = append(*e, *hf)
}

func (e *HeaderSet) Len() int {
	return len(*e)
}

// Sort Interface
func (e *HeaderSet) Swap(i, j int) {
	es := *e
	es[i], es[j] = es[j], es[i]
}

func (e *HeaderSet) Less(i, j int) bool {
	es := *e
	if es[i].Name == es[j].Name {
		return es[i].Value < es[j].Value
	}
	return es[i].Name < es[j].Name
}
