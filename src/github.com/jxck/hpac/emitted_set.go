package hpac

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

// remove ":" prefix and call http.Header.Add()
func (e EmittedSet) Emit(name, value string) {
	if strings.HasPrefix(name, ":") {
		name = strings.TrimLeft(name, ":")
	}
	e.Add(name, value)
}
