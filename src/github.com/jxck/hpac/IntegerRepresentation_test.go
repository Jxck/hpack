package hpac

import (
	"testing"
)

func TestEncodeInteger(t *testing.T) {
	buf := EncodeInteger(10, 5)
	expect := []uint8{10}
	for i, j := range buf.Bytes() {
		if expect[i] != j {
			t.Errorf("got %v\nwant %v", j, expect[i])
		}
	}

	buf = EncodeInteger(1337, 5)
	expect = []uint8{31, 154, 10}
	for i, j := range buf.Bytes() {
		if expect[i] != j {
			t.Errorf("got %v\nwant %v", j, expect[i])
		}
	}
}

func TestDecodeInteger(t *testing.T) {
	buf := EncodeInteger(10, 5)
	n := DecodeInteger(buf.Bytes(), 5)
	if n != 10 {
		t.Errorf("got %v\nwant %v", n, 10)
	}

	buf = EncodeInteger(40, 5)
	t.Log(buf.Bytes())
	n = DecodeInteger(buf.Bytes(), 5)
	if n != 40 {
		t.Errorf("got %v\nwant %v", n, 40)
	}

	buf = EncodeInteger(1337, 5)
	n = DecodeInteger(buf.Bytes(), 5)
	if n != 1337 {
		t.Errorf("got %v\nwant %v", n, 1337)
	}

	buf = EncodeInteger(3000000, 5)
	n = DecodeInteger(buf.Bytes(), 5)
	if n != 3000000 {
		t.Errorf("got %v\nwant %v", n, 3000000)
	}
}
