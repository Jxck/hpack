package hpack

import (
	"fmt"
	"net/http"
	"strings"
)

// A header set is a potentially ordered group of header fields that are encoded jointly.
// A complete set of key-value pairs contained in a HTTP request or response is a header set.
type HeaderSet []HeaderField

func ToHeaderSet(header http.Header) HeaderSet {
	hs := HeaderSet{}
	for key, values := range header {
		key := strings.ToLower(key)
		for _, value := range values {
			hs = append(hs, *NewHeaderField(key, value))
		}
	}
	return hs
}

func (hs HeaderSet) ToHeader() http.Header {
	header := make(http.Header)
	for _, hf := range hs {
		header.Add(hf.Name, hf.Value)
	}
	return header
}

func (hs HeaderSet) Dump() (str string) {
	str += fmt.Sprintf("\n--------- HS ---------\n")
	for i, v := range hs {
		str += fmt.Sprintln(i, v)
	}
	str += "--------------------------------\n"
	return str
}
