package main

import (
	//	"bytes"
	"github.com/jxck/hpac"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type Header struct {
	Name  string
	Value string
}

// ヘッダはポインタにしておく
type HeaderTable []*Header

//func (ht *HeaderTable) Add(header Header) {
//	*ht = append(*ht, header)
//}

// name と value が HeaderTable にあるかを探す // name, value とも一致 => index, *Header
// name はある          => index, nil
// ない                 => -1, nil
func (ht HeaderTable) SearchHeader(name, value string) (int, *Header) {
	// name が複数一致した時のために格納しておく
	// MEMO: スライスで持たず単一で最初だけもってもいいかもしれないが
	// もし無かった場合 0 になって、それが index=0 と紛らわしいので
	// slice でもって、長さで判断できるようにした
	var matching_name_indexes = []int{}

	// ヘッダテーブルの頭から探す
	for i, h := range ht {

		// Name がヘッダテーブルにあった場合
		if h.Name == name {

			// Value も一致したら
			if h.Value == value {
				// 一致した index とそこにある値を返す
				return i, h // index header
			}

			// name は一致したのでそのインデックスを加えておく
			matching_name_indexes = append(matching_name_indexes, i)
		}
	}

	// Name があっても value までは一致しなかった場合
	// 一番最初のヘッダを返す
	if len(matching_name_indexes) > 0 {
		return matching_name_indexes[0], nil // literal with index
	}

	// Name も一致しなかったら -1, nil
	return -1, nil // literal without index
}

func Search(headers http.Header, headerTable HeaderTable) {
	for name, values := range headers {
		for i := range values {
			value := values[i]
			index, h := headerTable.SearchHeader(name, value)
			if h != nil {
				log.Println("index header", index, h, name, value)
				frame := hpac.NewIndexedHeader()
				frame.Index = uint64(index)
				f := hpac.EncodeHeader(frame)
				log.Printf("%T %v", f, f.Bytes())
			} else if index != -1 {
				log.Println("literal with index", index, h, name, values)
				frame := hpac.NewIndexedNameWithIncrementalIndexing()
				frame.Index = uint64(index)
				frame.ValueLength = uint64(len(value))
				frame.ValueString = value
				f := hpac.EncodeHeader(frame)
				log.Printf("%T %v", f, f.Bytes())
			} else {
				log.Println("literal without index")
			}
		}
	}
}

func NewRequestHeaderTable() HeaderTable {
	return HeaderTable{
		{":scheme", "http"},
		{":scheme", "https"},
		{":host", ""},
		{":path", "/"},
	}
}


func main() {
	var headers = http.Header{
		":scheme":     []string{"http"},
		":path":       []string{"/index.html"},
		"mynewheader": []string{"first"},
	}

	context := hpac.NewContext()

	wire := context.Encode(headers)

	log.Println(wire)

////	headerTable := NewRequestHeaderTable()
////	Search(headers, headerTable)
}
