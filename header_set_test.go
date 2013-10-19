package hpack

import (
	"net/http"
	"reflect"
	"testing"
)

var headers = http.Header{
	"Method":      []string{"GET"},
	"Scheme":      []string{"http"},
	"Host":        []string{"example.com"},
	"Path":        []string{"/index.html"},
	"Accept":      []string{"*/*"},
	"Mynewheader": []string{"first,second"},
}

var headerset = HeaderSet{
	":method":     "GET",
	":scheme":     "http",
	":host":       "example.com",
	":path":       "/index.html",
	"accept":      "*/*",
	"mynewheader": "first,second",
}

func TestHeaderToHeaderSet(t *testing.T) {
	actual := HeaderToHeaderSet(headers)
	expected := headerset
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", headers, expected)
	}
}

func TestHeaderSetToHeader(t *testing.T) {
	actual := HeaderSetToHeader(headerset)
	expected := headers
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", headers, expected)
	}
}
