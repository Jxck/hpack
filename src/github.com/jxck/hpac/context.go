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
	// http.Header $B$r(B HeaderSet $B$KJQ49(B
	headerSet := NewHeaderSet(header)

	// ReferenceSet $B$NCf$+$i>C$9$Y$-CM$r>C$9(B
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
