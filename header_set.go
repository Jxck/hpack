package hpack

import (
	"net/http"
	"strings"
)

type HeaderSet map[string]string

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
	headerSet := make(HeaderSet, len(header))
	for name, value := range header {
		name = strings.ToLower(name)
		mustname, ok := MustHeader[name]
		if ok {
			name = mustname
		}
		headerSet[name] = strings.Join(value, ",")
	}
	return headerSet
}

// HeaderSet => http.Header
func HeaderSetToHeader(headerset HeaderSet) http.Header {
	headers := make(http.Header, len(headerset))
	for name, value := range headerset {
		name = RemovePrefix(name)
		headers.Add(name, value)
	}
	return headers
}
