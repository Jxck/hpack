package hpack

import (
	"fmt"
	"net/http"
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

func MapString(m http.Header) string {
	str := "\n====\n"
	for name, value := range m {
		str += fmt.Sprintf("%s:\t %s\n", name, value)
	}
	str += "====\n"
	return str
}
