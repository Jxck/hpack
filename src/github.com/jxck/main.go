package main

import (
	. "github.com/jxck/hpac"
	"log"
	"net/http"
)

var header = http.Header{
	":path":       []string{"/my-example/index.html"},
	"user-agent":  []string{"my-user-agent"},
	"mynewheader": []string{"first"},
}

func main() {
	log.SetFlags(log.Lshortfile)

	log.Println(EncodeInteger(10, 5))
}
