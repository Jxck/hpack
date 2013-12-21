package hpack

import (
	"bytes"
	integer "github.com/jxck/hpack/integer_representation"
)

func (frame *IndexedHeader) Encode() (buf *bytes.Buffer) {
	index := integer.Encode(frame.Index, 7).Bytes()
	buf = bytes.NewBuffer([]byte{128 + index[0]})
	index = index[1:]
	if len(index) > 0 {
		buf.Write(index)
	}
	return buf
}

func (frame *IndexedLiteral) Encode() (buf *bytes.Buffer) {
	index := integer.Encode(frame.Index, 6).Bytes()
	if frame.Indexing {
		buf = bytes.NewBuffer([]byte{index[0]}) // 00xx xxxx
	} else {
		buf = bytes.NewBuffer([]byte{64 + index[0]}) // 01xx xxxx
	}
	index = index[1:]
	if len(index) > 0 {
		buf.Write(index)
	}
	buf.Write(integer.Encode(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}

func (frame *StringLiteral) Encode() (buf *bytes.Buffer) {
	if frame.Indexing {
		buf = bytes.NewBuffer([]byte{0}) // 0000 0000
	} else {
		buf = bytes.NewBuffer([]byte{64}) // 0100 0000
	}
	buf.Write(integer.Encode(frame.NameLength, 8).Bytes())
	buf.WriteString(frame.NameString)
	buf.Write(integer.Encode(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}
