package hpac

import (
	"bytes"
	"testing"
)

func TestEncodeInteger(t *testing.T) {
	var I int
	var actual, expected []uint8

	I = 10
	expected = []uint8{10}
	actual = EncodeInteger(I, 5).Bytes()
	if !bytes.Equal(expected, actual) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	I = 40
	expected = []uint8{31, 9}
	actual = EncodeInteger(I, 5).Bytes()
	if !bytes.Equal(expected, actual) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	I = 1337
	expected = []uint8{31, 154, 10}
	actual = EncodeInteger(I, 5).Bytes()
	if !bytes.Equal(expected, actual) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	I = 3000000
	expected = []uint8{31, 161, 141, 183, 1}
	actual = EncodeInteger(I, 5).Bytes()
	if !bytes.Equal(expected, actual) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestDecodeInteger(t *testing.T) {
	var actual, expected uint32
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
