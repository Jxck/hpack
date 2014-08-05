package huffman

import (
	"fmt"
	assert "github.com/Jxck/assertion"
	"reflect"
	"testing"
	"testing/quick"
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

var testCase = []struct {
	str, hex string
}{
	{"www.example.com", "db6d883e68d1cb1225ba7f"},
	{"no-cache", "63654a1398ff"},
	{"custom-key", "4eb08b749790fa7f"},
	{"custom-value", "4eb08b74979a17a8ff"},
	{"302", "98a7"},
	{"gzip", "cbd54e"},
	{"private", "73d5cd111f"},
	{
		"Mon, 21 Oct 2013 20:13:21 GMT",
		"ef6b3a7a0e6e8fa263d0729a6e8397d869bd873747bbbfc7",
	},
	{
		"Mon, 21 Oct 2013 20:13:22 GMT",
		"ef6b3a7a0e6e8fa263d0729a6e8397d869bd873f47bbbfc7",
	},
	{
		"https://www.example.com",
		"ce31743d801b6db107cd1a396244b74f",
	},
	{
		"foo=ASDJKHQKBZXOQWEOPIUAXQWEOIU; max-age=3600; version=1",
		"c5adb77f876fc7fbf7fdbfbebff3f7f4fb7ebbbe9f5f87e37fefedfaeefa7c3f1d5d1a23ce546436cd494bd5d1cc5f0535969b",
	},
}

func TestEncode(t *testing.T) {
	for _, tc := range testCase {
		raw := []byte(tc.str)
		expected := tc.hex
		encoded := Encode(raw)
		actual := toHexString(encoded)
		assert.Equal(t, actual, expected)
	}
}

func TestDecode(t *testing.T) {
	expected := "302"
	// Show(root)
	var code = []byte{0x98, 0xa7}
	result := Decode(code)
	actual := string(result)
	assert.Equal(t, actual, expected)
}

func TestEncodeDecode(t *testing.T) {
	for _, tc := range testCase {
		expected := []byte(tc.str)
		encoded := Encode(expected)
		actual := Decode(encoded)
		assert.Equal(t, actual, expected)
	}
}

func TestQuickCheckEncodeDecode(t *testing.T) {
	f := func(expected []byte) bool {
		var encoded, actual []byte
		encoded = Encode(expected)
		actual = Decode(encoded)
		return reflect.DeepEqual(actual, expected)
	}

	c := &quick.Config{
		MaxCountScale: 100,
	}

	if err := quick.Check(f, c); err != nil {
		t.Error(err)
	}
}
