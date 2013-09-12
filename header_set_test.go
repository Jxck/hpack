package hpack

import (
	"net/http"
	"testing"
)

func TestNewHeaderSet(t *testing.T) {
	var headers = http.Header{
		":scheme":     []string{"http"},
		":path":       []string{"/index.html"},
		"mynewheader": []string{"first", "second"},
	}
	actual := NewHeaderSet(headers)

	var expected = HeaderSet{
		":scheme":     "http",
		":path":       "/index.html",
		"mynewheader": "first,second",
	}

	for name, value := range actual {
		if value != expected[name] {
			t.Errorf("got %v\nwant %v", value, expected[name])
		}
	}
}
