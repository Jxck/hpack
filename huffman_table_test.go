package hpack

import (
	"fmt"
	"reflect"
	"testing"
	"testing/quick"
)

// ===== Encode =====

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

// ===== Decode =====

func TestHuffmanDecodeResponse(t *testing.T) {
	expected := "302"
	// Show(root)
	var code = []byte{0x40, 0x9f}
	result := HuffmanDecodeResponse(code)
	actual := string(result)
	if actual != expected {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
	}
}

// ===== Encode -> Decode =====
func TestHuffmanEncodeDecode(t *testing.T) {
	// Request
	for _, tc := range requestTestCase {
		expected := []byte(tc.str)
		encoded := HuffmanEncodeRequest(expected)
		actual := HuffmanDecodeRequest(encoded)

		if reflect.DeepEqual(actual, expected) == false {
			t.Errorf("\ngot  %v\nwant %v", actual, expected)
		}
	}
	// Response
	for _, tc := range responseTestCase {
		expected := []byte(tc.str)
		encoded := HuffmanEncodeResponse(expected)
		actual := HuffmanDecodeResponse(encoded)

		if reflect.DeepEqual(actual, expected) == false {
			t.Errorf("\ngot  %v\nwant %v", actual, expected)
		}
	}
}

// ===== Quick Check =====
func TestQuickCheckHuffmanEncodeDecode(t *testing.T) {
	f := func(expected []byte) bool {
		var encoded, actual []byte
		// request
		encoded = HuffmanEncodeRequest(expected)
		actual = HuffmanDecodeRequest(encoded)
		req := reflect.DeepEqual(actual, expected)

		// response
		encoded = HuffmanEncodeResponse(expected)
		actual = HuffmanDecodeResponse(encoded)
		res := reflect.DeepEqual(actual, expected)

		return req && res
	}

	c := &quick.Config{
		MaxCountScale: 100,
	}

	if err := quick.Check(f, c); err != nil {
		t.Error(err)
	}
}
