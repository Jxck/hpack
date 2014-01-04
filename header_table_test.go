package hpack

import (
	"testing"
)

func TestHeaderTableSizeLen(t *testing.T) {
	ht := HeaderTable{
		DEFAULT_HEADER_TABLE_SIZE,
		[]*HeaderField{
			NewHeaderField("1234", "1234"),
			NewHeaderField("1234", "1234"),
			NewHeaderField("1234", "1234"),
			NewHeaderField("1234", "1234"),
			NewHeaderField("1234", "1234"),
		},
	}

	size := ht.Size()
	length := ht.Len()

	if size != 200 {
		t.Errorf("got %v\nwant %v", size, 200)
	}

	if length != 5 {
		t.Errorf("got %v\nwant %v", length, 5)
	}
}
