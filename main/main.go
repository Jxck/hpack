package main

import (
	"encoding/base64"
	"fmt"
	"github.com/jxck/hpack"
	. "github.com/jxck/logger"
	"log"
	"net/http"
	"os"
)

func init() {
	log.SetFlags(log.Lshortfile)
	LogLevel(4)
	Verbose(true)
}

func main() {
	var headers = http.Header{
		":scheme":    []string{"https"},
		":host":      []string{"example.com"},
		":path":      []string{"/"},
		":method":    []string{"GET"},
		"User-Agent": []string{"hpack-test-case"},
		"Cookie":     []string{"xxxxxxx1"},
		"X-Hello":    []string{"world"},
	}

	client := hpack.NewRequestContext()
	wire := client.Encode(headers)

	str := base64.StdEncoding.EncodeToString(wire)
	fmt.Println(str)
	headers.Write(os.Stdout)

	server := hpack.NewRequestContext()
	server.Decode(wire)

	log.Printf("refset: %v", server.ReferenceSet)
	log.Printf("emitted: %v", server.EmittedSet)
	log.Println("======================")

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
