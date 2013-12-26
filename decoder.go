package hpack

import (
	"bytes"
	"encoding/binary"
	integer "github.com/jxck/hpack/integer_representation"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// Decode Wire byte seq to Slice of Frames
// TODO: make it return channel
func Decode(wire []byte) (frames []Frame) {
	buf := bytes.NewBuffer(wire)
	for buf.Len() > 0 {
		frames = append(frames, DecodeHeader(buf))
	}
	return frames
}

// Decode single Frame from buffer and return it
func DecodeHeader(buf *bytes.Buffer) Frame {
	var types uint8
	if err := binary.Read(buf, binary.BigEndian, &types); err != nil {
		log.Println("binary.Read failed:", err)
	}
	if types >= 0x80 { // 1xxx xxxx
		// Indexed Header Representation

		// unread first byte for parse frame
		buf.UnreadByte()

		index := DecodePrefixedInteger(buf, 7)
		frame := NewIndexedHeader(index)
		return frame
	}
	if types == 0 { // 0000 0000
		// Literal Header Field with Incremental Indexing - New Name (F=0)
		indexing := true
		nameLength := DecodePrefixedInteger(buf, 8)
		name := DecodeString(buf, nameLength)
		valueLength := DecodePrefixedInteger(buf, 8)
		value := DecodeString(buf, valueLength)
		frame := NewStringLiteral(indexing, name, value)
		return frame
	}
	if types == 0x40 { // 0100 0000
		// Literal Header Field without Indexing - New Name (F=1)

		indexing := false
		nameLength := DecodePrefixedInteger(buf, 8)
		name := DecodeString(buf, nameLength)
		valueLength := DecodePrefixedInteger(buf, 8)
		value := DecodeString(buf, valueLength)
		frame := NewStringLiteral(indexing, name, value)
		return frame
	}
	if types&0xc0 == 0x40 { // 01xx xxxx & 1100 0000 == 0100 0000
		// Literal Header Field without Indexing - Indexed Name (F=1)

		// unread first byte for parse frame
		buf.UnreadByte()

		indexing := false
		index := DecodePrefixedInteger(buf, 6)
		valueLength := DecodePrefixedInteger(buf, 8)
		value := DecodeString(buf, valueLength)
		frame := NewIndexedLiteral(indexing, index, value)
		return frame
	}
	if types&0xc0 == 0 { // 00xx xxxx & 1100 0000 == 0000 0000
		// Literal Header Field with Incremental Indexing - Indexed Name (F=0)
		// unread first byte for parse frame
		buf.UnreadByte()

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
func DecodePrefixedInteger(buf *bytes.Buffer, N uint8) uint64 {
	tmp := integer.ReadPrefixedInteger(buf, N).Bytes()
	return integer.Decode(tmp, N)
}

// read n byte from buffer as string
func DecodeString(buf *bytes.Buffer, n uint64) string {
	valueBytes := make([]byte, n)
	binary.Read(buf, binary.BigEndian, &valueBytes) // err
	return string(valueBytes)
}
