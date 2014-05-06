package hpack

import (
	assert "github.com/jxck/assertion"
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

	assert.Equal(t, ht.Size(), 200)
	assert.Equal(t, ht.Len(), 5)
}
