package hpack

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRequestWithoutHuffman(t *testing.T) {
	buf := []byte{
		0x82, 0x87, 0x86, 0x04,
		0x0f, 0x77, 0x77, 0x77,
		0x2e, 0x65, 0x78, 0x61,
		0x6d, 0x70, 0x6c, 0x65,
		0x2e, 0x63, 0x6f, 0x6d,
	}

	expectedHeader := http.Header{
		"Method":    []string{"GET"},
		"Scheme":    []string{"http"},
		"Path":      []string{"/"},
		"Authority": []string{"www.example.com"},
	}

	expectedHeaderFields := []HeaderField{
		HeaderField{":authority", "www.example.com"},
		HeaderField{":path", "/"},
		HeaderField{":scheme", "http"},
		HeaderField{":method", "GET"},
	}

	client := NewContext()
	client.Decode(buf)

	// test Header Table
	if client.HT.Size() != 180 {
		t.Errorf("got %v\nwant %v", client.HT.Size(), 180)
	}

	// test Header Table
	for i, hf := range expectedHeaderFields {
		t.Log(*client.HT.HeaderFields[i] == hf)
	}

	// test Emitted Set
	if !reflect.DeepEqual(client.ES.Header, expectedHeader) {
		t.Errorf("got %v\nwant %v", client.ES.Header, expectedHeader)
	}

	// test Reference Set
	for i, hf := range *client.RS {
		t.Log(i, hf)
	}
}
