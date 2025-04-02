package main

import (
	"fmt"
	"strings"
)

// func StripDynamicPrefix(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		idx := strings.Index(r.URL.Path, "/static/")
// 		if idx != -1 {
// 			r.URL.Path = r.URL.Path[idx+len("/static/"):] // Hapus bagian sebelum "/static/"
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

func main() {
	value := strings.Index("/A/B/C/static/D", "/static/")

	v := []byte("/A/B/C/static/D")

	v = v[value+len("/static/"):]

	fmt.Println(string(v))
}
