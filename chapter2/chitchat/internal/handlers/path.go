package handlers

import (
	"net/http"
	"strings"
)

type PrefixHandler struct {
	next http.Handler
}

func NewPrefixHandler(next http.Handler) *PrefixHandler {
	return &PrefixHandler{
		next: next,
	}
}

func (ph *PrefixHandler) StripDynamicPrefix(w http.ResponseWriter, r *http.Request) {
	subStr := "/static/"

	idx := strings.Index(r.URL.Path, subStr)

	if idx != -1 {
		r.URL.Path = r.URL.Path[idx+len(subStr):]
	}

	ph.next.ServeHTTP(w, r)
}
