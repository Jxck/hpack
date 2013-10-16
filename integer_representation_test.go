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
	testcases := []struct {
		expected, actual uint64
	}{
		{10, DecodeInteger([]uint8{10}, 5)},
		{40, DecodeInteger([]uint8{31, 9}, 5)},
		{1337, DecodeInteger([]uint8{31, 154, 10}, 5)},
		{3000000, DecodeInteger([]uint8{31, 161, 141, 183, 1}, 5)},
	}

	for _, testcase := range testcases {
		actual := testcase.actual
		expected := testcase.expected
		if expected != actual {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
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
