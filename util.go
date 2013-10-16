package hpack

import (
	"fmt"
	"net/http"
	"strings"
)

// compare both slice has same value
func CompareSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

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
