package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hallo Mother Father !!!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: &handler,
	}

	server.ListenAndServe()
}
