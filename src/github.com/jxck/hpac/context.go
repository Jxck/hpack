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
		referenceSet:        ReferenceSet{
			// TODO: test data
			":scheme": "http",
			"hoge": "fuga",
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
	c.EmitIndexedInHeaderTable(headerSet)

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

// 3.1. Header Table にあるやつを indexed で emit
func (c *Context) EmitIndexedInHeaderTable(headerSet HeaderSet) {
	// Header Table にあるものは、 indexed で送れる
	for name, value := range headerSet {
		i, h := c.requestHeaderTable.SearchHeader(name, value)
		log.Println("search header", name, value, i, h)
	}
}
