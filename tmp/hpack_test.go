package hpack

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"reflect"
	"strings"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

const TestCaseDir string = "./hpack-test-case"

type TestCase struct {
	Context string
	Wire    string
	Header  map[string]string
}

func toJSON(jsoncase string) []TestCase {
	var testcases []TestCase
	err := json.NewDecoder(strings.NewReader(jsoncase)).Decode(&testcases)
	if err != nil {
		log.Fatal(err)
	}
	return testcases
}

func RunStory(testcases []TestCase, t *testing.T) {
	var context *Context
	for i, testcase := range testcases {
		log.Printf("== test case %d =========================\n", i)
		expected := NewEmittedSet()
		for k, v := range testcase.Header {
			expected.Emit(k, v)
		}

		wire, err := base64.StdEncoding.DecodeString(testcase.Wire)
		if err != nil {
			log.Fatal(err)
		}

		if context == nil {
			if testcase.Context == "request" {
				context = NewRequestContext()
			} else if testcase.Context == "response" {
				context = NewResponseContext()
			}
		}
		context.Decode(wire)

		actual := context.EmittedSet.Header
		// log.Println(MapString(expected.Header))
		// log.Println(MapString(actual))
		for name, values := range expected.Header {
			if !reflect.DeepEqual(actual[name], values) {
				log.Println(values, actual[name])
				t.Errorf("got %v\nwant %v", values, actual[name])
			}
		}
	}
}

func TestSingleCase(t *testing.T) {
	jsoncase := `
[
  {
    "context": "request",
    "wire": "Rkp0ZXh0L2h0bWwsYXBwbGljYXRpb24veGh0bWwreG1sLGFwcGxpY2F0aW9uL3htbDtxPTAuOSxpbWFnZS93ZWJwLCovKjtxPTAuOEgRZ3ppcCxkZWZsYXRlLHNkY2hJF2VuLVVTLGVuO3E9MC44LGphO3E9MC42UAltYXgtYWdlPTBRCmtlZXAtYWxpdmVABGhvc3QSYmxvZy5zdW1tZXJ3aW5kLmpwTHdNb3ppbGxhLzUuMCAoTWFjaW50b3NoOyBJbnRlbCBNYWMgT1MgWCAxMF84XzUpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8zMC4wLjE1OTkuNjkgU2FmYXJpLzUzNy4zNg==",
    "header": {
      "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
      "accept-encoding": "gzip,deflate,sdch",
      "accept-language": "en-US,en;q=0.8,ja;q=0.6",
      "cache-control": "max-age=0",
      "connection": "keep-alive",
      ":host": "blog.summerwind.jp",
      "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.69 Safari/537.36"
    }
  }
]
	`
	testcases := toJSON(jsoncase)
	for _, testcase := range testcases {
		RunStory([]TestCase{testcase}, t)
	}
}

