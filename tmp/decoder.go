package hpack

import (
	"bytes"
	"encoding/binary"
	integer "github.com/jxck/hpack/integer_representation"
	"log"
)

// Decode Wire byte seq to Slice of Frames
// TODO: make it return channel
func Decode(wire []byte) []Frame {
	buf := bytes.NewBuffer(wire)
	frames := []Frame{} // TODO: make()
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
	if types >= 0x80 { // >= 128
		// Indexed Header Representation

		// unread first byte for parse frame
		buf.UnreadByte()

		frame := new(IndexedHeader)
		frame.Index = DecodePrefixedInteger(buf, 7)
		return frame
	}
	if types == 0x40 {
		// Literal Header with Incremental Indexing - New Name

		frame := new(NewNameWithIncrementalIndexing)
		frame.NameLength = DecodePrefixedInteger(buf, 8)
		frame.NameString = DecodeString(buf, frame.NameLength)
		frame.ValueLength = DecodePrefixedInteger(buf, 8)
		frame.ValueString = DecodeString(buf, frame.ValueLength)
		return frame
	}
	if types == 0x60 {
		// Literal Header without Indexing - New Name

		frame := new(NewNameWithoutIndexing)
		frame.NameLength = DecodePrefixedInteger(buf, 8)
		frame.NameString = DecodeString(buf, frame.NameLength)
		frame.ValueLength = DecodePrefixedInteger(buf, 8)
		frame.ValueString = DecodeString(buf, frame.ValueLength)
		return frame
	}
	if types>>5 == 0x2 {
		// iteral Header with Incremental Indexing - Indexed Name

		// unread first byte for parse frame
		buf.UnreadByte()

		frame := new(IndexedNameWithIncrementalIndexing)
		// 0 describes "not in the header table", but index of Header Table start with 0
		// so Index is represented as +1 integer
		frame.Index = DecodePrefixedInteger(buf, 5) - 1
		frame.ValueLength = DecodePrefixedInteger(buf, 8)
		frame.ValueString = DecodeString(buf, frame.ValueLength)
		return frame
	}
	if types&0x60 == 0x60 {
		// Literal Header without Indexing - Indexed Name

		// unread first byte for parse frame
		buf.UnreadByte()

		frame := new(IndexedNameWithoutIndexing)
		frame.Index = DecodePrefixedInteger(buf, 5) - 1
		frame.ValueLength = DecodePrefixedInteger(buf, 8)
		frame.ValueString = DecodeString(buf, frame.ValueLength)
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
