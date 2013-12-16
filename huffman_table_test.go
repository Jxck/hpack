package hpack

import (
	"fmt"
	"testing"
)

func toHexString(hex []byte) (hexstr string) {
	for _, v := range hex {
		hexstr += fmt.Sprintf("%x", v)
	}
	return hexstr
}

var testcase = []struct {
	str, hex string
}{
	{"www.example.com", "db6d883e68d1cb1225ba7f"},
	{"no-cache", "63654a1398ff"},
	{"custom-key", "4eb08b749790fa7f"},
	{"custom-value", "4eb08b74979a17a8ff"},
}

func TestHuffmanEncode(t *testing.T) {
	for _, tc := range testcase {
		raw := []byte(tc.str)
		expected := tc.hex
		encoded := HuffmanEncode(raw)
		actual := toHexString(encoded)
		if actual != expected {
			t.Errorf("\ngot  %v\nwant %v", actual, expected)
		}
	}
}
