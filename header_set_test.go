package hpack

import (
	"net/http"
	"testing"
)

func TestNewHeaderSet(t *testing.T) {
	var headers = http.Header{
		"Method":      []string{"GET"},
		"Scheme":      []string{"http"},
		"Host":        []string{"example.com"},
		"Path":        []string{"/index.html"},
		"Accept":      []string{"*/*"},
		"Mynewheader": []string{"first,second"},
	}

	var expected = HeaderSet{
		":method":     "GET",
		":scheme":     "http",
		":host":       "example.com",
		":path":       "/index.html",
		"accept":      "*/*",
		"mynewheader": "first,second",
	}

	for name, value := range NewHeaderSet(headers) {
		if value != expected[name] {
			t.Errorf("got %v\nwant %v", value, expected[name])
		}
	}
}
