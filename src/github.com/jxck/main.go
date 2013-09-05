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

	client := hpac.NewContext()
	wire := client.Encode(headers)

	server := hpac.NewContext()
	server.Decode(wire)

	log.Println("======================")

	headers = http.Header{
		":method":     []string{"GET"},
		":scheme":     []string{"http"},
		":path":       []string{"/index.html"},
		"mynewheader": []string{"second"},
	}
	wire = client.Encode(headers)
	server.Decode(wire)
}
