package hpack

import (
	"bytes"
	"testing"
)

func TestEncodeInteger(t *testing.T) {
	testcases := []struct {
		expected, actual []uint8
	}{
		{[]uint8{10}, EncodeInteger(10, 5).Bytes()},
		{[]uint8{31, 9}, EncodeInteger(40, 5).Bytes()},
		{[]uint8{31, 154, 10}, EncodeInteger(1337, 5).Bytes()},
		{[]uint8{31, 161, 141, 183, 1}, EncodeInteger(3000000, 5).Bytes()},
	}

	for _, testcase := range testcases {
		actual := testcase.actual
		expected := testcase.expected
		if !bytes.Equal(expected, actual) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}

func TestDecodeInteger(t *testing.T) {
	var actual, expected uint64
	var buf []uint8

	buf = []uint8{10}
	expected = 10
	actual = DecodeInteger(buf, 5)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	buf = []uint8{31, 9}
	expected = 40
	actual = DecodeInteger(buf, 5)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	buf = []uint8{31, 154, 10}
	expected = 1337
	actual = DecodeInteger(buf, 5)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	buf = []uint8{31, 161, 141, 183, 1}
	expected = 3000000
	actual = DecodeInteger(buf, 5)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestReadPrefixedInteger(t *testing.T) {
	// 0x1F 0001 1111
	// 0x0a 0000 1010
	// 0x06 0000 0110
	buf := bytes.NewBuffer([]byte{0x1f, 0x0a, 0x06})
	expected := []byte{0x1F, 0xA}
	actual := ReadPrefixedInteger(buf, 5).Bytes()
	if !bytes.Equal(expected, actual) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
