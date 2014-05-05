package hpack

import (
	"github.com/jxck/hpack/huffman"
	integer "github.com/jxck/hpack/integer_representation"
	. "github.com/jxck/logger"
	"github.com/jxck/swrap"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// Decode Wire byte seq to Slice of Frames
func Decode(wire []byte) (frames []Frame) {
	buf := swrap.Make(wire)
	for buf.Len() > 0 {
		frames = append(frames, DecodeHeader(buf))
	}
	return frames
}

// Decode single Frame from buffer and return it
func DecodeHeader(buf *swrap.SWrap) Frame {
	// check first byte
	types := (*buf)[0]
	Trace("types = %v", types)
	if types >= 0x80 { // 1xxx xxxx
		// Indexed Header Representation

		index := DecodePrefixedInteger(buf, 7)
		Trace("Indexed = %v", index)
		frame := NewIndexedHeader(index)

		if index == 0 {
			// TODO: Decoding Error
			log.Fatal("Decoding Error: The index value of 0 is not used.")
		}
		return frame
	}
	if types == 0 { // 0000 0000
		// StringLiteral (indexing = true)

		// remove first byte defines type
		buf.Shift()

		indexing := true
		name := DecodeLiteral(buf)
		Trace("StringLiteral name = %v", name)
		value := DecodeLiteral(buf)
		Trace("StringLiteral value = %v", value)
		frame := NewStringLiteral(indexing, name, value)
		return frame
	}
	if types == 0x40 { // 0100 0000
		// StringLiteral (indexing = false)

		// remove first byte defines type
		buf.Shift()

		indexing := false
		name := DecodeLiteral(buf)
		Trace("StringLiteral name = %v", name)
		value := DecodeLiteral(buf)
		Trace("StringLiteral value = %v", value)
		frame := NewStringLiteral(indexing, name, value)
		return frame
	}
	if types&0xc0 == 0x40 { // 01xx xxxx & 1100 0000 == 0100 0000
		// IndexedLiteral (indexing = false)

		indexing := false
		index := DecodePrefixedInteger(buf, 6)
		Trace("IndexedLiteral index = %v", index)
		value := DecodeLiteral(buf)
		Trace("IndexedLiteral value = %v", value)
		frame := NewIndexedLiteral(indexing, index, value)
		return frame
	}
	if types&0xc0 == 0 { // 00xx xxxx & 1100 0000 == 0000 0000
		// IndexedLiteral (indexing = true)

		indexing := true
		index := DecodePrefixedInteger(buf, 6)
		Trace("IndexedLiteral index = %v", index)
		value := DecodeLiteral(buf)
		Trace("IndexedLiteral value = %v", value)
		frame := NewIndexedLiteral(indexing, index, value)
		return frame
	}
	if types == 0x30 { // 0011 0000
		// remove first byte defines type
		buf.Shift()
		frame := NewEmptyReferenceSet()
		return frame
	}
	if types&0xf0 == 0x20 { // 0010 xxxx & 1111 0000 == 0010 0000
		maxSize := DecodePrefixedInteger(buf, 4)
		frame := NewChangeHeaderTableSize(maxSize)
		return frame
	}
	return nil
}

// read N prefixed Integer from buffer as uint64
func DecodePrefixedInteger(buf *swrap.SWrap, N uint8) uint64 {
	tmp := integer.ReadPrefixedInteger(buf, N)
	return integer.Decode(tmp, N)
}

// read n byte from buffer as string
func DecodeString(buf *swrap.SWrap, n uint64) string {
	valueBytes := make([]byte, 0, n)
	for i := n; i > 0; i-- {
		valueBytes = append(valueBytes, buf.Shift())
	}
	return string(valueBytes)
}

func DecodeLiteral(buf *swrap.SWrap) (value string) {
	// 最初のバイトを取り出す
	first := (*buf)[0]

	// 最初の 1bit をみて huffman かどうか取得
	huffmanEncoded := (first&0x80 == 0x80)

	Trace("huffman = %t", huffmanEncoded)
	if huffmanEncoded {
		// 最初のバイトから 1 bit 目を消す
		(*buf)[0] = first & 127

		// ここで prefixed Integer 7 で読む。
		b := DecodePrefixedInteger(buf, 7)
		Trace("Literal Length = %v", b)

		// その長さの分だけバイト値を取り出す
		code := make([]byte, 0)
		for ; b > 0; b-- {
			code = append(code, buf.Shift())
		}

		// ハフマンデコード
		value = string(huffman.Decode(code))
		Trace("decoded = %v", value)
	} else {
		valueLength := DecodePrefixedInteger(buf, 7)
		value = DecodeString(buf, valueLength)
	}
	return value
}
