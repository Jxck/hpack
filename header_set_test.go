package hpack

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

var header = http.Header{
	"Method":      []string{"GET"},
	"Scheme":      []string{"http"},
	"Host":        []string{"example.com"},
	"Path":        []string{"/index.html"},
	"Accept":      []string{"*/*"},
	"Mynewheader": []string{"first", "second"},
}

var headerSet = HeaderSet{
	HeaderField{":method", "GET"},
	HeaderField{":scheme", "http"},
	HeaderField{":host", "example.com"},
	HeaderField{":path", "/index.html"},
	HeaderField{"accept", "*/*"},
	HeaderField{"mynewheader", "first,second"},
}

func TestHeaderSetToHeader(t *testing.T) {
	actual := headerSet.ToHeader()
	expected := header
	// But, multi value in single key like
	// myheader: ["first", "second"]
	// becames
	// myheader: ["first,second"]
	joined := []string{strings.Join(expected["Mynewheader"], ",")}
	expected["Mynewheader"] = joined
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
