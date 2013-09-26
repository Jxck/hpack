package hpack

import (
	"net/http"
	"testing"
)

// TODO: check refset, emmitedset in test
func TestIncrementalIndexingWithIndexedName(t *testing.T) {
	frame := CreateIndexedNameWithIncrementalIndexing(0, "ftp")

	server := NewRequestContext()
	server.Decode(frame.Encode().Bytes())
	last := len(server.HeaderTable.Headers) - 1
	if server.HeaderTable.Headers[last].Value != "ftp" {
		t.Errorf("got %v\nwant %v", server.HeaderTable.Headers[last].Name, "ftp")
	}
}

func TestIncrementalIndexingWithNewName(t *testing.T) {
	frame := CreateNewNameWithIncrementalIndexing("x-hello", "world")

	server := NewRequestContext()
	server.Decode(frame.Encode().Bytes())
	last := server.HeaderTable.Headers[len(server.HeaderTable.Headers)-1]
	if last.Name != "x-hello" || last.Value != "world" {
		t.Errorf("got (%v, %v)\nwant (%v %v)", last.Name, last.Value, "x-hello", "world")
	}
}

func TestSubstitutionIndexingWithIndexedName(t *testing.T) {
	frame := CreateIndexedNameWithSubstitutionIndexing(0, 10, "ftp")

	server := NewRequestContext()
	server.Decode(frame.Encode().Bytes())
	target := server.HeaderTable.Headers[10]
	if target.Name != ":scheme" || target.Value != "ftp" {
		t.Errorf("got (%v, %v)\nwant (%v %v)", target.Name, target.Value, ":scheme", "ftp")
	}
}

func TestSubstitutionIndexingWithNewName(t *testing.T) {
	frame := CreateNewNameWithSubstitutionIndexing("x-hello", 10, "world")

	server := NewRequestContext()
	server.Decode(frame.Encode().Bytes())
	target := server.HeaderTable.Headers[10]
	if target.Name != "x-hello" || target.Value != "world" {
		t.Errorf("got (%v, %v)\nwant (%v %v)", target.Name, target.Value, "x-hello", "world")
	}
}

func TestContextEncodeDecode(t *testing.T) {
	var headers = http.Header{
		"Scheme":     []string{"https"},
		"Host":       []string{"jxck.io"},
		"Path":       []string{"/"},
		"Method":     []string{"GET"},
		"User-Agent": []string{"http2cat"},
		"Cookie":     []string{"xxxxxxx1"},
		"X-Hello":    []string{"world"},
	}

	client := NewRequestContext()
	wire := client.Encode(headers)

	server := NewRequestContext()
	server.Decode(wire)

	headers = http.Header{
		"Scheme":     []string{"https"},
		"Host":       []string{"jxck.io"},
		"Path":       []string{"/"},
		"Method":     []string{"GET"},
		"User-Agent": []string{"http2cat"},
		"Cookie":     []string{"xxxxxxx2"},
	}

	wire = client.Encode(headers)
	server.Decode(wire)

	for name, values := range server.EmittedSet.Header {
		if !CompareSlice(headers[name], values) {
			t.Errorf("got %v\nwant %v", values, headers[name])
		}
	}
}

func TestSinerio(t *testing.T) {
	buf := []byte{128, 66, 1, 48, 69, 9, 116, 101, 120, 116, 47, 104, 116, 109, 108, 70, 29, 84, 104, 117, 44, 32, 50, 54, 32, 83, 101, 112, 32, 50, 48, 49, 51, 32, 48, 54, 58, 52, 51, 58, 53, 50, 32, 71, 77, 84, 71, 26, 34, 49, 55, 57, 98, 49, 45, 49, 57, 98, 48, 45, 52, 101, 54, 53, 55, 51, 51, 54, 97, 55, 50, 56, 48, 34, 73, 29, 83, 97, 116, 44, 32, 49, 52, 32, 83, 101, 112, 32, 50, 48, 49, 51, 32, 49, 50, 58, 51, 53, 58, 48, 54, 32, 71, 77, 84, 74, 22, 65, 112, 97, 99, 104, 101, 47, 50, 46, 50, 46, 50, 50, 32, 40, 85, 98, 117, 110, 116, 117, 41, 76, 15, 65, 99, 99, 101, 112, 116, 45, 69, 110, 99, 111, 100, 105, 110, 103, 77, 24, 49, 46, 49, 32, 118, 97, 114, 110, 105, 115, 104, 44, 32, 49, 46, 49, 32, 110, 103, 104, 116, 116, 112, 120, 64, 9, 120, 45, 118, 97, 114, 110, 105, 115, 104, 9, 50, 49, 53, 55, 57, 49, 56, 56, 54}
	client := NewResponseContext()
	client.Decode(buf)
	for name, values := range client.EmittedSet.Header {
		t.Log(name, len(values))
	}
}
