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

func MapString(m http.Header) string {
	str := "\n====\n"
	for name, value := range m {
		str += fmt.Sprintf("%s:\t%s\n", padding(name), value)
	}
	str += "====\n"
	return str
}
