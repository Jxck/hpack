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
	HeaderTableSize int                 `json:"header_table_size"`
	Wire            string              `json:"wire"`
	Headers         []map[string]string `json:"headers"`
}

type TestFile struct {
	Draft       int        `json:"draft"`
	Context     string     `json:"context"`
	Description string     `json:"description"`
	Cases       []TestCase `json:"cases"`
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

// func TestSingleStory(t *testing.T) {
// 	testcases := readJsonFile(dir + "story_29.json")
// 	RunStory(testcases, t)
// }

func TestStory(t *testing.T) {
	t.Skip()
	dirs := []string{
		"./hpack-test-case/node-http2-hpack/",
		"./hpack-test-case/nghttp2/",
		"./hpack-test-case/haskell-http2/",
		"./hpack-test-case/go-hpack/",
	}

	for _, dir := range dirs {
		files, _ := ioutil.ReadDir(dir)
		for _, f := range files {
			t.Log("==== test", dir+f.Name())
			testcases := readJsonFile(dir + f.Name())
			RunStory(testcases, t)
		}
	}
}

func writeJson(src, dst, filename string) {
	testFile := readJsonFile(src + filename)

	testFile.Draft = 5
	testFile.Description = "https://github.com/jxck/hpack implemeted in Golang. Encoded using String Literal, no Header/Static Table, and always start with emptied Reference Set. by Jxck."

	context := NewContext(testFile.Context == "request", DEFAULT_HEADER_TABLE_SIZE)
	// 一つのケースごと
	for i, c := range testFile.Cases {
		hs := HeaderSet{}
		// 一つのヘッダごと
		for _, header := range c.Headers {
			for key, value := range header {
				hs = append(hs, NewHeaderField(key, value))
			}
		}

		hexdump := context.Encode(hs)
		wire := hex.EncodeToString(hexdump)
		testFile.Cases[i].Wire = wire
		testFile.Cases[i].HeaderTableSize = DEFAULT_HEADER_TABLE_SIZE
	}

	b, _ := json.MarshalIndent(testFile, "", "  ")
	file, err := os.Create(dst + filename)
	if err != nil {
		log.Fatal(err)
	}
	n, err := file.Write(b)
	if n == 0 || err != nil {
		log.Fatal("failt to write file")
	}
}

func TestEncodeStory(t *testing.T) {
	src := "./hpack-test-case/raw-data/"
	dst := "./hpack-test-case/go-hpack/"
	files, _ := ioutil.ReadDir(src)
	for _, file := range files {
		writeJson(src, dst, file.Name())
	}
}
