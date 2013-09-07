package hpac

import (
	"net/http"
	"testing"
)

func TestContextEncodeDecode(t *testing.T) {
	var headers = http.Header{
		":scheme":    []string{"https"},
		":host":      []string{"jxck.io"},
		":path":      []string{"/"},
		":method":    []string{"GET"},
		"User-Agent": []string{"http2cat"},
		"Cookie":     []string{"xxxxxxx1"},
		"X-Hello":    []string{"world"},
	}

	client := NewContext()
	wire := client.Encode(headers)

	server := NewContext()
	server.Decode(wire)

	headers = http.Header{
		":scheme":    []string{"https"},
		":host":      []string{"jxck.io"},
		":path":      []string{"/labs/http2cat"},
		":method":    []string{"GET"},
		"User-Agent": []string{"http2cat"},
		"Cookie":     []string{"xxxxxxx2"},
	}

	wire = client.Encode(headers)
	server.Decode(wire)

	for name, values := range server.EmittedSet {
		if !CompareSlice(headers[name], values) {
			t.Errorf("got %v\nwant %v", values, headers[name])
		}
	}
}
