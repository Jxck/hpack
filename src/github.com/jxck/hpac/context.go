package hpac

import (
	"log"
	"net/http"
)

type Context struct {
	requestHeaderTable  HeaderTable
	responseHeaderTable HeaderTable
	referenceSet        ReferenceSet
}

func NewContext() *Context {
	var context = &Context{
		requestHeaderTable:  NewRequestHeaderTable(),
		responseHeaderTable: NewResponseHeaderTable(),
		referenceSet: ReferenceSet{
			// TODO: test data
			":scheme": "http",
			"hoge":    "fuga",
		},
	}
	return context
}

func (c *Context) Encode(header http.Header) []byte {
	// http.Header を HeaderSet に変換
	headerSet := NewHeaderSet(header)

	// ReferenceSet の中から消すべき値を消す
	c.CleanReferenceSet(headerSet)

	// Header Set の中から送らない値を消す
	c.CleanHeaderSet(headerSet)

	// Header Table にあるやつを indexed で emit
	c.ProcessHeader(headerSet)

	return nil
}

// 1. 不要なエントリを reference set から消す
func (c *Context) CleanReferenceSet(headerSet HeaderSet) {
	// reference set の中にあって、 header set の中に無いものは
	// 相手の reference set から消さないといけないので、
	// indexed representation でエンコードして
	// reference set からは消す
	for name, value := range c.referenceSet {
		if headerSet[name] != value {
			// TODO: integer representation でエンコード
			log.Println("remove from refset", name, value)
		}
	}
}

// 2. 送る必要の無いものを header set から消す
func (c *Context) CleanHeaderSet(headerSet HeaderSet) {
	for name, value := range c.referenceSet {
		if headerSet[name] == value {
			delete(headerSet, name)
			// TODO: "common-header" としてマーク
			log.Println("remove from header set", name, value)
		}
	}
}

// 3 と 4. 残りの処理
func (c *Context) ProcessHeader(headerSet HeaderSet) {
	for name, value := range headerSet {
		index, h := c.requestHeaderTable.SearchHeader(name, value)
		if h != nil { // 3.1 HT にエントリがある
			frame := CreateIndexedHeader(uint64(index))
			f := EncodeHeader(frame)
			log.Printf("indexed header {%v:%v} is in HT[%v] (%v)", name, value, index, f.Bytes())
		} else if index != -1 { // HT に name だけある
			frame := CreateIndexedNameWithIncrementalIndexing(uint64(index), value)
			f := EncodeHeader(frame)
			log.Printf("literal with index [%v:%v] is in HT[%v] %v", name, value, index, f.Bytes())
		} else { // HT に name も value もない
			frame := CreateNewNameWithoutIndexing(name, value)
			f := EncodeHeader(frame)
			log.Printf("literal without index [%v:%v] is not in HT %v", name, value, f.Bytes())
		}
	}
}
