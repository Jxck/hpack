package hpack

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestHpack(t *testing.T) {
	data, err := ioutil.ReadFile("./hpack-test-case/case_0.json")
	if err != nil {
		t.Fatal()
	}
	jsoncase := string(data)

	type TestCase struct {
		Wire   string
		Header map[string]string
	}

	var testcases []TestCase
	dec := json.NewDecoder(strings.NewReader(jsoncase))
	err = dec.Decode(&testcases)
	if err != nil {
		log.Fatal(err)
	}

	testcase := testcases[0]
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
