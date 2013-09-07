package hpac

import (
	"net/http"
	"testing"
)

// TODO: check refset, emmitedset in test
func TestIncrementalIndexingWithIndexedName(t *testing.T) {
	frame := CreateIndexedNameWithIncrementalIndexing(0, "ftp")

	server := NewContext()
	server.Decode(frame.Encode().Bytes())
	last := len(server.RequestHeaderTable) - 1
	if server.RequestHeaderTable[last].Value != "ftp" {
		t.Errorf("got %v\nwant %v", server.RequestHeaderTable[last].Name, "ftp")
	}
}

func TestIncrementalIndexingWithNewName(t *testing.T) {
	frame := CreateNewNameWithIncrementalIndexing("x-hello", "world")

	server := NewContext()
	server.Decode(frame.Encode().Bytes())
	last := server.RequestHeaderTable[len(server.RequestHeaderTable)-1]
	if last.Name != "x-hello" || last.Value != "world" {
		t.Errorf("got (%v, %v)\nwant (%v %v)", last.Name, last.Value, "x-hello", "world")
	}
}

func TestSubstitutionIndexingWithIndexedName(t *testing.T) {
	frame := CreateIndexedNameWithSubstitutionIndexing(0, 10, "ftp")

	server := NewContext()
	server.Decode(frame.Encode().Bytes())
	target := server.RequestHeaderTable[10]
	if target.Name != ":scheme" || target.Value != "ftp" {
		t.Errorf("got (%v, %v)\nwant (%v %v)", target.Name, target.Value, "x-hello", "world")
	}
}

func TestSubstitutionIndexingWithNewName(t *testing.T) {
	frame := CreateNewNameWithSubstitutionIndexing("x-hello", 10, "world")

	server := NewContext()
	server.Decode(frame.Encode().Bytes())
	target := server.RequestHeaderTable[10]
	if target.Name != "x-hello" || target.Value != "world" {
		t.Errorf("got (%v, %v)\nwant (%v %v)", target.Name, target.Value, "x-hello", "world")
	}
}

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
