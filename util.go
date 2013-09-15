package hpack

import (
	"fmt"
	"runtime"
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

func Debug(str string) {
	_, file, line, _ := runtime.Caller(1)
	f := strings.Split(file, "/")
	filename := f[len(f)-1]

	fmt.Printf("%v:%v %v\n", filename, line, str)
}
