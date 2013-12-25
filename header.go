package hpack

import (
	"strings"
)

// method, scheme, host, path, status
// are must and needs ":" prefix
var MustHeader = map[string]string{
	"scheme": ":scheme",
	"method": ":method",
	"path":   ":path",
	"host":   ":host",
	"status": ":status",
}

// A name-value pair.
// Both name and value are sequences of octets.
type HeaderField struct {
	Name  string
	Value string
}

// Add prefix if name is Must Header
func NewHeaderField(name, value string) HeaderField {
	name = strings.ToLower(name)
	mustname, ok := MustHeader[name]
	if ok {
		name = mustname
	}
	return HeaderField{name, value}
}

// The size of an entry is the sum of its name's length in octets
// of its value's length in octets and of 32 octets.
func (h *HeaderField) Size() int {
	return len(h.Name) + len(h.Value) + 32
}
