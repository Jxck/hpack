package hpack

import (
	"fmt"
	"net/http"
)

// List of Emitted Header
// This will pass to Application
type EmittedSet []HeaderField

func NewEmittedSet() *EmittedSet {
	return &EmittedSet{}
}

func (e *EmittedSet) Emit(hf *HeaderField) {
	*e = append(*e, *hf)
}

func (e *EmittedSet) Len() int {
	return len(*e)
}

// Sort Interface
func (e *EmittedSet) Swap(i, j int) {
	es := *e
	es[i], es[j] = es[j], es[i]
}

func (e *EmittedSet) Less(i, j int) bool {
	es := *e
	if es[i].Name == es[j].Name {
		return es[i].Value < es[j].Value
	}
	return es[i].Name < es[j].Name
}

// convert to http.Header
func (e *EmittedSet) ToHeader() http.Header {
	header := make(http.Header)
	for _, hf := range *e {
		header.Add(hf.Name, hf.Value)
	}
	return header
}

// Dump for Debug
func (e *EmittedSet) Dump() (str string) {
	str += "\n-------------- ES --------------\n"
	for i, v := range *e {
		str += fmt.Sprintln(i, v)
	}
	str += "--------------------------------\n"
	return str
}
