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


func main() {
	var header = http.Header{
		":scheme":     []string{"http"},
		":path":       []string{"/index.html"},
		"mynewheader": []string{"first"},
	}
	_ = header

	//	TestIndexedHeader()
	TestNewNameWithoutIndexing()
	//TestIndexedNameWithoutIndexing()
}
