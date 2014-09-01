package hpack

import (
	"fmt"
	"net/http"
	"strings"
)

// A header set is a potentially ordered group of header fields that are encoded jointly.
// A complete set of key-value pairs contained in a HTTP request or response is a header set.
type HeaderList []HeaderField

func NewHeaderList() *HeaderList {
	return new(HeaderList)
}

func ToHeaderList(header http.Header) HeaderList {
	hs := *new(HeaderList)
	for key, values := range header {
		key := strings.ToLower(key)
		for _, value := range values {
			hs = append(hs, *NewHeaderField(key, value))
		}
	}
	return hs
}

func (hs *HeaderList) Emit(hf *HeaderField) {
	*hs = append(*hs, *hf)
}

func (hs *HeaderList) Len() int {
	return len(*hs)
}

// Sort Interface
func (hs *HeaderList) Swap(i, j int) {
	h := *hs
	h[i], h[j] = h[j], h[i]
}

func (hs *HeaderList) Less(i, j int) bool {
	h := *hs
	if h[i].Name == h[j].Name {
		return h[i].Value < h[j].Value
	}
	return h[i].Name < h[j].Name
}

// convert to http.Header
func (hs HeaderList) ToHeader() http.Header {
	header := make(http.Header)
	for _, hf := range hs {
		header.Add(hf.Name, hf.Value)
	}
	return header
}

func (hs HeaderList) String() (str string) {
	str += fmt.Sprintf("\n--------- HS ---------\n")
	for i, v := range hs {
		str += fmt.Sprintln(i, v)
	}
	str += "--------------------------------\n"
	return str
}
