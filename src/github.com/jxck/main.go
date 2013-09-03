package main

import (
	"github.com/jxck/hpac"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	var headers = http.Header{
		":method":     []string{"GET"},
		":scheme":     []string{"http"},
		":path":       []string{"/index.html"},
		"mynewheader": []string{"first"},
	}

	context := hpac.NewContext()
	wire := context.Encode(headers)

	log.Println(wire)

	// headerTable := NewRequestHeaderTable()
	// Search(headers, headerTable)
}
