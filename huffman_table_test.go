package hpack

import (
	"fmt"
	"testing"
)

func toHexString(hex []byte) (hexstr string) {
	for _, v := range hex {
		s := fmt.Sprintf("%x", v)
		if len(s) < 2 {
			s = "0" + s
		}
		hexstr += s
	}
	return hexstr
}

var requestTestCase = []struct {
	str, hex string
}{
	{"www.example.com", "db6d883e68d1cb1225ba7f"},
	{"no-cache", "63654a1398ff"},
	{"custom-key", "4eb08b749790fa7f"},
	{"custom-value", "4eb08b74979a17a8ff"},
}

func TestHuffmanEncodeRequest(t *testing.T) {
	for _, tc := range requestTestCase {
		raw := []byte(tc.str)
		expected := tc.hex
		encoded := HuffmanEncodeRequest(raw)
		actual := toHexString(encoded)
		if actual != expected {
			t.Errorf("\ngot  %v\nwant %v", actual, expected)
		}
	}
}

var responseTestCase = []struct {
	str, hex string
}{
	{"302", "409f"},
	{"gzip", "e1fbb30f"},
	{"private", "c31b39bf387f"},
	{
		"Mon, 21 Oct 2013 20:13:21 GMT",
		"a2fba20320f2ab303124018b490d3209e877",
	},
	{
		"Mon, 21 Oct 2013 20:13:22 GMT",
		"a2fba20320f2ab303124018b490d3309e877",
	},
	{
		"https://www.example.com",
		"e39e7864dd7afd3d3d248747db87284955f6ff",
	},
	{
		"foo=ASDJKHQKBZXOQWEOPIUAXQWEOIU; max-age=3600; version=1",
		"df7dfb36d3d9e1fcfc3fafe7abfcfefcbfaf3edf2f977fd36ff7fd79f6f977fd3de16bfa46fe10d889447de1ce18e565f76c2f",
	},
}

func TestHuffmanEncodeResponse(t *testing.T) {
	for _, tc := range responseTestCase {
		raw := []byte(tc.str)
		expected := tc.hex
		encoded := HuffmanEncodeResponse(raw)
		actual := toHexString(encoded)
		if actual != expected {
			t.Errorf("\ngot  %v\nwant %v", actual, expected)
		}
	}
}
