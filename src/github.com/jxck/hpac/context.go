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
