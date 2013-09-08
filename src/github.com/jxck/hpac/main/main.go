package main

import (
	"fmt"
	"github.com/jxck/hpac"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	var headers = http.Header{
		":scheme":    []string{"https"},
		":host":      []string{"jxck.io"},
		":path":      []string{"/"},
		":method":    []string{"GET"},
		"User-Agent": []string{"http2cat"},
		"Cookie":     []string{"xxxxxxx1"},
		"X-Hello":    []string{"world"},
	}

	client := hpac.NewContext()
	wire := client.Encode(headers)

	server := hpac.NewContext()
	server.Decode(wire)

	log.Printf("refset: %v", server.ReferenceSet)
	log.Printf("emitted: %v", server.EmittedSet)
	fmt.Println("======================")

	headers = http.Header{
		":scheme":    []string{"https"},
		":host":      []string{"jxck.io"},
		":path":      []string{"/labs/http2cat"},
		":method":    []string{"GET"},
		"User-Agent": []string{"http2cat"},
		"Cookie":     []string{"xxxxxxx2"},
	}
	wire = client.Encode(headers)
	server.Decode(wire)
	log.Printf("refset: %v", server.ReferenceSet)
	log.Printf("emitted: %v", server.EmittedSet)
}
