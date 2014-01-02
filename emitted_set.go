package hpack

import (
	"fmt"
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
	return es[i].Name < es[j].Name
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
