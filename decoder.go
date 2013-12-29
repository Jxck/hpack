package hpack

import (
	integer "github.com/jxck/hpack/integer_representation"
	"github.com/jxck/swrap"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// Decode Wire byte seq to Slice of Frames
// TODO: make it return channel
func Decode(wire []byte) (frames []Frame) {
	sw := swrap.New(wire)
	buf := &sw
	for buf.Len() > 0 {
		frames = append(frames, DecodeHeader(buf))
	}
	return frames
}

// Decode single Frame from buffer and return it
func DecodeHeader(buf *swrap.SWrap) Frame {
	types := buf.Shift()
	if types >= 0x80 { // 1xxx xxxx
		// Indexed Header Representation

		// unread first byte for parse frame
		buf.UnShift(types)

		index := DecodePrefixedInteger(buf, 7)
		frame := NewIndexedHeader(index)
		return frame
	}
	if types == 0 { // 0000 0000
		// StringLiteral (indexing = true)

		indexing := true
		nameLength := DecodePrefixedInteger(buf, 8)
		name := DecodeString(buf, nameLength)
		valueLength := DecodePrefixedInteger(buf, 8)
		value := DecodeString(buf, valueLength)
		frame := NewStringLiteral(indexing, name, value)
		return frame
	}
	if types == 0x40 { // 0100 0000
		// StringLiteral (indexing = false)

		indexing := false
		nameLength := DecodePrefixedInteger(buf, 8)
		name := DecodeString(buf, nameLength)
		valueLength := DecodePrefixedInteger(buf, 8)
		value := DecodeString(buf, valueLength)
		frame := NewStringLiteral(indexing, name, value)
		return frame
	}
	if types&0xc0 == 0x40 { // 01xx xxxx & 1100 0000 == 0100 0000
		// IndexedLiteral (indexing = false)

		// unread first byte for parse frame
		buf.UnShift(types)

		indexing := false
		index := DecodePrefixedInteger(buf, 6)

		huff := DetectHuffman(buf)

		log.Println(huff)
		valueLength := DecodePrefixedInteger(buf, 7)
		value := DecodeString(buf, valueLength)
		frame := NewIndexedLiteral(indexing, index, value)
		return frame
	}
	if types&0xc0 == 0 { // 00xx xxxx & 1100 0000 == 0000 0000
		// IndexedLiteral (indexing = true)

		// unread first byte for parse frame
		buf.UnShift(types)

		indexing := true
		index := DecodePrefixedInteger(buf, 6)
		valueLength := DecodePrefixedInteger(buf, 8)
		value := DecodeString(buf, valueLength)
		frame := NewIndexedLiteral(indexing, index, value)
		return frame
	}
	return nil
}

// read N prefixed Integer from buffer as uint64
func DecodePrefixedInteger(buf *swrap.SWrap, N uint8) uint64 {
	tmp := integer.ReadPrefixedInteger(buf, N)
	log.Println(tmp)
	return integer.Decode(tmp, N)
}

// read n byte from buffer as string
func DecodeString(buf *swrap.SWrap, n uint64) string {
	log.Println(buf, n)
	valueBytes := make([]byte, 0, n)
	for i := n; i > 0; i-- {
		valueBytes = append(valueBytes, buf.Shift())
	}
	return string(valueBytes)
}

func DetectHuffman(buf *swrap.SWrap) bool {
	b := buf.Shift()
	huff := false
	if b&0x80 == 0x80 {
		huff = true
		b = b & 127
	}
	buf.UnShift(b)
	return huff
}
