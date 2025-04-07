package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Handler Called : %T\n", h)
			h.ServeHTTP(w, r)
		})
}

func protected(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	helloHandler := HelloHandler{}
	http.Handle("/hello", protected(log(helloHandler)))
	server.ListenAndServe()
}
