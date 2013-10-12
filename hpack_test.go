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
	jsoncase := `[
  {
    "context": "request",
    "wire": "hIBDC3lhaG9vLmNvLmpwgw==",
    "header": {
      "Method": "GET",
      "Scheme": "http",
      "Host": "yahoo.co.jp",
      "Path": "/"
    }
  },
  {
    "context": "request",
    "wire": "nl8AD3d3dy55YWhvby5jby5qcA==",
    "header": {
      "Method": "GET",
      "Scheme": "http",
      "Host": "www.yahoo.co.jp",
      "Path": "/"
    }
  },
  {
    "context": "request",
    "wire": "g59fAQlrLnlpbWcuanBEJi9pbWFnZXMvdG9wL3NwMi9jbW4vbG9nby1ucy0xMzA1MjgucG5n",
    "header": {
      "Method": "GET",
      "Scheme": "http",
      "Host": "k.yimg.jp",
      "Path": "/images/top/sp2/cmn/logo-ns-130528.png"
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
		header := http.Header{}
		for k, v := range testcase.Header {
			header.Add(k, v)
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

		for name, values := range context.EmittedSet.Header {
			if !CompareSlice(header[name], values) {
				log.Println(values, header[name])
				t.Errorf("got %v\nwant %v", values, header[name])
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
		header := http.Header{}
		for k, v := range testcase.Header {
			header.Add(k, v)
		}

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
