package hpack

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
	"testing"
)

type R struct {
	HeaderTable Headers `json:"header_table"`
	Hex         string  `json:"hpack_encoded_hex"`
	HeaderSet   Headers `json:"header_set"`
}

type JsonCase struct {
	Request  R
	Response R
}

func toJsonCase(samplecase string) []JsonCase {
	var testcases []JsonCase
	err := json.NewDecoder(strings.NewReader(samplecase)).Decode(&testcases)
	if err != nil {
		log.Fatal(err)
	}
	return testcases
}

func TestTmp(t *testing.T) {
	t.Skip()
	file, err := ioutil.ReadFile("./main/yahoo.co.jp_hpack_encoded.json")
	if err != nil {
		t.Fatal(err)
	}

	var samplecase string = string(file)
	context := NewRequestContext()
	for i, v := range toJsonCase(samplecase) {
		log.Println("=========case", i)
		wire, err := hex.DecodeString(v.Request.Hex)
		if err != nil {
			t.Fatal(err)
		}
		context.Decode(wire)

		// check HeaderTable
		actual_ht := context.HeaderTable.Headers
		expected_ht := v.Request.HeaderTable
		if !reflect.DeepEqual(actual_ht, expected_ht) {
			t.Log(HeadersString(actual_ht))
			t.Log(HeadersString(expected_ht))
			t.Errorf("got %v\nwant %v", actual_ht, expected_ht)
		}

		// check HeaderSet
		expected_es := NewEmittedSet()
		actual_es := context.EmittedSet
		for _, hs := range v.Request.HeaderSet {
			expected_es.Emit(hs.Name, hs.Value)
		}
		if !reflect.DeepEqual(actual_es, expected_es) {
			t.Log(HeaderString(actual_es.Header))
			t.Log(HeaderString(expected_es.Header))
			t.Log(RefSetString(context.ReferenceSet))
			t.Errorf("got %v\nwant %v", actual_es, expected_es)
			t.FailNow()
		}
	}
}
