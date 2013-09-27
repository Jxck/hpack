package hpack

import (
	"net/http"
	"strings"
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

func (e EmittedSet) Emit(name, value string) {
	name = RemovePrefix(name)
	e.Add(name, value)
}

func (e EmittedSet) Check(name, value string) bool {
	name = RemovePrefix(name)
	return e.Get(name) == value
}

// remove ":" prefix
func RemovePrefix(name string) string {
	if strings.HasPrefix(name, ":") {
		name = strings.TrimLeft(name, ":")
	}
	return name
}
