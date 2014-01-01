package hpack

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRemovePrefix(t *testing.T) {
	var params = []struct {
		fixture, expected string
	}{
		{":foo", "foo"},
		{":foo:", "foo:"},
		{"::foo", "foo"},
		{":foo:foo", "foo:foo"},
		{"bar", "bar"},
	}

	for _, param := range params {
		if RemovePrefix(param.fixture) != param.expected {
			t.Errorf("got %v\nwant %v", RemovePrefix(param.fixture), param.expected)
		}
	}
}

func TestHeaderToHeaderSet(t *testing.T) {
	header := http.Header{
		"Method":      []string{"GET"},
		"Scheme":      []string{"http"},
		"Host":        []string{"example.com"},
		"Path":        []string{"/index.html"},
		"Accept":      []string{"*/*"},
		"Mynewheader": []string{"first", "second"},
	}

	headerSet := HeaderSet{
		NewHeaderField(":method", "GET"),
		NewHeaderField(":scheme", "http"),
		NewHeaderField(":host", "example.com"),
		NewHeaderField(":path", "/index.html"),
		NewHeaderField("accept", "*/*"),
		NewHeaderField("mynewheader", "first,second"),
	}

	actual := HeaderToHeaderSet(header)
	expected := headerSet
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
