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
		referenceSet:        ReferenceSet{},
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

func (c *Context) CleanReferenceSet(headerSet HeaderSet) {
	for name, value := range c.referenceSet {
		if headerSet[name] != value {
			log.Println("remove from refset", name, value)
		}
	}

}
