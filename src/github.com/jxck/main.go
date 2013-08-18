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

	buf := EncodeInteger(10, 5)
//	log.Printf("%v %v", buf.Bytes(), len(buf.Bytes()))
//	for i, j := range buf.Bytes() {
//		log.Printf("%v, %b", i, j)
//	}

	buf = EncodeInteger(1337, 5)
	_ = buf
//	log.Printf("%v %v", buf.Bytes(), len(buf.Bytes()))
//	for i, j := range buf.Bytes() {
//		log.Printf("%v, %b", i, j)
//	}

}
