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
		"user-agent": []string{"http2cat"},
		"cookie":     []string{"xxxxxxx1"},
		"x-hello":    []string{"world"},
	}

	client := hpac.NewContext()
	wire := client.Encode(headers)

	server := hpac.NewContext()
	server.Decode(wire)

	fmt.Println("======================")

	headers = http.Header{
		":scheme":    []string{"https"},
		":host":      []string{"jxck.io"},
		":path":      []string{"/labs/http2cat"},
		":method":    []string{"GET"},
		"user-agent": []string{"http2cat"},
		"cookie":     []string{"xxxxxxx2"},
	}
	wire = client.Encode(headers)
	server.Decode(wire)
}
