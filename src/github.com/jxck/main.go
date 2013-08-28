package main

import (
	"bytes"
	. "github.com/jxck/hpac"
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
	var matching_name_indexes = []int{}
	for i, h := range ht {
		if h.Name == name {
			if h.Value == value {
				log.Println("index header")
				return i, h
			}
			matching_name_indexes = append(matching_name_indexes, i)
		}
	}
	if len(matching_name_indexes) > 0 {
		log.Println("literal with index")
		return matching_name_indexes[0], nil
	}
	log.Println("literal without index")
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
	log.SetFlags(log.Lshortfile)

	var headers = http.Header{
		":scheme":     []string{"http"},
		":path":       []string{"/index.html"},
		"mynewheader": []string{"first"},
	}

	headerTable := NewRequestHeaderTable()
	for name, values := range headers {
		for i := range values {
			value := values[i]
			log.Println(headerTable.SearchHeader(name, value))
		}
	}
}
