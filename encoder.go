package hpack

import (
	"github.com/jxck/hpack/huffman"
	integer "github.com/jxck/hpack/integer_representation"
	"github.com/jxck/swrap"
)

func (frame *IndexedHeader) Encode() (buf *swrap.SWrap) {
	buf = swrap.Make(integer.Encode(frame.Index, 7))
	(*buf)[0] += 0x80
	if frame.Index == 0 {
		// TODO: Encoding Error
	}
	return buf
}

func (frame *IndexedLiteral) Encode() (buf *swrap.SWrap) {
	buf = swrap.Make(integer.Encode(frame.Index, 6))
	if !frame.Indexing {
		(*buf)[0] += 0x40
	}
	// No Huffman
	buf.Merge(integer.Encode(frame.ValueLength, 7))
	buf.Merge([]byte(frame.ValueString))
	return buf
}

func (frame *IndexedLiteral) EncodeHuffman() (buf *swrap.SWrap) {
	buf = swrap.Make(integer.Encode(frame.Index, 6))
	if !frame.Indexing {
		(*buf)[0] += 0x40
	}

	var encoded, length []byte

	// Value With Huffman
	encoded = huffman.Encode([]byte(frame.ValueString))
	length = integer.Encode(uint64(len(encoded)), 7)
	length[0] += 0x80 // 1000 0000 (huffman flag)
	buf.Merge(length)
	buf.Merge(encoded)

	return buf
}

func (frame *StringLiteral) Encode() (buf *swrap.SWrap) {
	buf = new(swrap.SWrap)
	if frame.Indexing {
		buf.Add(0) // 0000 0000
	} else {
		buf.Add(0x40) // 0100 0000
	}
	// No Huffman
	buf.Merge(integer.Encode(frame.NameLength, 7))
	buf.Merge([]byte(frame.NameString))
	// No Huffman
	buf.Merge(integer.Encode(frame.ValueLength, 7))
	buf.Merge([]byte(frame.ValueString))
	return buf
}

func (frame *StringLiteral) EncodeHuffman() (buf *swrap.SWrap) {
	buf = new(swrap.SWrap)
	if frame.Indexing {
		buf.Add(0) // 0000 0000
	} else {
		buf.Add(0x40) // 0100 0000
	}

	var encoded, length []byte

	// Name With Huffman
	encoded = huffman.Encode([]byte(frame.NameString))

	length = integer.Encode(uint64(len(encoded)), 7)
	length[0] += 0x80 // 1000 0000 (huffman flag)
	buf.Merge(length)
	buf.Merge(encoded)

	// Value With Huffman
	encoded = huffman.Encode([]byte(frame.ValueString))
	length = integer.Encode(uint64(len(encoded)), 7)
	length[0] += 0x80 // 1000 0000 (huffman flag)
	buf.Merge(length)
	buf.Merge(encoded)

	return buf
}

func (frame *EmptyReferenceSet) Encode() (buf *swrap.SWrap) {
	buf = swrap.Make([]byte{0x30})
	return buf
}

func (frame *ChangeHeaderTableSize) Encode() (buf *swrap.SWrap) {
	buf = new(swrap.SWrap)
	buf.Add(0x20) // 0010 0000
	buf.Merge(integer.Encode(frame.MaxSize, 4))
	return buf
}
