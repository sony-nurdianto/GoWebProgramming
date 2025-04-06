package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

type WorldHandler struct{}

func (wh *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World")
}

func main() {
	helloHandler := HelloHandler{}
	worldHandler := WorldHandler{}

	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.Handle("/hello", &helloHandler)
	http.Handle("/world", &worldHandler)

	server.ListenAndServe()
}
