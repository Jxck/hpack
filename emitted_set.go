package hpack

import (
	"net/http"
)

// Wrapper Type of http.Header
// if you want to user range like http.Header
// you need call range like this
//
// es := NewEmittedSet()
// for name, value := range es.Header {
//    log.Println(name, value)
// }
type EmittedSet struct {
	http.Header
}

func NewEmittedSet() EmittedSet {
	return EmittedSet{http.Header{}}
}

func (e EmittedSet) Emit(hf HeaderField) {
	name := RemovePrefix(hf.Name)
	e.Add(name, hf.Value)
}
