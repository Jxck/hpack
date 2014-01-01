package hpack

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

const TestCaseDir string = "./hpack-test-case"

// {
//   "draft": 5,
//   "context": "request",
//   "description": "Encoded request headers with Literal without index only.",
//   "cases": [
//     {
//       "header_table_size": 4096,
//       "wire": "1234567890abcdef",
//       "headers": [
//         { ":method": "GET" },
//         { ":scheme": "http" },
//         { ":authority": "example.com" },
//         { ":path": "/" },
//         { "x-my-header": "value1,value2" }
//       ]
//     },
//     .....
//   ]
// }
type TestCase struct {
	HeaderTableSize int `json:"header_table_size"`
	Wire            string
	Headers         []map[string]string
}

type TestFile struct {
	Draft       int
	Context     string
	Description string
	Cases       []TestCase
}

func toJSON(jsoncase string) TestFile {
	var test TestFile
	err := json.NewDecoder(strings.NewReader(jsoncase)).Decode(&test)
	if err != nil {
		log.Fatal(err)
	}
	return test
}

func RunStory(testfile TestFile, t *testing.T) {
	context := NewContext(testfile.Context == "request", DEFAULT_HEADER_TABLE_SIZE)

	for _, v := range testfile.Cases {
		wire, err := hex.DecodeString(v.Wire)
		if err != nil {
			log.Fatal(err)
		}
		context.Decode(wire)

		expectedHeader := make(http.Header)
		for _, v := range v.Headers {
			for v, k := range v {
				expectedHeader.Add(RemovePrefix(v), k)
			}
		}
		if !reflect.DeepEqual(context.ES.Header, expectedHeader) {
			t.Errorf("\n got %v\nwant %v", context.ES.Header, expectedHeader)
		}
	}
}

func TestStory(t *testing.T) {
	// story_06 faild
	jsoncase := `
{
  "context": "request", 
  "cases": [
    {
      "header_table_size": 4096, 
      "wire": "82870388f4466d6912d2717787", 
      "headers": [
        {
          ":method": "GET"
        }, 
        {
          ":scheme": "http"
        }, 
        {
          ":authority": "yahoo.co.jp"
        }, 
        {
          ":path": "/"
        }
      ]
    }, 
    {
      "header_table_size": 4096, 
      "wire": "028bdb6d89e88cdad225a4e2ef83", 
      "headers": [
        {
          ":method": "GET"
        }, 
        {
          ":scheme": "http"
        }, 
        {
          ":authority": "www.yahoo.co.jp"
        }, 
        {
          ":path": "/"
        }
      ]
    }, 
    {
      "header_table_size": 4096, 
      "wire": "0187e44f45699138bb439905688c860116b820e38274602db9365642534c5a0f591cccbf8283", 
      "headers": [
        {
          ":method": "GET"
        }, 
        {
          ":scheme": "http"
        }, 
        {
          ":authority": "k.yimg.jp"
        }, 
        {
          ":path": "/images/top/sp2/cmn/logo-ns-130528.png"
        }
      ]
    }
  ], 
  "description": "Encoded by nghttp2. The basic encoding strategy is described in http://lists.w3.org/Archives/Public/ietf-http-wg/2013JulSep/1135.html We use huffman encoding only if it produces strictly shorter byte string than original. We make some headers not indexing at all, but this does not always result in less bits on the wire.", 
  "draft": 5
}
	`
	testcases := toJSON(jsoncase)
	RunStory(testcases, t)
}
