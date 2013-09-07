package hpac

import (
	"net/http"
	"strings"
)

type EmittedSet struct {
	http.Header
}

func (e EmittedSet) Emit(name, value string) {
	if strings.HasPrefix(name, ":") {
		name = strings.TrimLeft(name, ":")
	}
	e.Add(name, value)
}

func NewEmittedSet() EmittedSet {
	return EmittedSet{http.Header{}}
}
