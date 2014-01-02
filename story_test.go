package hpack

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"sort"
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

func readJsonFile(path string) TestFile {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	var test TestFile
	err = json.NewDecoder(file).Decode(&test)
	if err != nil {
		log.Fatal(err)
	}
	return test
}

func RunStory(testfile TestFile, t *testing.T) {
	context := NewContext(testfile.Context == "request", DEFAULT_HEADER_TABLE_SIZE)
	for _, cases := range testfile.Cases {
		wire, err := hex.DecodeString(cases.Wire)
		if err != nil {
			log.Fatal(err)
		}
		context.Decode(wire)

		expectedES := &EmittedSet{}
		for _, header := range cases.Headers {
			for key, value := range header {
				expectedES.Emit(NewHeaderField(key, value))
			}
		}

		sort.Sort(context.ES)
		sort.Sort(expectedES)

		if !reflect.DeepEqual(context.ES, expectedES) {
			t.Fatalf("actual %v expected %v", context.ES.Dump(), expectedES.Dump())
		}
	}
}

const dir string = "./hpack-test-case/node-http2-hpack/"

func TestSingleStory(t *testing.T) {
	testcases := readJsonFile(dir + "story_29.json")
	RunStory(testcases, t)
}

func TestStory(t *testing.T) {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		t.Log("==== test", dir+f.Name())
		testcases := readJsonFile(dir + f.Name())
		RunStory(testcases, t)
	}
}
