package main

import (
	"log"
)

type HeaderName string
type HeaderValue string

type Header struct {
	Name  HeaderName
	Value HeaderValue
}

type HeaderTable []*Header // Header はポインタにしておく

//func (ht *HeaderTable) Add(header Header) {
//	*ht = append(*ht, header)
//}

func (ht HeaderTable) SearchHeader(name HeaderName) int {
	for i, h := range ht {
		if h.Name == name {
			return i
		}
	}
	return -1
}

func NewRequestHeaderTable() HeaderTable {
	return HeaderTable{
		{":scheme", "http"},
		{":scheme", "https"},
		{":host", "A"},
		{":path", "/"},
	}
}
func main() {
	log.SetFlags(log.Lshortfile)

	header := NewRequestHeaderTable()

	i := header.SearchHeader(":host")
	log.Println(header[i])
}
