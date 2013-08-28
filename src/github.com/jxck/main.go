package main

import (
	"log"
	"net/http"
)

type Header struct {
	Name  string
	Value string
}

// ヘッダはポインタにしておく
type HeaderTable []*Header

//func (ht *HeaderTable) Add(header Header) {
//	*ht = append(*ht, header)
//}

func (ht HeaderTable) SearchHeader(name, value string) (int, *Header) {
	for i, h := range ht {
		if h.Name == name {
			if h.Value == value {
				return i, h
			}
		}
	}
	return -1, nil
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

	var header = http.Header{
		":scheme":     []string{"http"},
		":path":       []string{"/index.html"},
		"mynewheader": []string{"first"},
	}

	log.SetFlags(log.Lshortfile)

	_ = header

	headerTable := NewRequestHeaderTable()

	log.Println(headerTable.SearchHeader(":scheme", "http"))
	log.Println(headerTable.SearchHeader(":path", "/index.html"))
}
