package hpack

import (
	"github.com/jxck/hpack/huffman"
	integer "github.com/jxck/hpack/integer_representation"
	"github.com/jxck/swrap"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// Decode Wire byte seq to Slice of Frames
// TODO: make it return channel
func Decode(wire []byte, cxt CXT) (frames []Frame) {
	sw := swrap.New(wire)
	buf := &sw
	for buf.Len() > 0 {
		frames = append(frames, DecodeHeader(buf, cxt))
	}
	return frames
}

// Decode single Frame from buffer and return it
func DecodeHeader(buf *swrap.SWrap, cxt CXT) Frame {
	// check first byte
	types := (*buf)[0]
	// log.Println("types =", types)
	if types >= 0x80 { // 1xxx xxxx
		// Indexed Header Representation

		index := DecodePrefixedInteger(buf, 7)
		// log.Println("Indexed =", index)
		frame := NewIndexedHeader(index)
		return frame
	}
	if types == 0 { // 0000 0000
		// StringLiteral (indexing = true)

		// remove first byte defines type
		buf.Shift()

		indexing := true
		name := DecodeLiteral(buf, cxt)
		// log.Println("StringLiteral name =", name)
		value := DecodeLiteral(buf, cxt)
		// log.Println("StringLiteral value =", value)
		frame := NewStringLiteral(indexing, name, value)
		return frame
	}
	if types == 0x40 { // 0100 0000
		// StringLiteral (indexing = false)

		// remove first byte defines type
		buf.Shift()

		indexing := false
		name := DecodeLiteral(buf, cxt)
		// log.Println("StringLiteral name =", name)
		value := DecodeLiteral(buf, cxt)
		// log.Println("StringLiteral value =", value)
		frame := NewStringLiteral(indexing, name, value)
		return frame
	}
	if types&0xc0 == 0x40 { // 01xx xxxx & 1100 0000 == 0100 0000
		// IndexedLiteral (indexing = false)

		indexing := false
		index := DecodePrefixedInteger(buf, 6)
		// log.Println("IndexedLiteral index =", index)
		value := DecodeLiteral(buf, cxt)
		// log.Println("IndexedLiteral value =", value)
		frame := NewIndexedLiteral(indexing, index, value)
		return frame
	}
	if types&0xc0 == 0 { // 00xx xxxx & 1100 0000 == 0000 0000
		// IndexedLiteral (indexing = true)

		indexing := true
		index := DecodePrefixedInteger(buf, 6)
		// log.Println("IndexedLiteral index =", index)
		value := DecodeLiteral(buf, cxt)
		// log.Println("IndexedLiteral value =", value)
		frame := NewIndexedLiteral(indexing, index, value)
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

func DecodeLiteral(buf *swrap.SWrap, cxt CXT) (value string) {
	// log.Println(buf)

	// 最初のバイトを取り出す
	first := (*buf)[0]

	// 最初の 1bit をみて huffman かどうか取得
	huffmanEncoded := (first&0x80 == 0x80)

	// log.Println("buf.Len()", buf.Len())

	if huffmanEncoded {
		// 最初のバイトから 1 bit 目を消す
		(*buf)[0] = first & 127

		// ここで prefixed Integer 7 で読む。
		b := DecodePrefixedInteger(buf, 7)

		// その長さの分だけバイト値を取り出す
		code := []byte{}
		for ; b > 0; b-- {
			code = append(code, buf.Shift())
		}

		// log.Println("code", code)

		// コンテキストに合わせてデコード
		if cxt == REQUEST {
			value = string(huffman.DecodeRequest(code))
		} else if cxt == RESPONSE {
			value = string(huffman.DecodeResponse(code))
		}
		// log.Println(cxt, value)
	} else {
		valueLength := DecodePrefixedInteger(buf, 8)
		value = DecodeString(buf, valueLength)
	}
	return value
}
