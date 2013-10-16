package hpack

import (
	"fmt"
	"net/http"
	"strings"
)

func padding(str string) string {
	l := 20 - len(str)
	return str + strings.Repeat(" ", l)
}

func HeaderString(m http.Header) string {
	str := "\n====Header\n"
	for name, value := range m {
		str += fmt.Sprintf("%s:\t%s\n", padding(name), value)
	}
	str += "====\n"
	return str
}

func HeadersString(h Headers) string {
	str := "\n====Headers\n"
	for i, v := range h {
		str += fmt.Sprintf("%v:\t%s\n", i, v)
	}
	str += "====\n"
	return str
}

func ESString(e EmittedSet) string {
	str := "\n====EmittedSet\n"
	for k, v := range e.Header {
		str += fmt.Sprintf("%v:\t%s\n", k, v)
	}
	str += "====\n"
	return str
}

func RefSetString(r ReferenceSet) string {
	str := "\n====ReferenceSet\n"
	for k, v := range r {
		str += fmt.Sprintf("%v:\t%s\n", k, v)
	}
	str += "====\n"
	return str
}
