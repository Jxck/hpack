package hpack

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"reflect"
	"strings"
	"testing"
)

var samplecase string = `
[{
  "response": {
    "header_table": [
      { "value": "200", "name": ":status" }, { "value": null, "name": "age" }, { "value": null, "name": "cache-control" }, { "value": null, "name": "content-length" }, { "value": null, "name": "content-type" }, { "value": null, "name": "date" }, { "value": null, "name": "etag" }, { "value": null, "name": "expires" }, { "value": null, "name": "last-modified" }, { "value": null, "name": "server" }, { "value": null, "name": "set-cookie" }, { "value": null, "name": "vary" }, { "value": null, "name": "via" }, { "value": null, "name": "access-control-allow-origin" }, { "value": null, "name": "accept-ranges" }, { "value": null, "name": "allow" }, { "value": null, "name": "connection" }, { "value": null, "name": "content-disposition" }, { "value": null, "name": "content-encoding" }, { "value": null, "name": "content-language" }, { "value": null, "name": "content-location" }, { "value": null, "name": "content-range" }, { "value": null, "name": "link" }, { "value": null, "name": "location" }, { "value": null, "name": "proxy-authenticate" }, { "value": null, "name": "refresh" }, { "value": null, "name": "retry-after" }, { "value": null, "name": "strict-transport-security" }, { "value": null, "name": "transfer-encoding" }, { "value": null, "name": "www-authenticate" } ],
    "hpack_encoded_hex": "4103333031461d5361742c203033204e6f7620323031322031333a32363a313020474d545817687474703a2f2f7777772e7961686f6f2e636f2e6a702f4c0f4163636570742d456e636f64696e675105636c6f73655d076368756e6b65644518746578742f68746d6c3b20636861727365743d7574662d38430770726976617465",
    "header_set": [
      { "value": "301", "name": ":status" }, { "value": "Sat, 03 Nov 2012 13:26:10 GMT", "name": "date" }, { "value": "http://www.yahoo.co.jp/", "name": "location" }, { "value": "Accept-Encoding", "name": "vary" }, { "value": "close", "name": "connection" }, { "value": "chunked", "name": "transfer-encoding" }, { "value": "text/html; charset=utf-8", "name": "content-type" }, { "value": "private", "name": "cache-control" } ] },
	"request": {
		"header_table": [ { "value": "http", "name": ":scheme" }, { "value": "https", "name": ":scheme" }, { "value": null, "name": ":host" }, { "value": "/", "name": ":path" }, { "value": "GET", "name": ":method" }, { "value": null, "name": "accept" }, { "value": null, "name": "accept-charset" }, { "value": null, "name": "accept-encoding" }, { "value": null, "name": "accept-language" }, { "value": null, "name": "cookie" }, { "value": null, "name": "if-modified-since" }, { "value": null, "name": "user-agent" }, { "value": null, "name": "referer" }, { "value": null, "name": "authorization" }, { "value": null, "name": "allow" }, { "value": null, "name": "cache-control" }, { "value": null, "name": "connection" }, { "value": null, "name": "content-length" }, { "value": null, "name": "content-type" }, { "value": null, "name": "date" }, { "value": null, "name": "expect" }, { "value": null, "name": "from" }, { "value": null, "name": "if-match" }, { "value": null, "name": "if-none-match" }, { "value": null, "name": "if-range" }, { "value": null, "name": "if-unmodified-since" }, { "value": null, "name": "max-forwards" }, { "value": null, "name": "proxy-authorization" }, { "value": null, "name": "range" }, { "value": null, "name": "via" }, { "value": "yahoo.co.jp", "name": ":host" }, { "value": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0", "name": "user-agent" }, { "value": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8", "name": "accept" }, { "value": "en-US,en;q=0.5", "name": "accept-language" }, { "value": "gzip, deflate", "name": "accept-encoding" }, { "value": "keep-alive", "name": "connection" }, { "value": "B=76j09a189a6h4&b=3&s=0b", "name": "cookie" } ],
    "hpack_encoded_hex": "8084430b7961686f6f2e636f2e6a70834c514d6f7a696c6c612f352e3020284d6163696e746f73683b20496e74656c204d6163204f5320582031302e383b2072763a31362e3029204765636b6f2f32303130303130312046697265666f782f31362e30463f746578742f68746d6c2c6170706c69636174696f6e2f7868746d6c2b786d6c2c6170706c69636174696f6e2f786d6c3b713d302e392c2a2f2a3b713d302e38490e656e2d55532c656e3b713d302e35480d677a69702c206465666c617465510a6b6565702d616c6976654a18423d37366a3039613138396136683426623d3326733d3062",
    "header_set": [ { "value": "http", "name": ":scheme" }, { "value": "GET", "name": ":method" }, { "value": "yahoo.co.jp", "name": ":host" }, { "value": "/", "name": ":path" }, { "value": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0", "name": "user-agent" }, { "value": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8", "name": "accept" }, { "value": "en-US,en;q=0.5", "name": "accept-language" }, { "value": "gzip, deflate", "name": "accept-encoding" }, { "value": "keep-alive", "name": "connection" }, { "value": "B=76j09a189a6h4&b=3&s=0b", "name": "cookie" }
    ]
  }
}]
`

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
	for _, v := range toJsonCase(samplecase) {
		wire, err := hex.DecodeString(v.Request.Hex)
		if err != nil {
			t.Fatal(err)
		}
		context := NewRequestContext()
		context.Decode(wire)

		// check HeaderTable
		actual_ht := context.HeaderTable.Headers
		expected_ht := v.Request.HeaderTable
		if !reflect.DeepEqual(actual_ht, expected_ht) {
			t.Errorf("got %v\nwant %v", actual_ht, expected_ht)
		}

		// check HeaderSet
		expected_es := NewEmittedSet()
		actual_es := context.EmittedSet
		for _, hs := range v.Request.HeaderSet {
			expected_es.Emit(hs.Name, hs.Value)
		}
		if !reflect.DeepEqual(actual_es, expected_es) {
			t.Errorf("got %v\nwant %v", actual_es, expected_es)
		}
	}
}
