package hpack

import (
	assert "github.com/Jxck/assertion"
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

	var actual uint64 = 200
	assert.Equal(t, ht.Size(), actual)
	assert.Equal(t, ht.Len(), 5)
}
