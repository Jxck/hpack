package hpack

import (
	integer "github.com/jxck/hpack/integer_representation"
	"github.com/jxck/swrap"
)

func (frame *IndexedHeader) Encode() (buf *swrap.SWrap) {
	index := integer.Encode(frame.Index, 7)
	index[0] += 0x80
	return &index
}

func (frame *IndexedLiteral) Encode() (buf *swrap.SWrap) {
	// TODO: support huff encode
	index := integer.Encode(frame.Index, 6)
	if !frame.Indexing {
		index[0] += 0x40
	}
	buf = &index
	buf.Merge(integer.Encode(frame.ValueLength, 8))
	buf.Merge([]byte(frame.ValueString))
	return buf
}

func (frame *StringLiteral) Encode() (buf *swrap.SWrap) {
	// TODO: support huff encode
	sw := swrap.SWrap{}
	if frame.Indexing {
		sw.Add(0) // 0000 0000
	} else {
		sw.Add(0x40) // 0100 0000
	}
	buf = &sw
	buf.Merge(integer.Encode(frame.NameLength, 8))
	buf.Merge([]byte(frame.NameString))
	buf.Merge(integer.Encode(frame.ValueLength, 8))
	buf.Merge([]byte(frame.ValueString))
	return buf
}
