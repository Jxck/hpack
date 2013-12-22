package hpack

// A name-value pair.
// Both name and value are sequences of octets.
type HeaderField struct {
	Name  string
	Value string
}

func NewHeaderField(name, value string) HeaderField {
	return HeaderField{name, value}
}

func (h *HeaderField) Size() int {
	return len(h.Name) + len(h.Value) + 32
}
