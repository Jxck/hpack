package integer_representation

import (
	"bytes"
	"testing"
	"testing/quick"
)

func TestEncode(t *testing.T) {
	testcases := []struct {
		expected, actual []byte
	}{
		{[]byte{10}, Encode(10, 5)},
		{[]byte{31, 9}, Encode(40, 5)},
		{[]byte{31, 154, 10}, Encode(1337, 5)},
		{[]byte{31, 161, 141, 183, 1}, Encode(3000000, 5)},
	}

	for _, testcase := range testcases {
		actual := testcase.actual
		expected := testcase.expected
		if !bytes.Equal(expected, actual) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}

func TestDecode(t *testing.T) {
	testcases := []struct {
		expected, actual uint64
	}{
		{10, Decode([]byte{10}, 5)},
		{40, Decode([]byte{31, 9}, 5)},
		{1337, Decode([]byte{31, 154, 10}, 5)},
		{3000000, Decode([]byte{31, 161, 141, 183, 1}, 5)},
	}

	for _, testcase := range testcases {
		actual := testcase.actual
		expected := testcase.expected
		if expected != actual {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}

func TestEncodeDecodeQuickCheck(t *testing.T) {
	f := func(I uint64) bool {
		var N uint8 = 5
		buf := Encode(I, N)
		actual := Decode(buf, N)
		t.Log(I)
		t.Log(actual)
		return actual == I
	}
	c := &quick.Config{}

	if err := quick.Check(f, c); err != nil {
		t.Error(err)
	}
}

func TestReadPrefixedInteger(t *testing.T) {
	// 0x1F 0001 1111
	// 0x95 1001 0101
	// 0x0A 0000 1010
	// 0x06 0000 0110
	var prefix uint8 = 5
	buf := []byte{0x1F, 0x95, 0x0A, 0x06}
	expected := []byte{0x1F, 0x95, 0xA}
	actual := ReadPrefixedInteger(buf, prefix)
	if !bytes.Equal(expected, actual) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
