package hpack

import (
	"net/http"
	"strings"
)

// remove ":" prefix
func RemovePrefix(name string) string {
	if strings.HasPrefix(name, ":") {
		name = strings.TrimLeft(name, ":")
	}
	return name
}

// http.Header => HeaderSet
func HeaderToHeaderSet(header http.Header) HeaderSet {
	headerSet := make(HeaderSet, 0, len(header))
	for name, values := range header {
		// process name
		name = strings.ToLower(name)
		mustname, ok := MustHeader[name]
		if ok {
			name = mustname
		}
		// process values
		value := strings.Join(values, ",")
		headerField := NewHeaderField(name, value)
		headerSet = append(headerSet, headerField)
	}
	return headerSet
}

/*
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
*/