func TestStory(t *testing.T) {
	// story_06 faild
	jsoncase := `
[
  {
    "context": "request",
    "wire": "hIBDEmdlby5jcmFpZ3NsaXN0Lm9yZ4NMUU1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwLjg7IHJ2OjE2LjApIEdlY2tvLzIwMTAwMTAxIEZpcmVmb3gvMTYuMEY/dGV4dC9odG1sLGFwcGxpY2F0aW9uL3hodG1sK3htbCxhcHBsaWNhdGlvbi94bWw7cT0wLjksKi8qO3E9MC44SQ5lbi1VUyxlbjtxPTAuNUgNZ3ppcCwgZGVmbGF0ZVEKa2VlcC1hbGl2ZUogY2xfYj1BQjJCS2JzbDRoR003TTRuSDVQWVdnaFRNNUE=",
    "header": {
      ":method": "GET",
      ":scheme": "http",
      ":host": "geo.craigslist.org",
      ":path": "/",
      "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0",
      "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
      "accept-language": "en-US,en;q=0.5",
      "accept-encoding": "gzip, deflate",
      "connection": "keep-alive",
      "cookie": "cl_b=AB2BKbsl4hGM7M4nH5PYWghTM5A"
    }
  },
  {
    "context": "request",
    "wire": "noNfABJ3d3cuY3JhaWdzbGlzdC5vcmdEDS9hYm91dC9zaXRlcy8=",
    "header": {
      ":method": "GET",
      ":scheme": "http",
      ":host": "www.craigslist.org",
      ":path": "/about/sites/",
      "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0",
      "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
      "accept-language": "en-US,en;q=0.5",
      "accept-encoding": "gzip, deflate",
      "connection": "keep-alive",
      "cookie": "cl_b=AB2BKbsl4hGM7M4nH5PYWghTM5A"
    }
  },
  {
    "context": "request",
    "wire": "oKZfCBUvc3R5bGVzL2NvdW50cmllcy5jc3NfAhJ0ZXh0L2NzcywqLyo7cT0wLjFNJmh0dHA6Ly93d3cuY3JhaWdzbGlzdC5vcmcvYWJvdXQvc2l0ZXMv",
    "header": {
      ":method": "GET",
      ":scheme": "http",
      ":host": "www.craigslist.org",
      ":path": "/styles/countries.css",
      "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0",
      "accept": "text/css,*/*;q=0.1",
      "accept-language": "en-US,en;q=0.5",
      "accept-encoding": "gzip, deflate",
      "connection": "keep-alive",
      "referer": "http://www.craigslist.org/about/sites/",
      "cookie": "cl_b=AB2BKbsl4hGM7M4nH5PYWghTM5A"
    }
  },
  {
    "context": "request",
    "wire": "p6hfCQ4vanMvZm9ybWF0cy5qc18KAyovKg==",
    "header": {
      ":method": "GET",
      ":scheme": "http",
      ":host": "www.craigslist.org",
      ":path": "/js/formats.js",
      "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0",
      "accept": "*/*",
      "accept-language": "en-US,en;q=0.5",
      "accept-encoding": "gzip, deflate",
      "connection": "keep-alive",
      "referer": "http://www.craigslist.org/about/sites/",
      "cookie": "cl_b=AB2BKbsl4hGM7M4nH5PYWghTM5A"
    }
  },
  {
    "context": "request",
    "wire": "ql8MEy9qcy9qcXVlcnktMS40LjIuanM=",
    "header": {
      ":method": "GET",
      ":scheme": "http",
      ":host": "www.craigslist.org",
      ":path": "/js/jquery-1.4.2.js",
      "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0",
      "accept": "*/*",
      "accept-language": "en-US,en;q=0.5",
      "accept-encoding": "gzip, deflate",
      "connection": "keep-alive",
      "referer": "http://www.craigslist.org/about/sites/",
      "cookie": "cl_b=AB2BKbsl4hGM7M4nH5PYWghTM5A"
    }
  },
  {
    "context": "request",
    "wire": "qausXw4ML2Zhdmljb24uaWNvXw0haW1hZ2UvcG5nLGltYWdlLyo7cT0wLjgsKi8qO3E9MC41",
    "header": {
      ":method": "GET",
      ":scheme": "http",
      ":host": "www.craigslist.org",
      ":path": "/favicon.ico",
      "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0",
      "accept": "image/png,image/*;q=0.8,*/*;q=0.5",
      "accept-language": "en-US,en;q=0.5",
      "accept-encoding": "gzip, deflate",
      "connection": "keep-alive",
      "cookie": "cl_b=AB2BKbsl4hGM7M4nH5PYWghTM5A"
    }
  },
  {
    "context": "request",
    "wire": "pa2uXwcVc2hvYWxzLmNyYWlnc2xpc3Qub3Jng6Cp",
    "header": {
      ":method": "GET",
      ":scheme": "http",
      ":host": "shoals.craigslist.org",
      ":path": "/",
      "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0",
      "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
      "accept-language": "en-US,en;q=0.5",
      "accept-encoding": "gzip, deflate",
      "connection": "keep-alive",
      "referer": "http://www.craigslist.org/about/sites/",
      "cookie": "cl_b=AB2BKbsl4hGM7M4nH5PYWghTM5A"
    }
  },
  {
    "context": "request",
    "wire": "pK+DoKmlXw8WL3N0eWxlcy9jcmFpZ3NsaXN0LmNzc6hfCx1odHRwOi8vc2hvYWxzLmNyYWlnc2xpc3Qub3JnL18GQmNsX2I9QUIyQktic2w0aEdNN000bkg1UFlXZ2hUTTVBOyBjbF9kZWZfbGFuZz1lbjsgY2xfZGVmX2hwPXNob2Fscw==",
    "header": {
      ":method": "GET",
      ":scheme": "http",
      ":host": "www.craigslist.org",
      ":path": "/styles/craigslist.css",
      "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0",
      "accept": "text/css,*/*;q=0.1",
      "accept-language": "en-US,en;q=0.5",
      "accept-encoding": "gzip, deflate",
      "connection": "keep-alive",
      "referer": "http://shoals.craigslist.org/",
      "cookie": "cl_b=AB2BKbsl4hGM7M4nH5PYWghTM5A; cl_def_lang=en; cl_def_hp=shoals"
    }
  },
  {
    "context": "request",
    "wire": "sKiqqw==",
    "header": {
      ":method": "GET",
      ":scheme": "http",
      ":host": "www.craigslist.org",
      ":path": "/js/formats.js",
      "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0",
      "accept": "*/*",
      "accept-language": "en-US,en;q=0.5",
      "accept-encoding": "gzip, deflate",
      "connection": "keep-alive",
      "referer": "http://shoals.craigslist.org/",
      "cookie": "cl_b=AB2BKbsl4hGM7M4nH5PYWghTM5A; cl_def_lang=en; cl_def_hp=shoals"
    }
  },
  {
    "context": "request",
    "wire": "ql8SDy9qcy9ob21lcGFnZS5qcw==",
    "header": {
      ":method": "GET",
      ":scheme": "http",
      ":host": "www.craigslist.org",
      ":path": "/js/homepage.js",
      "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0",
      "accept": "*/*",
      "accept-language": "en-US,en;q=0.5",
      "accept-encoding": "gzip, deflate",
      "connection": "keep-alive",
      "referer": "http://shoals.craigslist.org/",
      "cookie": "cl_b=AB2BKbsl4hGM7M4nH5PYWghTM5A; cl_def_lang=en; cl_def_hp=shoals"
    }
  }
]
	`
	testcases := toJSON(jsoncase)
	RunStory(testcases, t)
}
