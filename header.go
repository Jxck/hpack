package hpack

type Header struct {
	Name  string
	Value string
}

func (h *Header) Size() int {
	return len(h.Name) + len(h.Value) + 32
}
