package hpack

import (
	"bytes"
	integer "github.com/jxck/hpack/integer_representation"
)

func (frame *IndexedHeader) Encode() *bytes.Buffer {
	index := integer.Encode(frame.Index, 7).Bytes()
	buf := bytes.NewBuffer([]byte{0x80 + index[0]})
	index = index[1:]
	if len(index) > 0 {
		buf.Write(index)
	}
	return buf
}

func (frame *NewNameWithoutIndexing) Encode() *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{0x60})
	buf.Write(integer.Encode(frame.NameLength, 8).Bytes())
	buf.WriteString(frame.NameString)
	buf.Write(integer.Encode(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}

func (frame *IndexedNameWithoutIndexing) Encode() *bytes.Buffer {
	index := integer.Encode(frame.Index+1, 5).Bytes()
	buf := bytes.NewBuffer([]byte{0x60 + index[0]})
	index = index[1:]
	if len(index) > 0 {
		buf.Write(index)
	}
	buf.Write(integer.Encode(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}

func (frame *IndexedNameWithIncrementalIndexing) Encode() *bytes.Buffer {
	index := integer.Encode(frame.Index+1, 5).Bytes()
	buf := bytes.NewBuffer([]byte{0x40 + index[0]})
	index = index[1:]
	if len(index) > 0 {
		buf.Write(index)
	}
	buf.Write(integer.Encode(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}

func (frame *NewNameWithIncrementalIndexing) Encode() *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{0x40})
	buf.Write(integer.Encode(frame.NameLength, 8).Bytes())
	buf.WriteString(frame.NameString)
	buf.Write(integer.Encode(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}

func (frame *IndexedNameWithSubstitutionIndexing) Encode() *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{})
	buf.Write(integer.Encode(frame.Index+1, 6).Bytes())
	buf.Write(integer.Encode(frame.SubstitutedIndex, 8).Bytes())
	buf.Write(integer.Encode(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}

func (frame *NewNameWithSubstitutionIndexing) Encode() *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{0})
	buf.Write(integer.Encode(frame.NameLength, 8).Bytes())
	buf.WriteString(frame.NameString)
	buf.Write(integer.Encode(frame.SubstitutedIndex, 8).Bytes())
	buf.Write(integer.Encode(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}
