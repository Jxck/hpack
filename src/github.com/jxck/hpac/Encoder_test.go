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
