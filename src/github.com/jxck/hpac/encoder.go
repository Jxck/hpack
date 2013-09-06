package hpac

import (
	"bytes"
	"log"
)

func EncodeHeader(frame Frame) *bytes.Buffer {
	switch frame.(type) {
	case *IndexedHeader:
		f := frame.(*IndexedHeader)
		return f.Encode()
	case *NewNameWithoutIndexing:
		f := frame.(*NewNameWithoutIndexing)
		return f.Encode()
	case *IndexedNameWithoutIndexing:
		f := frame.(*IndexedNameWithoutIndexing)
		return f.Encode()
	case *IndexedNameWithIncrementalIndexing:
		f := frame.(*IndexedNameWithIncrementalIndexing)
		return f.Encode()
	case *NewNameWithIncrementalIndexing:
		f := frame.(*NewNameWithIncrementalIndexing)
		return f.Encode()
	case *IndexedNameWithSubstitutionIndexing:
		f := frame.(*IndexedNameWithSubstitutionIndexing)
		return f.Encode()
	case *NewNameWithSubstitutionIndexing:
		f := frame.(*NewNameWithSubstitutionIndexing)
		return f.Encode()
	default:
		log.Println("unmatch")
		return nil
	}
}

func (frame *IndexedHeader) Encode() *bytes.Buffer {
	index := EncodeInteger(frame.Index, 7).Bytes()
	buf := bytes.NewBuffer([]byte{0x80 + index[0]})
	index = index[1:]
	if len(index) > 0 {
		buf.Write(index)
	}
	return buf
}

func (frame *NewNameWithoutIndexing) Encode() *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{0x60})
	buf.Write(EncodeInteger(frame.NameLength, 8).Bytes())
	buf.WriteString(frame.NameString)
	buf.Write(EncodeInteger(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}

func (frame *IndexedNameWithoutIndexing) Encode() *bytes.Buffer {
	index := EncodeInteger(frame.Index+1, 5).Bytes()
	buf := bytes.NewBuffer([]byte{0x60 + index[0]})
	index = index[1:]
	if len(index) > 0 {
		buf.Write(index)
	}
	buf.Write(EncodeInteger(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}

func (frame *IndexedNameWithIncrementalIndexing) Encode() *bytes.Buffer {
	index := EncodeInteger(frame.Index+1, 5).Bytes()
	buf := bytes.NewBuffer([]byte{0x40 + index[0]})
	index = index[1:]
	if len(index) > 0 {
		buf.Write(index)
	}
	buf.Write(EncodeInteger(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}

func (frame *NewNameWithIncrementalIndexing) Encode() *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{0x40})
	buf.Write(EncodeInteger(frame.NameLength, 8).Bytes())
	buf.WriteString(frame.NameString)
	buf.Write(EncodeInteger(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}

func (frame *IndexedNameWithSubstitutionIndexing) Encode() *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{})
	buf.Write(EncodeInteger(frame.Index+1, 6).Bytes())
	buf.Write(EncodeInteger(frame.SubstitutedIndex, 8).Bytes())
	buf.Write(EncodeInteger(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}

func (frame *NewNameWithSubstitutionIndexing) Encode() *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{0})
	buf.Write(EncodeInteger(frame.NameLength, 8).Bytes())
	buf.WriteString(frame.NameString)
	buf.Write(EncodeInteger(frame.SubstitutedIndex, 8).Bytes())
	buf.Write(EncodeInteger(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}
