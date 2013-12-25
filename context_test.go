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
	client := NewContext()
	client.Decode(buf)

	expected := http.Header{
		"Method":    []string{"GET"},
		"Scheme":    []string{"http"},
		"Path":      []string{"/"},
		"Authority": []string{"www.example.com"},
	}

	if client.HT.Size() != 180 {
		t.Errorf("got %v\nwant %v", client.HT.Size(), 180)
	}

	if !reflect.DeepEqual(client.ES.Header, expected) {
		t.Errorf("got %v\nwant %v", client.ES.Header, expected)
	}
}
