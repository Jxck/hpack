package main

import (
	"github.com/jxck/hpac"
	"log"
	"net/http"
	"fmt"
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

	fmt.Println(server)
	fmt.Println("======================")

	headers = http.Header{
		":method":     []string{"GET"},
		":scheme":     []string{"http"},
		":path":       []string{"/index.html"},
		"mynewheader": []string{"second"},
	}
	wire = client.Encode(headers)
	log.Println(wire)
	server.Decode(wire)
	fmt.Println(server)
}

/*

refset
  ":scheme":     "http"
  "hoge"   :     "fuga"

header
  ":method":     []string{"GET"},
  ":scheme":     []string{"http"},
  ":path":       []string{"/index.html"},
  "mynewheader": []string{"first"},

context.go:56: remove from refset hoge fuga
context.go:77: remove from header set :scheme http
context.go:91: indexed header {:method:GET} is in HT[4] ([132])
context.go:98: literal with index {:path:/index.html} is in HT[3] ([68 11 47 105 110 100 101 120 46 104 116 109 108])
context.go:105: literal without index {mynewheader:first} is not in HT ([96 11 109 121 110 101 119 104 101 97 100 101 114 5 102 105 114 115 116])
context.go:43: map[:scheme:http]
main.go:24: [255 128 255 255 255 255 255 255 255 255 1 132 68 11 47 105 110 100 101 120 46 104 116 109 108 96 11 109 121 110 101 119 104 101 97 100 101 114 5 102 105 114 115 116]
decoder.go:37: Indexed Header Representation
decoder.go:17: 33
decoder.go:37: Indexed Header Representation
decoder.go:17: 32
decoder.go:90: Literal Header with Incremental Indexing - Indexed Name
decoder.go:17: 19
decoder.go:74: Literal Header without Indexing - New Name
decoder.go:17: 0
decoder.go:19: [0x1846e390 0x1846e3b0 0x18467ca0 0x184663c0]
*/
