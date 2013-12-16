package hpack

import (
	"fmt"
	"testing"
)

func TestHuffmanEncode(t *testing.T) {
	expected := "db6d 883e 68d1 cb12 25ba 7f"
	raw := []byte("www.example.com")
	encoded := HuffmanEncode(raw)

	actual := ""
	for i, v := range encoded {
		actual += fmt.Sprintf("%x", v)
		if i%2 == 1 {
			actual += fmt.Sprint(" ")
		}
	}
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
