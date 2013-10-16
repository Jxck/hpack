package hpack

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

const TestCaseDir string = "./hpack-test-case"

func TestStory(t *testing.T) {
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

	type TestCase struct {
		Context string
		Wire    string
		Header  map[string]string
	}

	var testcases []TestCase
	dec := json.NewDecoder(strings.NewReader(jsoncase))
	err := dec.Decode(&testcases)
	if err != nil {
		log.Fatal(err)
	}

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
		log.Println(MapString(expected.Header))
		log.Println(MapString(actual))
		for name, values := range expected.Header {
			if !CompareSlice(actual[name], values) {
				log.Println(values, actual[name])
				t.Errorf("got %v\nwant %v", values, actual[name])
			}
		}
	}
}

func RunCase(filename string, t *testing.T) {
	data, err := ioutil.ReadFile(TestCaseDir + "/" + filename)
	if err != nil {
		t.Fatal()
	}
	jsoncase := string(data)

	type TestCase struct {
		Context string
		Wire    string
		Header  map[string]string
	}

	var testcases []TestCase
	dec := json.NewDecoder(strings.NewReader(jsoncase))
	err = dec.Decode(&testcases)
	if err != nil {
		log.Fatal(err)
	}

	for _, testcase := range testcases {
		es := NewEmittedSet()
		for k, v := range testcase.Header {
			es.Add(k, v)
		}
		header := es.Header

		wire, err := base64.StdEncoding.DecodeString(testcase.Wire)
		if err != nil {
			log.Fatal(err)
		}
		var context *Context
		if testcase.Context == "request" {
			context = NewRequestContext()
		} else if testcase.Context == "response" {
			context = NewResponseContext()
		}
		context.Decode(wire)

		for name, values := range context.EmittedSet.Header {
			if !CompareSlice(header[name], values) {
				log.Println(values, header[name])
				t.Errorf("got %v\nwant %v", values, header[name])
			}
		}
	}
}

func RunStory(filename string, t *testing.T) {
	data, err := ioutil.ReadFile(TestCaseDir + "/" + filename)
	if err != nil {
		t.Fatal()
	}
	jsoncase := string(data)

	type TestCase struct {
		Context string
		Wire    string
		Header  map[string]string
	}

	var testcases []TestCase
	dec := json.NewDecoder(strings.NewReader(jsoncase))
	err = dec.Decode(&testcases)
	if err != nil {
		log.Fatal(err)
	}

	var context *Context
	for i, testcase := range testcases {
		log.Printf("== test case %d =========================\n", i)
		header := http.Header{}
		for k, v := range testcase.Header {
			header.Add(k, v)
		}

		wire, err := base64.StdEncoding.DecodeString(testcase.Wire)
		if err != nil {
			fmt.Println(filename, i, testcase)
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

		for name, values := range context.EmittedSet.Header {
			if !CompareSlice(header[name], values) {
				log.Println(values, header[name])
				t.Errorf("got %v\nwant %v", values, header[name])
			}
		}
	}
}

func TestHpack(t *testing.T) {
	t.Skip() // TODO:fix me
	files, err := ioutil.ReadDir(TestCaseDir)
	if err != nil {
		t.Fatal()
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "case") {
			RunCase(file.Name(), t)
		}
		if strings.HasPrefix(file.Name(), "story") {
			RunStory(file.Name(), t)
		}
	}
}
