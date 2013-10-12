package hpack

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestHpack(t *testing.T) {
	jsoncase := `{
   "wire": "Rkp0ZXh0L2h0bWwsYXBwbGljYXRpb24veGh0bWwreG1sLGFwcGxpY2F0aW9uL3htbDtxPTAuOSxpbWFnZS93ZWJwLCovKjtxPTAuOEgRZ3ppcCxkZWZsYXRlLHNkY2hJF2VuLVVTLGVuO3E9MC44LGphO3E9MC42UAltYXgtYWdlPTBRCmtlZXAtYWxpdmVABGhvc3QSYmxvZy5zdW1tZXJ3aW5kLmpwTHdNb3ppbGxhLzUuMCAoTWFjaW50b3NoOyBJbnRlbCBNYWMgT1MgWCAxMF84XzUpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8zMC4wLjE1OTkuNjkgU2FmYXJpLzUzNy4zNg==",
   "header": {
     "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
     "accept-encoding": "gzip,deflate,sdch",
     "accept-language": "en-US,en;q=0.8,ja;q=0.6",
     "cache-control": "max-age=0",
     "connection":  "keep-alive",
     "host": "blog.summerwind.jp",
     "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.69 Safari/537.36"
   }
 }`

	type TestCase struct {
		Wire   string
		Header map[string]string
	}

	var testcase TestCase
	dec := json.NewDecoder(strings.NewReader(jsoncase))
	err := dec.Decode(&testcase)
	if err != nil {
		log.Fatal(err)
	}

	header := http.Header{}
	for k, v := range testcase.Header {
		header.Add(k, v)
	}

	wire, err := base64.StdEncoding.DecodeString(testcase.Wire)
	if err != nil {
		log.Fatal(err)
	}
	context := NewRequestContext()
	context.Decode(wire)

	for name, values := range context.EmittedSet.Header {
		if !CompareSlice(header[name], values) {
			log.Println(values, header[name])
			t.Errorf("got %v\nwant %v", values, header[name])
		}
	}
}
