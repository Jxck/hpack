package hpack

import (
	"net/http"
	"strings"
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

// remove ":" prefix
func RemovePrefix(name string) string {
	if strings.HasPrefix(name, ":") {
		name = strings.TrimLeft(name, ":")
	}
	return name
}

// http.Header => HeaderSet
func HeaderToHeaderSet(header http.Header) HeaderSet {
	headerSet := make(HeaderSet, 0, len(header))
	for name, values := range header {
		// process name
		name = strings.ToLower(name)
		mustname, ok := MustHeader[name]
		if ok {
			name = mustname
		}
		// process values
		value := strings.Join(values, ",")
		headerField := NewHeaderField(name, value)
		headerSet = append(headerSet, headerField)
	}
	return headerSet
}

// HeaderSet => http.Header
// But, multi value in single key like
// myheader: ["first", "second"]
// becames
// myheader: ["first,second"]
func HeaderSetToHeader(headerSet HeaderSet) http.Header {
	headers := make(http.Header, len(headerSet))
	for _, headerField := range headerSet {
		name := RemovePrefix(headerField.Name)
		value := headerField.Value
		headers.Add(name, value)
	}
	return headers
}
