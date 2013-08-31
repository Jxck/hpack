package hpac

import (
	"bytes"
	"log"
)

type Encoder struct {
	requestHeaderTable  HeaderTable
	responseHeaderTable HeaderTable
	referenceSet        ReferenceSet
}

func NewEncoder() Encoder {
	var encoder = Encoder{
		requestHeaderTable:  RequestHeaderTable,
		responseHeaderTable: ResponseHeaderTable,
		referenceSet:        ReferenceSet{},
	}
	return encoder
}

func EncodeHeader(frame Frame) *bytes.Buffer {
	switch frame.(type) {
	case *IndexedHeader:
		f := frame.(*IndexedHeader)
		return encodeIndexedHeader(f)
	case *NewNameWithoutIndexing:
		f := frame.(*NewNameWithoutIndexing)
		return encodeNewNameWithoutIndexing(f)
	case *IndexedNameWithoutIndexing:
		f := frame.(*IndexedNameWithoutIndexing)
		return encodeIndexedNameWithoutIndexing(f)
	case *IndexedNameWithIncrementalIndexing:
		f := frame.(*IndexedNameWithIncrementalIndexing)
		return encodeIndexedNameWithIncrementalIndexing(f)
	case *NewNameWithIncrementalIndexing:
		f := frame.(*NewNameWithIncrementalIndexing)
		return encodeNewNameWithIncrementalIndexing(f)
	case *IndexedNameWithSubstitutionIndexing:
		f := frame.(*IndexedNameWithSubstitutionIndexing)
		return encodeIndexedNameWithSubstitutionIndexing(f)
	case *NewNameWithSubstitutionIndexing:
		f := frame.(*NewNameWithSubstitutionIndexing)
		return encodeNewNameWithSubstitutionIndexing(f)
	default:
		log.Println("unmatch")
		return nil
	}
}

func encodeIndexedHeader(frame *IndexedHeader) *bytes.Buffer {
	index := frame.Index
	buf := bytes.NewBuffer([]byte{index + 0x80})
	return buf
}

func encodeNewNameWithoutIndexing(frame *NewNameWithoutIndexing) *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{0x60})
	buf.Write(EncodeInteger(frame.NameLength, 8).Bytes())
	buf.WriteString(frame.NameString)
	buf.Write(EncodeInteger(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}

func encodeIndexedNameWithoutIndexing(frame *IndexedNameWithoutIndexing) *bytes.Buffer {
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

func encodeIndexedNameWithIncrementalIndexing(frame *IndexedNameWithIncrementalIndexing) *bytes.Buffer {
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

func encodeNewNameWithIncrementalIndexing(frame *NewNameWithIncrementalIndexing) *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{0x40})
	buf.Write(EncodeInteger(frame.NameLength, 8).Bytes())
	buf.WriteString(frame.NameString)
	buf.Write(EncodeInteger(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}

func encodeIndexedNameWithSubstitutionIndexing(frame *IndexedNameWithSubstitutionIndexing) *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{})
	buf.Write(EncodeInteger(frame.Index+1, 6).Bytes())
	buf.Write(EncodeInteger(frame.SubstitutedIndex, 8).Bytes())
	buf.Write(EncodeInteger(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}

func encodeNewNameWithSubstitutionIndexing(frame *NewNameWithSubstitutionIndexing) *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{0})
	buf.Write(EncodeInteger(frame.NameLength, 8).Bytes())
	buf.WriteString(frame.NameString)
	buf.Write(EncodeInteger(frame.SubstitutedIndex, 8).Bytes())
	buf.Write(EncodeInteger(frame.ValueLength, 8).Bytes())
	buf.WriteString(frame.ValueString)
	return buf
}
