package hpack

import (
	"net/http"
)

// A header set is a potentially ordered group of header fields that are encoded jointly.
// A complete set of key-value pairs contained in a HTTP request or response is a header set.
// TODO: make it slice of pointer ?
type HeaderSet []HeaderField

// method, scheme, host, path, status
// are must and needs ":" prefix
var MustHeader = map[string]string{
	"scheme": ":scheme",
	"method": ":method",
	"path":   ":path",
	"host":   ":host",
	"status": ":status",
}

// HeaderSet => http.Header
// But, multi value in single key like
// myheader: ["first", "second"]
// becames
// myheader: ["first,second"]
func (headerSet HeaderSet) ToHeader() http.Header {
	headers := make(http.Header, len(headerSet))
	for _, headerField := range headerSet {
		name := RemovePrefix(headerField.Name)
		value := headerField.Value
		headers.Add(name, value)
	}
	return headers
}
