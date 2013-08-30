package main

import (
	"bytes"
	. "github.com/jxck/hpac"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func EncodeHeader(frame Frame) *bytes.Buffer {
	switch frame.(type) {
	case *IndexedHeader:
		f := frame.(*IndexedHeader)
		return encodeIndexedHeader(f)
//	case *NewNameWithoutIndexing:
//		f := frame.(*NewNameWithoutIndexing)
//		return encodeNewNameWithoutIndexing(f)
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

//func encodeNewNameWithoutIndexing(frame *NewNameWithoutIndexing) *bytes.Buffer {
//	buf := bytes.NewBuffer([]byte{0x60})
//	buf.WriteByte(EncodeInteger(frame.NameLength, 8))
//	buf.WriteString(frame.NameString)
//	buf.WriteByte(EncodeInteger(frame.ValueLength, 8))
//	buf.WriteString(frame.ValueString)
//	return buf
//}

func TestIndexedHeader() {
	frame := NewIndexedHeader()
	frame.Index = 10
	buf := EncodeHeader(frame)
	log.Printf("%v", buf.Bytes())
}

//func NewNameWithoutIndexing() {
//	frame := NewNewNameWithoutIndexing()
//	frame.NameLength = 11
//	frame.NameString = "mynewheader"
//	frame.ValueLength = 5
//	frame.ValueString = "first"
//	buf := EncodeHeader(frame)
//	log.Printf("%v", buf.Bytes())
//}

func main() {
	var header = http.Header{
		":scheme":     []string{"http"},
		":path":       []string{"/index.html"},
		"mynewheader": []string{"first"},
	}
	_ = header

	TestIndexedHeader()
}
