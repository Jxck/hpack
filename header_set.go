package hpack

import (
	"fmt"
	"net/http"
	"strings"
)

// A header set is a potentially ordered group of header fields that are encoded jointly.
// A complete set of key-value pairs contained in a HTTP request or response is a header set.
type HeaderSet []HeaderField

func NewHeaderSet() *HeaderSet {
	return new(HeaderSet)
}

func ToHeaderSet(header http.Header) HeaderSet {
	hs := *new(HeaderSet)
	for key, values := range header {
		key := strings.ToLower(key)
		for _, value := range values {
			hs = append(hs, *NewHeaderField(key, value))
		}
	}
	return hs
}

func (hs *HeaderSet) Emit(hf *HeaderField) {
	*hs = append(*hs, *hf)
}

func (hs *HeaderSet) Len() int {
	return len(*hs)
}

// Sort Interface
func (hs *HeaderSet) Swap(i, j int) {
	h := *hs
	h[i], h[j] = h[j], h[i]
}

func (hs *HeaderSet) Less(i, j int) bool {
	h := *hs
	if h[i].Name == h[j].Name {
		return h[i].Value < h[j].Value
	}
	return h[i].Name < h[j].Name
}

// convert to http.Header
func (hs HeaderSet) ToHeader() http.Header {
	header := make(http.Header)
	for _, hf := range hs {
		header.Add(hf.Name, hf.Value)
	}
	return header
}

func (hs HeaderSet) String() (str string) {
	str += fmt.Sprintf("\n--------- HS ---------\n")
	for i, v := range hs {
		str += fmt.Sprintln(i, v)
	}
	str += "--------------------------------\n"
	return str
}
