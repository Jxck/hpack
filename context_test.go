package hpack

import (
	"net/http"
	"testing"
)

/*
func TestIndexedHeader(t *testing.T) {
	frame := NewIndexedHeader(0)

	server := NewRequestContext()
	server.Decode(frame.Encode().Bytes())
	actual := server.EmittedSet.Get("Scheme")
	if actual != "http" {
		t.Errorf("got %v\nwant %v", actual, "http")
	}
}

func TestIndexedNameWithoutIndexing(t *testing.T) {
	frame := NewIndexedNameWithoutIndexing(0, "ftp")

	server := NewRequestContext()
	server.Decode(frame.Encode().Bytes())
	actual := server.EmittedSet.Get("Scheme")
	if actual != "ftp" {
		t.Errorf("got %v\nwant %v", actual, "ftp")
	}
}

func TestNewNameWithoutIndexing(t *testing.T) {
	frame := NewNewNameWithoutIndexing("x-hello", "world")

	server := NewRequestContext()
	server.Decode(frame.Encode().Bytes())
	actual := server.EmittedSet.Get("X-Hello")
	if actual != "world" {
		t.Errorf("got %v\nwant %v", actual, "ftp")
	}
}

// TODO: check refset, emmitedset in test
func TestIncrementalIndexingWithIndexedName(t *testing.T) {
	frame := NewIndexedNameWithIncrementalIndexing(0, "ftp")

	server := NewRequestContext()
	server.Decode(frame.Encode().Bytes())
	last := len(server.HeaderTable.Headers) - 1
	if server.HeaderTable.Headers[last].Value != "ftp" {
		t.Errorf("got %v\nwant %v", server.HeaderTable.Headers[last].Name, "ftp")
	}
}

func TestIncrementalIndexingWithNewName(t *testing.T) {
	frame := NewNewNameWithIncrementalIndexing("x-hello", "world")

	server := NewRequestContext()
	server.Decode(frame.Encode().Bytes())
	last := server.HeaderTable.Headers[len(server.HeaderTable.Headers)-1]
	if last.Name != "x-hello" || last.Value != "world" {
		t.Errorf("got (%v, %v)\nwant (%v %v)", last.Name, last.Value, "x-hello", "world")
	}
}

func TestContextEncodeDecode(t *testing.T) {
	var headers = http.Header{
		"Scheme":     []string{"https"},
		"Host":       []string{"example.com"},
		"Path":       []string{"/"},
		"Method":     []string{"GET"},
		"User-Agent": []string{"hpack-test"},
		"Cookie":     []string{"xxxxxxx1"},
		"X-Hello":    []string{"world"},
	}

	client := NewRequestContext()
	wire := client.Encode(headers)

	server := NewRequestContext()
	server.Decode(wire)

	headers = http.Header{
		"Scheme":     []string{"https"},
		"Host":       []string{"example.com"},
		"Path":       []string{"/"},
		"Method":     []string{"GET"},
		"User-Agent": []string{"hpack-test"},
		"Cookie":     []string{"xxxxxxx2"},
	}

	wire = client.Encode(headers)
	server.Decode(wire)

	for name, values := range server.EmittedSet.Header {
		if !reflect.DeepEqual(headers[name], values) {
			t.Errorf("got %v\nwant %v", values, headers[name])
		}
	}
}

*/

func TestScenario(t *testing.T) {
	buf := []byte{
		0x82, 0x87, 0x86, 0x04,
		0x0f, 0x77, 0x77, 0x77,
		0x2e, 0x65, 0x78, 0x61,
		0x6d, 0x70, 0x6c, 0x65,
		0x2e, 0x63, 0x6f, 0x6d,
	}
	client := NewContext()
	client.Decode(buf)

	expected := http.Header{
		"Method":    []string{"GET"},
		"Scheme":    []string{"http"},
		"Path":      []string{"/"},
		"Authority": []string{"www.example.com"},
	}

	_ = expected
	//if !reflect.DeepEqual(client.EmittedSet.Header, expected) {
	//	t.Errorf("got %v\nwant %v", client.EmittedSet.Header, expected)
	//}
}
