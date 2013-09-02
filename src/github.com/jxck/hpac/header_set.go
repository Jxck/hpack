package hpac

import (
	"net/http"
	"strings"
)

type HeaderSet map[string]string

func NewHeaderSet(header http.Header) HeaderSet {
	headerSet := make(HeaderSet, len(header))
	for name, value := range header {
		headerSet[name] = strings.Join(value, ",")
	}
	return headerSet
}
