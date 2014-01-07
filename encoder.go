package hpack

import (
	"github.com/jxck/hpack/huffman"
	integer "github.com/jxck/hpack/integer_representation"
	"github.com/jxck/swrap"
)

func (frame *IndexedHeader) Encode() (buf *swrap.SWrap) {
	buf = new(swrap.SWrap)
	buf.Merge(integer.Encode(frame.Index, 7))
	(*buf)[0] += 0x80
	return buf
}

func (frame *IndexedLiteral) Encode() (buf *swrap.SWrap) {
	// TODO: support huff encode
	buf = new(swrap.SWrap)
	buf.Merge(integer.Encode(frame.Index, 6))
	if !frame.Indexing {
		(*buf)[0] += 0x40
	}
	// No Huffman
	buf.Merge(integer.Encode(frame.ValueLength, 7))
	buf.Merge([]byte(frame.ValueString))
	return buf
}

func (frame *StringLiteral) Encode() (buf *swrap.SWrap) {
	// TODO: support huff encode
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

// Encode() に request/response や flag などを渡すと、
// Frame Interface が変わるし、 IndexedHeader にはいらない。
// type に足すのも考えたけど、 New の引数を増やすか、使う側が設定する必要が
// あるのも微妙かと思い、別メソッドとすることにした。
func (frame *StringLiteral) EncodeHuffman(cxt CXT) (buf *swrap.SWrap) {
	buf = new(swrap.SWrap)
	if frame.Indexing {
		buf.Add(0) // 0000 0000
	} else {
		buf.Add(0x40) // 0100 0000
	}

	var encoded, length []byte

	// Name With Huffman
	if cxt == REQUEST {
		encoded = huffman.EncodeRequest([]byte(frame.NameString))
	} else {
		encoded = huffman.EncodeResponse([]byte(frame.NameString))
	}

	length = integer.Encode(uint64(len(encoded)), 7)
	length[0] += 0x80 // 1000 0000 (huffman flag)
	buf.Merge(length)
	buf.Merge(encoded)

	// Value With Huffman
	if cxt == REQUEST {
		encoded = huffman.EncodeRequest([]byte(frame.ValueString))
	} else {
		encoded = huffman.EncodeResponse([]byte(frame.ValueString))
	}
	length = integer.Encode(uint64(len(encoded)), 7)
	length[0] += 0x80 // 1000 0000 (huffman flag)
	buf.Merge(length)
	buf.Merge(encoded)

	return buf
}
