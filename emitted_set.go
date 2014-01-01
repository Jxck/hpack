package hpack

import (
	"fmt"
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

func NewEmittedSet() *EmittedSet {
	return &EmittedSet{http.Header{}}
}

// TODO: 重複したキーを登録した場合
// e.Header["Hoge"] しないと map が取れない問題
func (e *EmittedSet) Emit(hf *HeaderField) {
	e.Add(hf.Name, hf.Value)
}

func (e *EmittedSet) Dump() (str string) {
	str += "\n-------------- ES --------------\n"
	for i, v := range e.Header {
		str += fmt.Sprintln(i, v)
	}
	str += "--------------------------------\n"
	return str
}
